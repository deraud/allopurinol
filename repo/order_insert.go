package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *OrderRepoImpl) InsertOrderGroup(ctx context.Context,  o *entity.OrderGroup) (*entity.OrderGroup, error) {
	orderGroup := *o

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("order_groups").
		Cols("user_id").
		Values(o.UserId).
		SQL("RETURNING id")

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&orderGroup.Id)
	if err != nil {
		return nil, err
	}

	return &orderGroup, nil
}

func (r *OrderRepoImpl) InsertOrder(ctx context.Context, o *entity.Order) (*entity.Order, error) {
	order := *o

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("orders").
		Cols("user_id", "address_id", "status_id", "payment_method_id", "pharmacy_id", "logistic_partner_id", "order_group_id", "total_price_product", "total_price_shipping").
		Values(o.UserId, o.AddressId, o.StatusId, o.PaymentMethodId, o.PharmacyId, o.LogisticPartnerId, o.OrderGroupId, o.TotalPriceProduct, o.TotalPriceShipping).
		SQL("RETURNING id")

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&order.Id)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepoImpl) InsertOrderItemBulk(ctx context.Context, items []*entity.OrderItem) ([]*entity.OrderItem, error) {
	orderItems := make([]*entity.OrderItem, len(items))

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("order_items").
		Cols("order_id", "catalog_id", "quantity")
	for _, i := range items {
		ib.Values(i.OrderId, i.CatalogId, i.Quantity)
	}
	ib.SQL("RETURNING id")

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		orderItem := *items[i]
		err := rows.Scan(&orderItem.Id)
		if err != nil {
			return nil, err
		}
		orderItems[i] = &orderItem
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}
