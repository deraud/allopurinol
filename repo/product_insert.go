package repo

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"

	"github.com/huandu/go-sqlbuilder"
)

func (r *ProductRepoImpl) InsertProductCategory(ctx context.Context, pc *entity.ProductCategory) (*entity.ProductCategory, error) {
	cat := *pc

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("product_categories").
		Cols("name").
		Values(pc.Name).
		SQL("RETURNING id")

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&cat.Id)
	if err != nil {
		return nil, err
	}

	return &cat, nil
}

func (r *ProductRepoImpl) InsertProduct(ctx context.Context, p *entity.Product) (*entity.Product, error) {
	product := *p

	ib := sqlbuilder.NewInsertBuilder()
	cols := []string{"product_classification_id", "product_form_id", "manufacturer_id", "name", "generic_name",
		"description", "weight", "height", "length", "width", "image"}
	values := []any{p.ClassificationId, p.FormId, p.ManufacturerId, p.Name, p.GenericName, p.Description, p.Weight,
		p.Height, p.Length, p.Width, p.ImageLink}

	if p.UnitInPack != nil {
		cols = append(cols, "unit_in_pack")
		values = append(values, p.UnitInPack)
	}
	if p.SellingUnit != nil {
		cols = append(cols, "selling_unit")
		values = append(values, p.SellingUnit)
	}
	if !p.IsActive {
		cols = append(cols, "is_active")
		values = append(values, p.IsActive)
	}

	ib.InsertInto("products").
		Cols(cols...).
		Values(values...).
		SQL("RETURNING id")

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := tx.QueryRowContext(ctx, query, args...).Scan(&product.Id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepoImpl) InsertProductCategoryMaps(ctx context.Context, productId int64, categoryIds []int64) error {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("product_category_maps").
		Cols("product_id", "product_category_id")

	for _, id := range categoryIds {
		ib.Values(productId, id)
	}

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
