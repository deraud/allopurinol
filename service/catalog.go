package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type CatalogServiceItf interface {
	CreateCatalogService(ctx context.Context, pharmacistCredId int64, c *entity.Catalog) (*entity.Catalog, error)
	UpdateCatalogService(ctx context.Context, pharmacistCredId int64, c *entity.Catalog) (*entity.Catalog, error)
	DeleteCatalogService(ctx context.Context, pharmacistCredId int64, id int64) error
	GetCatalogsService(ctx context.Context, options *entity.CatalogOptions, pharmacistCredId int64) ([]*entity.Catalog, *entity.CatalogOptions, error)
	GetCatalogService(ctx context.Context, id int64, pharmacistCredId int64) (*entity.Catalog, error)
	GetAvailableCatalogsService(ctx context.Context, options *entity.AvailableCatalogOptions, userCredId int64) ([]*entity.Catalog, *entity.AvailableCatalogOptions, error)
	GetAvailableCatalogService(ctx context.Context, id int64) (*entity.Catalog, error)
	GetCheckoutCatalogsService(ctx context.Context, userCredId int64, addressId int64) ([]*entity.Pharmacy, map[int64][]*entity.Catalog, []*entity.Catalog, error)
	GetMostBoughtCatalogsService(ctx context.Context) ([]*entity.Catalog, error)
}

type CatalogServiceImpl struct {
	productRepo    repo.ProductRepoItf
	catalogRepo    repo.CatalogRepoItf
	pharmacistRepo repo.PharmacistRepoItf
	pharmacyRepo   repo.PharmacyRepoItf
	addressRepo    repo.AddressRepoItf
	userRepo       repo.UserRepoItf
	cartRepo       repo.CartRepoItf
	shippingRepo   repo.ShippingRepoItf
}

func NewCatalogService(productRepo repo.ProductRepoItf, catalogRepo repo.CatalogRepoItf, pharmacistRepo repo.PharmacistRepoItf, pharmacyRepo repo.PharmacyRepoItf, addressRepo repo.AddressRepoItf, userRepo repo.UserRepoItf, cartRepo repo.CartRepoItf, shippingRepo repo.ShippingRepoItf) CatalogServiceItf {
	return &CatalogServiceImpl{
		productRepo:    productRepo,
		catalogRepo:    catalogRepo,
		pharmacistRepo: pharmacistRepo,
		pharmacyRepo:   pharmacyRepo,
		addressRepo:    addressRepo,
		userRepo:       userRepo,
		cartRepo:       cartRepo,
		shippingRepo:   shippingRepo,
	}
}
