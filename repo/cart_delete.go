package repo

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
)

func (r *CartRepoImpl) DeleteCartItem(ctx context.Context, cartItem *entity.CartItem) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("cart_items").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("user_id", cartItem.UserId),
			ub.Equal("product_id", cartItem.ProductId),
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

func (r *CartRepoImpl) DeleteCartItemByUserId(ctx context.Context, userId int64) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("cart_items").
		Set(
			ub.Assign("deleted_at", "NOW()"),
			ub.Assign("updated_at", "NOW()"),
		).
		Where(
			ub.Equal("user_id", userId),
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
