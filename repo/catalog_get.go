package repo

import (
	"context"
	"fmt"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *CatalogRepoImpl) GetCatalogs(ctx context.Context, options *entity.CatalogOptions, pharmacyId int64) ([]*entity.Catalog, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	catalogs := make([]*entity.Catalog, size)

	if options.SortBy != "stock" {
		options.SortBy = "p." + options.SortBy 
	}else{
		options.SortBy = "c." + options.SortBy
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("c.id", "p.name", "p.generic_name", "m.name", "pc.name", "pf.name", "c.stock", "c.is_active").
		From("catalogs c").
		Join("products p", "c.product_id = p.id").
		Join("manufacturers m", "p.manufacturer_id = m.id").
		Join("product_classifications pc", "p.product_classification_id = pc.id").
		JoinWithOption(sqlbuilder.LeftOuterJoin ,"product_forms pf", "p.product_form_id = pf.id").
		Where(sb.Equal("c.pharmacy_id", pharmacyId)).
		AddWhereClause(r.generateGetCatalogsWhereClause(options)).
		OrderBy(fmt.Sprintf("%s %s", options.SortBy, options.SortOrder)).
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var catalog entity.Catalog
		err := rows.Scan(
			&catalog.Id,
			&catalog.Name,
			&catalog.GenericName,
			&catalog.Manufacturer,
			&catalog.Classification,
			&catalog.Form,
			&catalog.Stock,
			&catalog.IsActive,
		)
		if err != nil {
			return nil, err
		}
		catalogs[i] = &catalog
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return catalogs, nil
}

func (r *CatalogRepoImpl) GetCatalogsCount(ctx context.Context, options *entity.CatalogOptions, pharmacyId int64) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(c.id)").
		From("catalogs c").
		Join("products p", "c.product_id = p.id").
		Join("manufacturers m", "p.manufacturer_id = m.id").
		Join("product_classifications pc", "p.product_classification_id = pc.id").
		JoinWithOption(sqlbuilder.LeftOuterJoin ,"product_forms pf", "p.product_form_id = pf.id").
		Where(sb.Equal("c.pharmacy_id", pharmacyId)).
		AddWhereClause(r.generateGetCatalogsWhereClause(options))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *CatalogRepoImpl) generateGetCatalogsWhereClause(options *entity.CatalogOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.Like(fmt.Sprintf("p.%s", options.SearchBy), fmt.Sprintf("%%%s%%", options.SearchValue)),
		cond.IsNull("c.deleted_at"),
	)

	if options.ManufacturerId != nil {
		whereClause.AddWhereExpr(cond.Args, cond.Equal("p.manufacturer_id", *options.ManufacturerId))
	}
	if options.ClassificationId != nil {
		whereClause.AddWhereExpr(cond.Args, cond.Equal("p.product_classification_id", *options.ClassificationId))
	}
	if options.FormId != nil {
		whereClause.AddWhereExpr(cond.Args, cond.Equal("p.product_form_id", *options.FormId))
	}
	if options.IsActive != nil {
		whereClause.AddWhereExpr(cond.Args, cond.Equal("c.is_active", *options.IsActive))
	}
	return whereClause
}

func (r *CatalogRepoImpl) GetCatalog(ctx context.Context, id int64) (*entity.Catalog, error) {
	var catalog entity.Catalog

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"c.id",
		"c.stock",
		"c.price",
		"c.is_active",
		"p.name",
		"p.generic_name",
		"m.name",
		"pc.name",
		"pf.name",
		"p.description",
		"p.unit_in_pack",
		"p.selling_unit",
		"p.image",
	).
		From("catalogs c").
		Join("products p", "c.product_id = p.id").
		Join("manufacturers m", "p.manufacturer_id = m.id").
		Join("product_classifications pc", "p.product_classification_id = pc.id").
		JoinWithOption(sqlbuilder.LeftOuterJoin ,"product_forms pf", "p.product_form_id = pf.id").
		Where(
			sb.Equal("c.id", id),
			sb.IsNull("p.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&catalog.Id,
		&catalog.Stock,
		&catalog.Price,
		&catalog.IsActive,
		&catalog.Name,
		&catalog.GenericName,
		&catalog.Manufacturer,
		&catalog.Classification,
		&catalog.Form,
		&catalog.Description,
		&catalog.UnitInPack,
		&catalog.SellingUnit,
		&catalog.Image,
	)
	if err != nil {
		return nil, err
	}

	return &catalog, nil
}

