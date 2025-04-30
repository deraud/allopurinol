package repo

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
)

func (r *ProductRepoImpl) UpdateProductCategory(ctx context.Context, pc *entity.ProductCategory) (*entity.ProductCategory, error) {
	cat := *pc

	ub := sqlbuilder.NewUpdateBuilder()
	values := []string{ub.Assign("updated_at", "NOW()")}
	if pc.Name != "" {
		values = append(values, ub.Assign("name", pc.Name))
	}
	ub.Update("product_categories").
		Set(values...).
		Where(ub.Equal("id", pc.Id)).
		SQL("RETURNING name")

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&cat.Name,
	)
	if err != nil {
		return nil, err
	}

	return &cat, nil
}

func (r *ProductRepoImpl) UpdateProduct(ctx context.Context, p *entity.Product) (*entity.Product, error) {
	product := *p

	ub := sqlbuilder.NewUpdateBuilder()
	values := []string{
		ub.Assign("product_classification_id", p.ClassificationId),
		ub.Assign("product_form_id", p.FormId),
		ub.Assign("manufacturer_id", p.ManufacturerId),
		ub.Assign("name", p.Name),
		ub.Assign("generic_name", p.GenericName),
		ub.Assign("description", p.Description),
		ub.Assign("weight", p.Weight),
		ub.Assign("height", p.Height),
		ub.Assign("length", p.Length),
		ub.Assign("width", p.Width),
		ub.Assign("updated_at", "NOW()"),
	}
	if p.ImageLink != "" {
		values = append(values, ub.Assign("image", p.ImageLink))
	}
	ub.Update("products").
		Set(values...).
		Where(ub.Equal("id", p.Id)).
		SQL("RETURNING image")

	query, args := ub.BuildWithFlavor(constant.DATABASE_FLAVOR)
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&p.ImageLink)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
