package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type OrderServiceItf interface {
	CreateOrderGroupService(ctx context.Context, pharmacies []*entity.Pharmacy, catalogs map[int64][]*entity.Catalog, details *entity.CheckoutDetails, userCredId int64) (*entity.OrderGroup, error)
	UpdateOrderGroupService(ctx context.Context, o *entity.OrderGroup, userCredId int64) (*entity.OrderGroup, error)
	GetPendingOrdersService(ctx context.Context, options *entity.PendingOrderOptions, credId int64) ([]*entity.PendingOrderGroup, error)
	UpdateUserOrderCanceledService(ctx context.Context, o *entity.Order, userCredId int64) (*entity.Order, error)
	GetPharmacyOrdersService(ctx context.Context, options *entity.PharmacyOrderOptions, pharmacistCredId int64) ([]*entity.Order, *entity.PharmacyOrderOptions, error)
	GetPharmacyOrderService(ctx context.Context, id int64, pharmacistCredId int64) (*entity.Order, error)
	UpdatePharmacyOrderShippedService(ctx context.Context, o *entity.Order, pharmacistCredId int64) (*entity.Order, error)
	UpdatePharmacyOrderCanceledService(ctx context.Context, o *entity.Order, pharmacistCredId int64) (*entity.Order, error)
	GetOrdersService(ctx context.Context, options *entity.OrderOptions) ([]*entity.Order, *entity.OrderOptions, error)
	GetOrderService(ctx context.Context, id int64) (*entity.Order, error)
	UpdateUserOrderConfirmedService(ctx context.Context, o *entity.Order, userCredId int64) (*entity.Order, error)
	GetUserOrdersService(ctx context.Context, options *entity.UserOrderOptions, userCredId int64) ([]*entity.Order, *entity.UserOrderOptions, error)
}

type OrderServiceImpl struct {
	transactorRepo repo.TransactorItf
	redisRepo      repo.RedisRepoItf
	pharmacistRepo repo.PharmacistRepoItf
	orderRepo      repo.OrderRepoItf
	userRepo       repo.UserRepoItf
	catalogRepo    repo.CatalogRepoItf
	cartRepo       repo.CartRepoItf
}

func NewOrderService(transactorRepo repo.TransactorItf, redisRepo repo.RedisRepoItf, orderRepo repo.OrderRepoItf, userRepo repo.UserRepoItf, catalogRepo repo.CatalogRepoItf, cartRepo repo.CartRepoItf, pharmacistRepo repo.PharmacistRepoItf) OrderServiceItf {
	return &OrderServiceImpl{
		transactorRepo: transactorRepo,
		redisRepo:      redisRepo,
		pharmacistRepo: pharmacistRepo,
		orderRepo:      orderRepo,
		userRepo:       userRepo,
		catalogRepo:    catalogRepo,
		cartRepo:       cartRepo,
	}
}
