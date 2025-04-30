package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *CatalogRepoImpl) UpdateCatalog(ctx context.Context, c *entity.Catalog) (*entity.Catalog, error) {
	catalog := *c

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("catalogs").
		Set(
			ub.Assign("stock", c.Stock),
			ub.Assign("price", c.Price),
			ub.Assign("is_active", c.IsActive),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("id", c.Id),
			ub.IsNull("deleted_at"),
		).
		SQL("RETURNING pharmacy_id, product_id")

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&catalog.PharmacyId, &catalog.ProductId)
	if err != nil {
		return nil, err
	}

	return &catalog, nil
}

func (r *CatalogRepoImpl) UpdateStock(ctx context.Context, id int64, value int) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("catalogs").
		Set(
			ub.Add("stock", value),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("id", id),
			ub.IsNull("deleted_at"),
		)

	tx := extractTx(ctx)
	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
