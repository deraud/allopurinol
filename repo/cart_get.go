package repo

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
)

func (r *CartRepoImpl) GetCartItems(ctx context.Context, userId int64) ([]*entity.CartItem, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("product_id", "quantity").
		From("cart_items").
		Where(
			sb.Equal("user_id", userId),
			sb.IsNull("deleted_at"),
		).
		OrderBy("product_id")

	query, args := sb.BuildWithFlavor(constant.DATABASE_FLAVOR)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*entity.CartItem

	for i := 0; rows.Next(); i++ {
		var item = &entity.CartItem{}
		err := rows.Scan(
			&item.ProductId,
			&item.Quantity,
		)
		if err != nil {
			return nil, err
		}
		item.UserId = userId
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return items, nil
}
