package repo

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
)

func (r *CartRepoImpl) InsertCartItem(ctx context.Context, cartItem *entity.CartItem, quantity int) error {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("cart_items").
		Cols("user_id", "product_id", "quantity").
		Values(cartItem.UserId, cartItem.ProductId, quantity)

	tx := extractTx(ctx)
	query, args := ib.BuildWithFlavor(constant.DATABASE_FLAVOR)
	_, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
