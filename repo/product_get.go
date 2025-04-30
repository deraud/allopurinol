package repo

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
	"strconv"
	"strings"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

func (r *ProductRepoImpl) GetProductCategories(ctx context.Context, options *entity.ProductCategoryOptions) ([]*entity.ProductCategory, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	productCats := make([]*entity.ProductCategory, size)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name").
		From("product_categories").
		AddWhereClause(generateGetProductCategoriesWhereClause(options)).
		OrderBy(fmt.Sprintf("name %s", options.SortOrder)).
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var productCat = &entity.ProductCategory{}
		err := rows.Scan(
			&productCat.Id,
			&productCat.Name,
		)
		if err != nil {
			return nil, err
		}
		productCats[i] = productCat
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return productCats, nil
}

func generateGetProductCategoriesWhereClause(options *entity.ProductCategoryOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	whereClause.AddWhereExpr(
		cond.Args,
		cond.ILike(options.SearchBy, fmt.Sprintf("%%%s%%", options.SearchValue)),
		cond.IsNull("deleted_at"),
	)
	return whereClause
}

func (r *ProductRepoImpl) GetProductCategoriesCount(ctx context.Context, options *entity.ProductCategoryOptions) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(id)").
		From("product_categories").
		AddWhereClause(generateGetProductCategoriesWhereClause(options))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func ParseArrayAggString(s string) []string {
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	s = strings.ReplaceAll(s, "\"", "")
	return strings.Split(s, ",")
}

