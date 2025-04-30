package repo

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
)

func (r *ProductRepoImpl) DeleteProductCategory(ctx context.Context, id int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("product_categories").
		Set(
			ub.Assign("deleted_at", "NOW()"),
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

func (r *ProductRepoImpl) DeleteProductCategoryMaps(ctx context.Context, id int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("product_category_maps").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("product_id", id),
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

func (r *ProductRepoImpl) DeleteProductCarts(ctx context.Context, ids []int64) error {
	ub := sqlbuilder.NewUpdateBuilder()

	var ors []string
	for _, id := range ids {
		ors = append(ors, ub.Equal("catalog_id", id))
	}

	ub.Update("cart_items").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Or(ors...),
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

func (r *ProductRepoImpl) DeleteProductCatalog(ctx context.Context, id int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("catalogs").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("product_id", id),
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

func (r *ProductRepoImpl) DeleteProduct(ctx context.Context, id int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("products").
		Set(
			ub.Assign("deleted_at", "NOW()"),
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
