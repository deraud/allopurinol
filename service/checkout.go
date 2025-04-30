package service

import (
	"context"
	"healthcare-allopurinol/bridge"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type CheckoutServiceItf interface {
	GetShippingCosts(ctx context.Context, shipping *entity.Shipping, credId, addressId int64) (*entity.Shipping, error)
}

type CheckoutServiceImpl struct {
	transactorRepo   repo.TransactorItf
	rajaOngkirBridge bridge.ROBridgeItf
	shippingRepo     repo.ShippingRepoItf
	rajaOngkirRepo   repo.RORepoItf
	pharmacyRepo     repo.PharmacyRepoItf
	addressRepo      repo.AddressRepoItf
	userRepo         repo.UserRepoItf
	redisRepo        repo.RedisRepoItf
}

func NewCheckoutService(transactorRepo repo.TransactorItf, rajaOngkirBridge bridge.ROBridgeItf, shippingRepo repo.ShippingRepoItf, rajaOngkirRepo repo.RORepoItf, pharmacyRepo repo.PharmacyRepoItf, addressRepo repo.AddressRepoItf, userRepo repo.UserRepoItf, redisRepo repo.RedisRepoItf) CheckoutServiceItf {
	return &CheckoutServiceImpl{
		transactorRepo:   transactorRepo,
		rajaOngkirBridge: rajaOngkirBridge,
		shippingRepo:     shippingRepo,
		rajaOngkirRepo:   rajaOngkirRepo,
		pharmacyRepo:     pharmacyRepo,
		addressRepo:      addressRepo,
		userRepo:         userRepo,
		redisRepo:        redisRepo,
	}
}
