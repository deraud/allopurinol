package repo

import (
	"context"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
	"strings"
)

func (r *OrderRepoImpl) GetPendingOrders(ctx context.Context, options *entity.PendingOrderOptions, userId int64) ([]*entity.PendingOrderGroup, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	groupOrders := make([]*entity.PendingOrderGroup, size)

	sortBy := options.SortBy
	if options.SortBy == "created_at" {
		sortBy = "og.created_at"
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("og.id", "o.id", "o.pharmacy_id", "MIN(DISTINCT p.name)", "ARRAY_AGG(oi.catalog_id)",
		"ARRAY_AGG(oi.quantity)", "ARRAY_AGG(c.price*oi.quantity)", "MIN(a.name)", "og.created_at",
		"ARRAY_AGG(p2.name)", "o.total_price_product", "o.total_price_shipping").
		From("orders o").
		Join("order_groups og", "og.id = o.order_group_id").
		Join("order_items oi", "oi.order_id = o.id").
		Join("addresses a", "a.id = o.address_id ").
		Join("pharmacies p", "p.id = o.pharmacy_id").
		Join("catalogs c", "c.id = oi.catalog_id").
		Join("products p2", "p2.id = c.product_id").
		AddWhereClause(generateGetOrdersWhereClause(userId)).
		OrderBy(fmt.Sprintf("%s %s", sortBy, options.SortOrder)).
		GroupBy("og.id, o.id").
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := map[int64][]*entity.PendingOrder{}

	var total = map[int64]decimal.Decimal{}
	var shipping = map[int64]decimal.Decimal{}
	var addresses = map[int64]string{}
	var createdAts = map[int64]pgtype.Timestamp{}

	for i := 0; rows.Next(); i++ {
		var groupId int64

		var catalogIds []int64
		var catalogQuantities []int64
		var catalogPrices []decimal.Decimal
		var productNames []string

		var totalScan, shippingScan decimal.Decimal
		var address string
		var createdAt pgtype.Timestamp

		var order = &entity.PendingOrder{}
		err := rows.Scan(
			&groupId,
			&order.Id,
			&order.PharmacyId,
			&order.PharmacyName,
			pq.Array(&catalogIds),
			pq.Array(&catalogQuantities),
			pq.Array(&catalogPrices),
			&address,
			&createdAt,
			pq.Array(&productNames),
			&totalScan,
			&shippingScan,
		)
		if err != nil {
			return nil, err
		}
		order.ShippingCost = shippingScan

		_, exist := total[groupId]
		if !exist {
			total[groupId] = totalScan
		} else {
			total[groupId] = total[groupId].Add(totalScan)
		}

		_, exist = shipping[groupId]
		if !exist {
			shipping[groupId] = shippingScan
		} else {
			shipping[groupId] = shipping[groupId].Add(shippingScan)
		}

		var catalogs []*entity.PendingCatalog
		for i := 0; i < len(catalogIds); i++ {
			catalogs = append(catalogs, &entity.PendingCatalog{
				Id:       catalogIds[i],
				Name:     productNames[i],
				Quantity: catalogQuantities[i],
				Price:    catalogPrices[i],
			})
		}
		order.Catalogs = catalogs

		_, exist = ordersMap[groupId]
		if !exist {
			ordersMap[groupId] = []*entity.PendingOrder{order}
			addresses[groupId] = address
			createdAts[groupId] = createdAt
		} else {
			ordersMap[groupId] = append(ordersMap[groupId], order)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	counter := 0
	for groupId, orders := range ordersMap {
		groupOrders[counter] = &entity.PendingOrderGroup{
			Id:           groupId,
			Order:        orders,
			ShippingCost: shipping[groupId],
			TotalPrice:   total[groupId].Add(shipping[groupId]),
			UserAddress:  addresses[groupId],
			CreatedAt:    createdAts[groupId],
		}
		counter++
	}

	return groupOrders, nil
}

func generateGetOrdersWhereClause(userId int64) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.Equal("o.user_id", userId),
		cond.Equal("o.status_id", 1),
		cond.IsNull("o.deleted_at"),
		cond.IsNull("og.deleted_at"),
		cond.IsNull("oi.deleted_at"),
		cond.IsNull("p.deleted_at"),
		cond.IsNull("c.deleted_at"),
	)
	return whereClause
}

func (r *OrderRepoImpl) GetPendingOrdersCount(ctx context.Context, userId int64) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(DISTINCT og.id)").
		From("order_groups og").
		Join("orders o", "o.order_group_id = og.id").
		Where(
			sb.Equal("og.user_id", userId),
			sb.Equal("o.status_id", 1),
			sb.IsNull("og.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *OrderRepoImpl) GetPharmacyOrders(ctx context.Context, options *entity.PharmacyOrderOptions, pharmacyId int64) ([]*entity.Order, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	orders := make([]*entity.Order, size)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("o.id", "os.name", "COUNT(oi.id)", "o.total_price_product").
		From("orders o").
		Join("order_statuses os", "o.status_id = os.id").
		Join("order_items oi", "o.id = oi.order_id").
		Where(sb.Equal("o.pharmacy_id", pharmacyId)).
		GroupBy("o.id, os.id").
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			order       entity.Order
			orderStatus entity.OrderStatus
			itemCount   int
		)
		err := rows.Scan(
			&order.Id,
			&orderStatus.Name,
			&itemCount,
			&order.TotalPriceProduct,
		)
		if err != nil {
			return nil, err
		}
		order.OrderStatus = &orderStatus
		order.OrderItems = make([]*entity.OrderItem, itemCount)
		orders[i] = &order
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepoImpl) GetPharmacyOrdersCount(ctx context.Context, options *entity.PharmacyOrderOptions, pharmacyId int64) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(DISTINCT(o.id))").
		From("orders o").
		Join("order_statuses os", "o.status_id = os.id").
		Join("order_items oi", "o.id = oi.order_id").
		Where(sb.Equal("o.pharmacy_id", pharmacyId))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *OrderRepoImpl) GetOrder(ctx context.Context, id int64) (*entity.Order, error) {
	var (
		order           entity.Order
		orderStatus     entity.OrderStatus
		paymentMethod   entity.PaymentMethod
		logisticPartner entity.LogisticPartner
		user            entity.User
		address         entity.Address
		pharmacy        entity.Pharmacy
	)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("o.id", "o.total_price_product", "o.total_price_shipping", "u.id", "u.name", "a.id", "a.name", "os.id", "os.name", "pm.id", "pm.name", "lp.id", "lp.name", "p.id", "p.name").
		From("orders o").
		Join("order_statuses os", "o.status_id = os.id").
		Join("payment_methods pm", "o.payment_method_id = pm.id").
		Join("logistic_partners lp", "o.logistic_partner_id = lp.id").
		Join("users u", "o.user_id = u.id").
		Join("addresses a", "o.address_id = a.id").
		Join("pharmacies p", "o.pharmacy_id = p.id").
		Where(
			sb.Equal("o.id", id),
			sb.IsNull("o.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&order.Id,
		&order.TotalPriceProduct,
		&order.TotalPriceShipping,
		&user.Id,
		&user.Name,
		&address.Id,
		&address.Name,
		&orderStatus.Id,
		&orderStatus.Name,
		&paymentMethod.Id,
		&paymentMethod.Name,
		&logisticPartner.Id,
		&logisticPartner.Name,
		&pharmacy.Id,
		&pharmacy.Name,
	)
	if err != nil {
		return nil, err
	}

	order.User = &user
	order.Address = &address
	order.OrderStatus = &orderStatus
	order.LogisticPartner = &logisticPartner
	order.PaymentMethod = &paymentMethod
	order.Pharmacy = &pharmacy
	return &order, nil
}

func (r *OrderRepoImpl) GetOrderItems(ctx context.Context, orderId int64) ([]*entity.OrderItem, error) {
	var orderItems []*entity.OrderItem

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("oi.id", "oi.quantity", "c.id", "c.price", "p.id", "p.name", "p.image").
		From("order_items oi").
		Join("catalogs c", "oi.catalog_id = c.id").
		Join("products p", "c.product_id = p.id").
		Where(
			sb.Equal("oi.order_id", orderId),
			sb.IsNull("oi.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			orderItem entity.OrderItem
			catalog   entity.Catalog
		)
		err := rows.Scan(
			&orderItem.Id,
			&orderItem.Quantity,
			&catalog.Id,
			&catalog.Price,
			&catalog.ProductId,
			&catalog.Name,
			&catalog.Image,
		)
		if err != nil {
			return nil, err
		}
		orderItem.Catalog = &catalog
		orderItems = append(orderItems, &orderItem)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *OrderRepoImpl) GetOrders(ctx context.Context, options *entity.OrderOptions) ([]*entity.Order, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	orders := make([]*entity.Order, size)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("o.id", "p.id", "p.name", "os.id", "os.name").
		From("orders o").
		Join("order_statuses os", "o.status_id = os.id").
		Join("pharmacies p", "o.pharmacy_id = p.id").
		AddWhereClause(r.generateGetOrdersWhereClause(options)).
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			order       entity.Order
			orderStatus entity.OrderStatus
			pharmacy    entity.Pharmacy
		)
		err := rows.Scan(
			&order.Id,
			&pharmacy.Id,
			&pharmacy.Name,
			&orderStatus.Id,
			&orderStatus.Name,
		)
		if err != nil {
			return nil, err
		}
		order.OrderStatus = &orderStatus
		order.Pharmacy = &pharmacy
		orders[i] = &order
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepoImpl) GetUserOrders(ctx context.Context, options *entity.UserOrderOptions, userId int64) ([]*entity.Order, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	orders := make([]*entity.Order, size)

	sb := sqlbuilder.NewSelectBuilder()

	filter := ""
	if options.FilterValue != "" {
		filter = sb.Equal("LOWER(os.name)", strings.ToLower(options.FilterValue))
	}

	sb.Select("o.id", "os.name", "ARRAY_AGG(p.name)", "ARRAY_AGG(p.image)", "ARRAY_AGG(oi.id)",
		"ARRAY_AGG(oi.quantity)", "ARRAY_AGG(c.price*oi.quantity)", "MIN(p2.name)", "MIN(a.name)", "o.created_at",
		"o.total_price_product", "o.total_price_shipping ").
		From("orders o").
		Join("order_statuses os", "o.status_id = os.id").
		Join("order_items oi", "o.id = oi.order_id").
		Join("addresses a", "a.id = o.address_id ").
		Join("catalogs c", "c.id = oi.catalog_id").
		Join("products p", "p.id = c.product_id").
		Join("pharmacies p2", "p2.id = c.pharmacy_id").
		Where(
			sb.Equal("o.user_id", userId),
			filter,
		).
		GroupBy("o.id, os.id").
		OrderBy("created_at DESC").
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			order       entity.Order
			orderStatus entity.OrderStatus
			names       []string
			images      []string
			itemIds     []int64
			quantities  []int64
			prices      []float64
		)
		err := rows.Scan(
			&order.Id,
			&orderStatus.Name,
			pq.Array(&names),
			pq.Array(&images),
			pq.Array(&itemIds),
			pq.Array(&quantities),
			pq.Array(&prices),
			&order.PharmacyName,
			&order.AddressName,
			&order.CreatedAt,
			&order.TotalPriceProduct,
			&order.TotalPriceShipping,
		)
		if err != nil {
			return nil, err
		}
		order.OrderStatus = &orderStatus
		order.TotalPriceProduct = order.TotalPriceProduct.Add(order.TotalPriceShipping)
		order.OrderItems = make([]*entity.OrderItem, len(names))
		for i := 0; i < len(names); i++ {
			order.OrderItems[i] = &entity.OrderItem{
				Quantity: int(quantities[i]),
				Price:    decimal.NewFromFloat(prices[i]),
				Catalog: &entity.Catalog{
					Id:    itemIds[i],
					Name:  names[i],
					Image: images[i],
				},
			}
		}
		orders[i] = &order
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepoImpl) GetOrdersCount(ctx context.Context, options *entity.OrderOptions) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(o.id)").
		From("orders o").
		Join("order_statuses os", "o.status_id = os.id").
		Join("pharmacies p", "o.pharmacy_id = p.id").
		AddWhereClause(r.generateGetOrdersWhereClause(options))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *OrderRepoImpl) GetUserOrdersCount(ctx context.Context, options *entity.UserOrderOptions, userId int64) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()

	filter := ""
	if options.FilterValue != "" {
		filter = sb.Equal("LOWER(os.name)", strings.ToLower(options.FilterValue))
	}

	sb.Select("COUNT(o.id)").
		From("orders o").
		Join("order_statuses os", "o.status_id = os.id").
		Where(
			sb.Equal("o.user_id", userId),
			filter,
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *OrderRepoImpl) generateGetOrdersWhereClause(options *entity.OrderOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.IsNull("o.deleted_at"),
	)

	if options.PharmacyId != nil {
		whereClause.AddWhereExpr(cond.Args, cond.Equal("o.pharmacy_id", *options.PharmacyId))
	}
	if options.StatusId != nil {
		whereClause.AddWhereExpr(cond.Args, cond.Equal("o.status_id", *options.StatusId))
	}
	return whereClause
}