func ParseArrayAggInt64(s string) ([]int64, error) {
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	s = strings.ReplaceAll(s, "\"", "")
	var res []int64
	for _, v := range strings.Split(s, ",") {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}

func (r *ProductRepoImpl) GetProducts(ctx context.Context, options *entity.ProductOptions) ([]*entity.Product, error) {
	size := options.Limit
	remaining := options.TotalRows - (options.Page-1)*options.Limit
	if remaining <= 0 {
		size = 0
	} else if remaining < options.Limit {
		size = remaining
	}
	products := make([]*entity.Product, size)

	if options.SortBy != "stock" && options.SortBy != "usage" {
		options.SortBy = fmt.Sprintf("p.%s", options.SortBy)
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "ARRAY_AGG(DISTINCT pc.id) as category_ids", "ARRAY_AGG(DISTINCT pc.name) as categories", "pcl.name", "pf.name", "m.name", "p.name",
		"p.generic_name", "SUM(c.stock) as stock", "p.description", "COUNT(ph.id) as usage", "p.image", "p.product_classification_id",
		"p.product_form_id", "p.manufacturer_id").
		From("products p").
		JoinWithOption(sqlbuilder.LeftOuterJoin, "catalogs c", "c.product_id = p.id").
		JoinWithOption(sqlbuilder.LeftOuterJoin, "pharmacies ph", "ph.id = c.pharmacy_id").
		Join("product_category_maps pcm", "pcm.product_id = p.id").
		Join("product_categories pc", "pc.id = pcm.product_category_id").
		Join("product_classifications pcl", "pcl.id = p.product_classification_id").
		JoinWithOption(sqlbuilder.LeftJoin, "product_forms pf", "pf.id = p.product_form_id").
		Join("manufacturers m", "m.id = p.manufacturer_id").
		AddWhereClause(generateGetProductsWhereClause(options)).
		OrderBy(fmt.Sprintf("%s %s", options.SortBy, options.SortOrder)).
		GroupBy("p.id, pcl.name, pf.name, m.name").
		Offset((options.Page - 1) * options.Limit).
		Limit(options.Limit)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var product = &entity.Product{}
		var categoryIds []int64
		var categories []string
		err := rows.Scan(
			&product.Id,
			pq.Array(&categoryIds),
			pq.Array(&categories),
			&product.Classification,
			&product.Form,
			&product.Manufacturer,
			&product.Name,
			&product.GenericName,
			&product.Stock,
			&product.Description,
			&product.Usage,
			&product.ImageLink,
			&product.ClassificationId,
			&product.FormId,
			&product.ManufacturerId,
		)
		if err != nil {
			return nil, err
		}
		products[i] = product
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func generateGetProductsWhereClause(options *entity.ProductOptions) *sqlbuilder.WhereClause {
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	var equal string
	if options.FilterBy == "classification_id" || options.FilterBy == "form_id" {
		options.FilterBy = fmt.Sprintf("product_%s", options.FilterBy)
	}

	if options.FilterBy != "" && options.FilterValue != "" {
		equal = cond.Equal(fmt.Sprintf("p.%s", options.FilterBy), options.FilterValue)
	}

	whereClause.AddWhereExpr(
		cond.Args,
		equal,
		cond.ILike(fmt.Sprintf("p.%s", options.SearchBy), fmt.Sprintf("%%%s%%", options.SearchValue)),
		cond.IsNull("p.deleted_at"),
	)
	return whereClause
}

func (r *ProductRepoImpl) GetProductsCount(ctx context.Context, options *entity.ProductOptions) (int, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(p.id)").
		From("products p").
		AddWhereClause(generateGetProductsWhereClause(options))

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *ProductRepoImpl) GetProductCatalogIds(ctx context.Context, id int64) ([]int64, error) {
	var catalogs []int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("ARRAY_AGG(c.id)").
		From("products p").
		JoinWithOption(sqlbuilder.LeftOuterJoin, "catalogs c", "c.product_id = p.id").
		Where(
			sb.Equal("c.product_id", id),
			sb.IsNull("p.deleted_at"),
			sb.IsNull("c.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)

	var ids []int64
	err := r.db.QueryRowContext(ctx, query, args...).Scan(pq.Array(&ids))
	if err != nil {
		return nil, err
	}
	if ids == nil {
		return nil, nil
	}

	return catalogs, nil
}

func (r *ProductRepoImpl) GetProductDetail(ctx context.Context, id int64) (*entity.Product, error) {
	product := entity.Product{}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "ARRAY_AGG(DISTINCT pc.id) as category_ids", "ARRAY_AGG(DISTINCT pc.name) as categories", "pcl.name", "pf.name", "m.name", "p.name",
		"p.generic_name", "SUM(c.stock) as stock", "p.description", "COUNT(ph.id) as usage", "p.unit_in_pack", "p.selling_unit",
		"p.weight", "p.height", "p.length", "p.width", "p.image", "p.product_classification_id", "p.product_form_id", "p.manufacturer_id").
		From("products p").
		JoinWithOption(sqlbuilder.LeftOuterJoin, "catalogs c", "c.product_id = p.id").
		JoinWithOption(sqlbuilder.LeftOuterJoin, "pharmacies ph", "ph.id = c.pharmacy_id").
		Join("product_category_maps pcm", "pcm.product_id = p.id").
		Join("product_categories pc", "pc.id = pcm.product_category_id").
		Join("product_classifications pcl", "pcl.id = p.product_classification_id").
		JoinWithOption(sqlbuilder.LeftJoin, "product_forms pf", "pf.id = p.product_form_id").
		Join("manufacturers m", "m.id = p.manufacturer_id").
		Where(sb.Equal("p.id", id)).
		GroupBy("p.id, pcl.name, pf.name, m.name")

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	var categoryIds []int64
	var categories []string
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&product.Id,
		pq.Array(&categoryIds),
		pq.Array(&categories),
		&product.Classification,
		&product.Form,
		&product.Manufacturer,
		&product.Name,
		&product.GenericName,
		&product.Stock,
		&product.Description,
		&product.Usage,
		&product.UnitInPack,
		&product.SellingUnit,
		&product.Weight,
		&product.Height,
		&product.Length,
		&product.Width,
		&product.ImageLink,
		&product.ClassificationId,
		&product.FormId,
		&product.ManufacturerId,
	)
	if err != nil {
		return nil, err
	}

	product.CategoryIds = categoryIds
	product.Categories = categories

	return &product, nil
}

func getProductExtras(r *ProductRepoImpl, ctx context.Context, table string) ([]*entity.ProductExtra, error) {
	var extras []*entity.ProductExtra

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id, name").
		From(table).
		Where(
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var extra = &entity.ProductExtra{}
		err := rows.Scan(
			&extra.Id,
			&extra.Name,
		)
		if err != nil {
			return nil, err
		}

		extras = append(extras, extra)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return extras, nil
}

func (r *ProductRepoImpl) GetProductClassifications(ctx context.Context) ([]*entity.ProductExtra, error) {
	return getProductExtras(r, ctx, "product_classifications")
}

func (r *ProductRepoImpl) GetProductForms(ctx context.Context) ([]*entity.ProductExtra, error) {
	return getProductExtras(r, ctx, "product_forms")
}

func (r *ProductRepoImpl) GetProductManufacturers(ctx context.Context) ([]*entity.ProductExtra, error) {
	return getProductExtras(r, ctx, "manufacturers")
}

func (r *ProductRepoImpl) GetProductByIdBulk(ctx context.Context, ids []int64) ([]*entity.Product, error) {
	products := make([]*entity.Product, len(ids))

	whereClause := make([]string, len(ids))
	for _, id := range ids {
		whereClause = append(whereClause, fmt.Sprintf("p.id=%d", id))
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id", "p.name", "p.image").
		From("products p").
		Where(
			sb.Or(whereClause...),
			sb.IsNull("p.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Name, &product.ImageLink)
		if err != nil {
			return nil, err
		}
		products[i] = &product
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepoImpl) GetMostBoughtProductIdsToday(ctx context.Context) ([]int64, error) {
	var productIds []int64
	currentTime := time.Now()

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id").
		From("products p").
		Join("catalogs c", "p.id = c.product_id").
		Join("order_items oi", "c.id = oi.catalog_id").
		Where(
			sb.GTE("oi.created_at", currentTime.Format("2006-01-02")),
			sb.LT("oi.created_at", currentTime.Add(24*time.Hour).Format("2006-01-02")),
			sb.IsNull("oi.deleted_at"),
		).
		GroupBy("p.id").
		OrderBy("SUM(oi.quantity) DESC").
		Limit(20)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var productId int64
		err := rows.Scan(&productId)
		if err != nil {
			return nil, err
		}
		productIds = append(productIds, productId)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return productIds, nil
}

func (r *ProductRepoImpl) GetMostBoughtProductIdsAllTime(ctx context.Context) ([]int64, error) {
	var productIds []int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("p.id").
		From("products p").
		Join("catalogs c", "p.id = c.product_id").
		Join("order_items oi", "c.id = oi.catalog_id").
		Where(
			sb.IsNull("oi.deleted_at"),
		).
		GroupBy("p.id").
		OrderBy("SUM(oi.quantity) DESC").
		Limit(20)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var productId int64
		err := rows.Scan(&productId)
		if err != nil {
			return nil, err
		}
		productIds = append(productIds, productId)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return productIds, nil
}