func (r *CatalogRepoImpl) GetAvailableCatalogs(ctx context.Context, options *entity.AvailableCatalogOptions, userAddress *entity.Address) ([]*entity.Catalog, int, error) {
	var (
		catalogs []*entity.Catalog
		count int
	)

	filterCatId := `
	    WHERE 
	   		is_active = TRUE AND 
	   		deleted_at IS NULL AND
			name ILIKE $1`
	if options.CategoryId != nil {
		filterCatId = fmt.Sprintf(`
		INNER JOIN
		    product_category_maps pcm
				ON p.id = pcm.product_id
		WHERE 
	   		p.is_active = TRUE AND 
	   		p.deleted_at IS NULL AND
			p.name ILIKE $1 AND
			pcm.product_category_id = %d`, *options.CategoryId)
	}

	query := fmt.Sprintf(`
	WITH main_data AS (
	SELECT
		c.id, c.price, c.stock,
		p.id AS product_id, p.name, p.image, p.selling_unit,
		ph.location, ph.days
	FROM (
		SELECT 
			id,
			product_id,
			pharmacy_id,
			price,
			stock
		FROM 
			catalogs
		WHERE 
			is_active = TRUE AND 
			deleted_at IS NULL AND
			stock > 0
		) AS c
	   		INNER JOIN (
	   			SELECT 
	   				p.id, name, image, selling_unit
	   			FROM 
	   				products p
	   			%s
	   		) AS p
	         	ON c.product_id = p.id
	        INNER JOIN (
	        	SELECT 
	   				id, location, days
	   			FROM 
	   				pharmacies
	   			WHERE 
	   				is_active = TRUE AND 
	   				deleted_at IS NULL AND
	   				ST_DWithin(location, ST_SetSRID(ST_MakePoint($2, $3), 4326), $4)
	        ) AS ph
				ON c.pharmacy_id = ph.id
	),
	data_summary AS (
		SELECT 
			MIN(price) AS min_price, 	
			MAX(price) AS max_price,
			MIN(ST_DistanceSphere(location::geometry, ST_MakePoint($2, $3))) AS min_dist, 
			MAX(ST_DistanceSphere(location::geometry, ST_MakePoint($2, $3))) AS max_dist,
			CAST(MIN(days) AS FLOAT) AS min_days,
			CAST(MAX(days) AS FLOAT) AS max_days
		FROM
		    main_data
	)
	SELECT
		id, product_id, name, image, selling_unit, price, stock, COUNT(1) OVER() AS total
	FROM (
		SELECT 
			DISTINCT ON(product_id)
				id, product_id, name, image, selling_unit, price, stock, score
		FROM
		(
			SELECT 
				main_data.*,
				CASE
					WHEN data_summary.max_days = data_summary.min_days THEN $5
					ELSE $5 * (data_summary.max_days - main_data.days)/(data_summary.max_days - data_summary.min_days)
				END +
				CASE
					WHEN data_summary.max_price = data_summary.min_price THEN $6
					ELSE $6 * (data_summary.max_price - main_data.price)/(data_summary.max_price - data_summary.min_price)
				END +
				CASE
					WHEN data_summary.max_dist = data_summary.min_dist THEN $7
					ELSE $7 * (data_summary.max_dist - ST_DistanceSphere(main_data.LOCATION::geometry, ST_MakePoint($2, $3)))/(data_summary.max_dist - data_summary.min_dist)
				END AS score
			FROM 
				main_data
					CROSS JOIN data_summary
			ORDER BY score DESC
		)
	)
	ORDER BY score DESC
	LIMIT $8
	OFFSET $9;
	`, filterCatId)
	args := []any{
		fmt.Sprintf("%%%s%%", options.SearchValue),
		userAddress.Longitude,
		userAddress.Latitude,
		constant.PHARMACY_MAX_DISTANCE,
		constant.FASTEST_WEIGHT,
		constant.CHEAPEST_WEIGHT,
		constant.NEAREST_WEIGHT,
		options.Limit,
		(options.Page - 1) * options.Limit,
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			catalog entity.Catalog
			product entity.Product
			total int
		)
		err := rows.Scan(
			&catalog.Id,
			&product.Id,
			&product.Name,
			&product.ImageLink,
			&product.SellingUnit,
			&catalog.Price,
			&catalog.Stock,
			&total,
		)
		if err != nil {
			return nil, 0, err
		}
		if i == 0 {
			count = total
		}
		catalog.Product = &product
		catalogs = append(catalogs, &catalog)
	}
	err = rows.Err()
	if err != nil {
		return nil, 0, err
	}
	return catalogs, count, nil
}

