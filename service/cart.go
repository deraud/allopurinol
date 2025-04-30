package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type CartServiceItf interface {
	IncrementCartItemService(ctx context.Context, cartItem *entity.CartItem, credId int64) error
	DecrementCartItemService(ctx context.Context, cartItem *entity.CartItem, credId int64) error
	UpdateCartItemService(ctx context.Context, cartItem *entity.CartItem, value int, credId int64) error
	DeleteCartItemService(ctx context.Context, cartItem *entity.CartItem, credId int64) error
	GetCartItemsService(ctx context.Context, credId int64) ([]*entity.CartItem, error)
}

type CartServiceImpl struct {
	transactorRepo repo.TransactorItf
	userRepo       repo.UserRepoItf
	productRepo    repo.ProductRepoItf
	cartRepo       repo.CartRepoItf
}

func NewCartService(transactorRepo repo.TransactorItf, userRepo repo.UserRepoItf, productRepo repo.ProductRepoItf, cartRepo repo.CartRepoItf) CartServiceItf {
	return &CartServiceImpl{
		transactorRepo: transactorRepo,
		userRepo:       userRepo,
		productRepo:    productRepo,
		cartRepo:       cartRepo,
	}
}
