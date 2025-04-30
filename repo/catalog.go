package repo

import (
	"context"
	"database/sql"
	"errors"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

type CatalogRepoItf interface {
	IsIdExist(ctx context.Context, id int64) (bool, error)
	GetPharmacyId(ctx context.Context, id int64) (int64, error)
	InsertCatalog(ctx context.Context, c *entity.Catalog) (*entity.Catalog, error)
	UpdateCatalog(ctx context.Context, c *entity.Catalog) (*entity.Catalog, error)
	IsCatalogExist(ctx context.Context, pharmacyId int64, productId int64) (bool, error)
	DeleteCatalog(ctx context.Context, id int64) error
	GetCatalogs(ctx context.Context, options *entity.CatalogOptions, pharmacyId int64) ([]*entity.Catalog, error)
	GetCatalogsCount(ctx context.Context, options *entity.CatalogOptions, pharmacyId int64) (int, error)
	GetCatalog(ctx context.Context, id int64) (*entity.Catalog, error)
	GetAvailableCatalogs(ctx context.Context, options *entity.AvailableCatalogOptions, userAddress *entity.Address) ([]*entity.Catalog, int, error)
	GetAvailableCatalogsCount(ctx context.Context, options *entity.AvailableCatalogOptions, address *entity.Address) (int, error)
	GetAvailableCatalog(ctx context.Context, id int64) (*entity.Catalog, error)
	GetCheckoutCatalogs(ctx context.Context, cartItems []*entity.CartItem, userAddress *entity.Address) ([]*entity.Catalog, error)
	UpdateStock(ctx context.Context, id int64, value int) error
	IsOrdered(ctx context.Context, id int64) (bool, error)
	GetCatalogsByProductIds(ctx context.Context, productIds []int64) ([]*entity.Catalog, error)
}

type CatalogRepoImpl struct {
	db *sql.DB
}

func NewCatalogRepo(database *sql.DB) CatalogRepoItf {
	return &CatalogRepoImpl{
		db: database,
	}
}

func (r *CatalogRepoImpl) IsIdExist(ctx context.Context, id int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("catalogs").
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

func (r *CatalogRepoImpl) GetPharmacyId(ctx context.Context, id int64) (int64, error) {
	var pharmacyId int64

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("pharmacy_id").
		From("catalogs").
		Where(
			sb.Equal("id", id),
			sb.IsNull("deleted_at"),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&pharmacyId)
	if err != nil {
		return 0, err
	}

	return pharmacyId, nil
}

func (r *CatalogRepoImpl) IsCatalogExist(ctx context.Context, pharmacyId int64, productId int64) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("catalogs").
		Where(
			sb.Equal("pharmacy_id", pharmacyId),
			sb.Equal("product_id", productId),
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

func (r *CatalogRepoImpl) IsOrdered(ctx context.Context, id int64) (bool, error) {
	var count int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(oi.id)").
		From("catalogs c").
		Join("order_items oi", "c.id = oi.catalog_id").
		Where(
			sb.Equal("c.id", id),
		)

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}