func (r *CatalogRepoImpl) GetAvailableCatalogsCount(ctx context.Context, options *entity.AvailableCatalogOptions, userAddress *entity.Address) (int, error) {
	var count int

	query := `
	SELECT
		COUNT(DISTINCT p.id) AS data_count
	FROM (
		SELECT 
			product_id,
			pharmacy_id
		FROM 
			catalogs
		WHERE 
			is_active = TRUE AND 
			deleted_at IS NULL AND
			stock > 0
		) AS c
			INNER JOIN (
				SELECT 
					id
				FROM 
					products
				WHERE 
					is_active = TRUE AND 
					deleted_at IS NULL AND
					name ILIKE $1 
			) AS p
				ON c.product_id = p.id
			INNER JOIN (
				SELECT 
					id
				FROM 
					pharmacies
				WHERE 
					is_active = TRUE AND 
					deleted_at IS NULL AND
					ST_DWithin(location, ST_SetSRID(ST_MakePoint($2, $3), 4326), $4)
			) AS ph
				ON c.pharmacy_id = ph.id
	`
	args := []any{
		fmt.Sprintf("%%%s%%", options.SearchValue),
		userAddress.Longitude,
		userAddress.Latitude,
		constant.PHARMACY_MAX_DISTANCE,
	}
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *CatalogRepoImpl) GetAvailableCatalog(ctx context.Context, id int64) (*entity.Catalog, error) {
	var (
		catalog  entity.Catalog
		product  entity.Product
		pharmacy entity.Pharmacy
		address  entity.Address
		catIds   string
		catNames string
	)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"c.id",
		"c.stock",
		"c.price",
		"p.id",
		"p.name",
		"p.generic_name",
		"m.name",
		"pcl.name",
		"pf.name",
		"p.description",
		"p.unit_in_pack",
		"p.selling_unit",
		"p.weight",
		"p.height",
		"p.length",
		"p.width",
		"p.image",
		"ARRAY_AGG(pc.id)",
		"ARRAY_AGG(pc.name)",
		"ph.id",
		"ph.name",
		"a.name",
		"a.province",
		"a.city",
		"a.district",
		"a.subdistrict",
		"a.postal_code",
	).
		From("catalogs c").
		Join("products p", "c.product_id = p.id").
		Join("pharmacies ph", "c.pharmacy_id = ph.id").
		Join("product_category_maps pcm", "p.id = pcm.product_id").
		Join("product_categories pc", "pcm.product_category_id = pc.id").
		Join("product_classifications pcl", "p.product_classification_id = pcl.id").
		JoinWithOption(sqlbuilder.LeftOuterJoin, "product_forms pf", "p.product_form_id = pf.id").
		Join("manufacturers m", "p.manufacturer_id = m.id").
		Join("addresses a", "ph.id = a.pharmacy_id").
		Where(
			sb.Equal("c.id", id),
			sb.IsNull("c.deleted_at"),
		).
		GroupBy("c.id", "p.id", "m.name", "pcl.name", "pf.name", "pc.id", "ph.id", "a.id")

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&catalog.Id,
		&catalog.Stock,
		&catalog.Price,
		&product.Id,
		&product.Name,
		&product.GenericName,
		&product.Manufacturer,
		&product.Classification,
		&product.Form,
		&product.Description,
		&product.UnitInPack,
		&product.SellingUnit,
		&product.Weight,
		&product.Height,
		&product.Length,
		&product.Width,
		&product.ImageLink,
		&catIds,
		&catNames,
		&pharmacy.Id,
		&pharmacy.Name,
		&address.Name,
		&address.Province,
		&address.City,
		&address.District,
		&address.Subdistrict,
		&address.PostalCode,
	)
	if err != nil {
		return nil, err
	}

	product.CategoryIds, err = ParseArrayAggInt64(catIds)
	if err != nil {
		return nil, err
	}

	product.Categories = ParseArrayAggString(catNames)
	catalog.Product = &product
	pharmacy.Address = &address
	catalog.Pharmacy = &pharmacy
	return &catalog, nil
}

