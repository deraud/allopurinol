package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *CatalogRepoImpl) InsertCatalog(ctx context.Context, c *entity.Catalog) (*entity.Catalog, error) {
	catalog := *c

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("catalogs").
		Cols("pharmacy_id", "product_id", "stock", "price").
		Values(c.PharmacyId, c.ProductId, c.Stock, c.Price).
		SQL("RETURNING id, is_active")

	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&catalog.Id, &catalog.IsActive)
	if err != nil {
		return nil, err
	}

	return &catalog, nil
}
