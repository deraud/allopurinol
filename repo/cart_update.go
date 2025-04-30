package repo

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
)

func updateCartItemOne(ctx context.Context, cartItem *entity.CartItem, increment bool) error {
	ub := sqlbuilder.NewUpdateBuilder()
	var change int
	if increment {
		change = 1
	} else {
		change = -1
	}
	ub.Update("cart_items").
		Set(
			ub.Add("quantity", change),
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

func (r *CartRepoImpl) UpdateCartItemIncrement(ctx context.Context, cartItem *entity.CartItem) error {
	return updateCartItemOne(ctx, cartItem, true)
}

func (r *CartRepoImpl) UpdateCartItemDecrement(ctx context.Context, cartItem *entity.CartItem) error {
	return updateCartItemOne(ctx, cartItem, false)
}

func (r *CartRepoImpl) UpdateCartItem(ctx context.Context, cartItem *entity.CartItem, value int) error {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("cart_items").
		Set(
			ub.Assign("quantity", value),
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