func (r *CatalogRepoImpl) GetCheckoutCatalogs(ctx context.Context, cartItems []*entity.CartItem, userAddress *entity.Address) ([]*entity.Catalog, error) {
	var catalogs []*entity.Catalog

	wb := r.generateGetCheckoutCatalogsWhereClause(cartItems, userAddress)
	whereQuery, args := wb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	query := fmt.Sprintf(`
		SELECT 
			DISTINCT ON(c.product_id)
				c.id, c.stock, c.price, p.id, p.name, p.image, p.weight, ph.id, ph.name, a.name
		FROM
			catalogs c
				INNER JOIN products p
					ON c.product_id = p.id
				INNER JOIN pharmacies ph
					ON c.pharmacy_id = ph.id
				INNER JOIN addresses a
					ON c.pharmacy_id = a.pharmacy_id
		%s
		ORDER BY	 
			c.product_id, 
			ST_Distance(a.location, ST_SetSRID(ST_MakePoint(%f, %f), 4326)) DESC
	`, whereQuery, userAddress.Longitude, userAddress.Latitude)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			catalog  entity.Catalog
			product  entity.Product
			pharmacy entity.Pharmacy
			address  entity.Address
		)
		err := rows.Scan(
			&catalog.Id,
			&catalog.Stock,
			&catalog.Price,
			&product.Id,
			&product.Name,
			&product.ImageLink,
			&product.Weight,
			&pharmacy.Id,
			&pharmacy.Name,
			&address.Name,
		)
		if err != nil {
			return nil, err
		}
		catalog.Product = &product
		pharmacy.Address = &address
		catalog.Pharmacy = &pharmacy
		catalogs = append(catalogs, &catalog)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return catalogs, nil
}

func (r *CatalogRepoImpl) generateGetCheckoutCatalogsWhereClause(cartItems []*entity.CartItem, userAddress *entity.Address) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereIdStatement := make([]string, len(cartItems))
	for _, item := range cartItems {
		whereIdStatement = append(whereIdStatement, fmt.Sprintf("(c.product_id=%d AND c.stock >= %d)", item.ProductId, item.Quantity))
	}

	whereClause.AddWhereExpr(
		cond.Args,
		cond.LTE(fmt.Sprintf("ST_Distance(a.location, ST_SetSRID(ST_MakePoint(%f, %f), 4326))", userAddress.Longitude, userAddress.Latitude), constant.PHARMACY_MAX_DISTANCE),
		cond.Equal("p.is_active", true),
		cond.Equal("ph.is_active", true),
		cond.Equal("c.is_active", true),
		cond.GT("c.stock", 0),
		cond.IsNull("c.deleted_at"),
		cond.Or(whereIdStatement...),
	)
	return whereClause
}

func (r *CatalogRepoImpl) GetCatalogsByProductIds(ctx context.Context, productIds []int64) ([]*entity.Catalog, error) {
	var catalogs []*entity.Catalog

	wb := r.generateGetCatalogsByProductIdsWhereClause(productIds)
	whereQuery, args := wb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	query := fmt.Sprintf(`
		SELECT 
			DISTINCT ON(c.product_id)
				c.id, p.id, p.name, p.image, p.selling_unit, c.stock, c.price
		FROM
			catalogs c
				INNER JOIN products p
					ON c.product_id = p.id
		%s
	`, whereQuery)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var (
			catalog entity.Catalog
			product entity.Product
		)
		err := rows.Scan(
			&catalog.Id,
			&product.Id,
			&product.Name,
			&product.ImageLink,
			&product.SellingUnit,
			&catalog.Price,
			&catalog.Stock,
		)
		if err != nil {
			return nil, err
		}
		catalog.Product = &product
		catalogs = append(catalogs, &catalog)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return catalogs, nil
}

func (r *CatalogRepoImpl) generateGetCatalogsByProductIdsWhereClause(productIds []int64) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereIdStatement := make([]string, len(productIds))
	for i, id := range productIds {
		whereIdStatement[i] = fmt.Sprintf("p.id=%d", id)
	}

	whereClause.AddWhereExpr(
		cond.Args,
		cond.Equal("p.is_active", true),
		cond.Equal("c.is_active", true),
		cond.GT("c.stock", 0),
		cond.IsNull("c.deleted_at"),
		cond.Or(whereIdStatement...),
	)

	return whereClause
}
