package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
	"strings"
)

type ProductRepoItf interface {
	IsCategoryExist(ctx context.Context, category string) (bool, error)
	IsIdExist(ctx context.Context, id int64) (bool, error)
	IsCategoryInProductExist(ctx context.Context, id int64) (bool, error)
	IsProductRequirementsExist(ctx context.Context, product *entity.Product) (bool, error)
	IsProductExist(ctx context.Context, product *entity.Product) (bool, error)
	IsProductIdExist(ctx context.Context, id int64) (bool, error)
	IsProductBoughtExist(ctx context.Context, id int64) (bool, error)
	InsertProductCategory(ctx context.Context, pc *entity.ProductCategory) (*entity.ProductCategory, error)
	UpdateProductCategory(ctx context.Context, pc *entity.ProductCategory) (*entity.ProductCategory, error)
	DeleteProductCategory(ctx context.Context, id int64) error
	GetProductCategories(ctx context.Context, options *entity.ProductCategoryOptions) ([]*entity.ProductCategory, error)
	GetProductCategoriesCount(ctx context.Context, options *entity.ProductCategoryOptions) (int, error)
	InsertProduct(ctx context.Context, p *entity.Product) (*entity.Product, error)
	InsertProductCategoryMaps(ctx context.Context, productId int64, categoryIds []int64) error
	GetProducts(ctx context.Context, options *entity.ProductOptions) ([]*entity.Product, error)
	GetProductsCount(ctx context.Context, options *entity.ProductOptions) (int, error)
	DeleteProductCategoryMaps(ctx context.Context, id int64) error
	DeleteProduct(ctx context.Context, id int64) error
	GetProductCatalogIds(ctx context.Context, id int64) ([]int64, error)
	DeleteProductCarts(ctx context.Context, ids []int64) error
	DeleteProductCatalog(ctx context.Context, id int64) error
	GetProductDetail(ctx context.Context, id int64) (*entity.Product, error)
	UpdateProduct(ctx context.Context, p *entity.Product) (*entity.Product, error)
	GetProductClassifications(ctx context.Context) ([]*entity.ProductExtra, error)
	GetProductForms(ctx context.Context) ([]*entity.ProductExtra, error)
	GetProductManufacturers(ctx context.Context) ([]*entity.ProductExtra, error)
	GetProductByIdBulk(ctx context.Context, ids []int64) ([]*entity.Product, error)
	GetMostBoughtProductIdsToday(ctx context.Context) ([]int64, error)
	GetMostBoughtProductIdsAllTime(ctx context.Context) ([]int64, error)
}

type ProductRepoImpl struct {
	db *sql.DB
}

func NewProductRepo(database *sql.DB) ProductRepoItf {
	return &ProductRepoImpl{
		db: database,
	}
}

func isCustomExist(r *ProductRepoImpl, ctx context.Context, table string, equal string, compare any) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From(table).
		Where(
			sb.Equal(equal, compare),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *ProductRepoImpl) IsCategoryExist(ctx context.Context, category string) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("product_categories").
		Where(
			sb.Equal("LOWER(name)", strings.ToLower(category)),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *ProductRepoImpl) IsProductIdExist(ctx context.Context, id int64) (bool, error) {
	return isCustomExist(r, ctx, "products", "id", id)
}

func (r *ProductRepoImpl) IsIdExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("product_categories").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *ProductRepoImpl) IsCategoryInProductExist(ctx context.Context, id int64) (bool, error) {
	return isCustomExist(r, ctx, "product_category_maps", "product_category_id", id)
}

func (r *ProductRepoImpl) IsProductRequirementsExist(ctx context.Context, product *entity.Product) (bool, error) {
	classExist, err := isCustomExist(r, ctx, "product_classifications", "id", product.ClassificationId)
	if err != nil {
		return false, err
	}
	if !classExist {
		return false, nil
	}

	formExist, err := isCustomExist(r, ctx, "product_forms", "id", product.FormId)
	if err != nil {
		return false, err
	}
	if !formExist {
		return false, nil
	}

	manExist, err := isCustomExist(r, ctx, "manufacturers", "id", product.ManufacturerId)
	if err != nil {
		return false, err
	}
	if !manExist {
		return false, nil
	}

	for _, category := range product.CategoryIds {
		catExist, err := isCustomExist(r, ctx, "product_categories", "id", category)
		if err != nil {
			return false, err
		}
		if !catExist {
			return false, nil
		}
	}

	return true, nil
}

func (r *ProductRepoImpl) IsProductExist(ctx context.Context, product *entity.Product) (bool, error) {
	nameExist, err := isCustomExist(r, ctx, "products", "LOWER(name)", strings.ToLower(product.Name))
	if err != nil {
		return false, err
	}
	if !nameExist {
		return false, nil
	}

	genExist, err := isCustomExist(r, ctx, "products", "LOWER(generic_name)", strings.ToLower(product.GenericName))
	if err != nil {
		return false, err
	}
	if !genExist {
		return false, nil
	}

	manExist, err := isCustomExist(r, ctx, "products", "manufacturer_id", product.ManufacturerId)
	if err != nil {
		return false, err
	}
	if !manExist {
		return false, nil
	}

	return true, nil
}

func (r *ProductRepoImpl) IsProductBoughtExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("catalogs c").
		Join("order_items o", "o.catalog_id = c.id").
		Where(
			sb.Equal("c.product_id", id),
			sb.IsNull("c.deleted_at"),
			sb.IsNull("o.deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exist, nil
}
