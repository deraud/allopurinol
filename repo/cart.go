package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/entity"
)

type CartRepoItf interface {
	UpdateCartItemIncrement(ctx context.Context, cartItem *entity.CartItem) error
	UpdateCartItemDecrement(ctx context.Context, cartItem *entity.CartItem) error
	UpdateCartItem(ctx context.Context, cartItem *entity.CartItem, value int) error
	IsCartItemExist(ctx context.Context, cartItem *entity.CartItem) (bool, error)
	InsertCartItem(ctx context.Context, cartItem *entity.CartItem, quantity int) error
	IsCartItemOnOne(ctx context.Context, cartItem *entity.CartItem) (bool, error)
	DeleteCartItem(ctx context.Context, cartItem *entity.CartItem) error
	GetCartItems(ctx context.Context, userId int64) ([]*entity.CartItem, error)
	DeleteCartItemByUserId(ctx context.Context, userId int64) error
}

type CartRepoImpl struct {
	db *sql.DB
}

func NewCartRepo(database *sql.DB) CartRepoItf {
	return &CartRepoImpl{
		db: database,
	}
}

func (r *CartRepoImpl) IsCartItemExist(ctx context.Context, cartItem *entity.CartItem) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("cart_items").
		Where(
			sb.Equal("user_id", cartItem.UserId),
			sb.Equal("product_id", cartItem.ProductId),
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

func (r *CartRepoImpl) IsCartItemOnOne(ctx context.Context, cartItem *entity.CartItem) (bool, error) {
	var exist bool

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("1").
		From("cart_items").
		Where(
			sb.Equal("user_id", cartItem.UserId),
			sb.Equal("product_id", cartItem.ProductId),
			sb.Equal("quantity", 1),
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
