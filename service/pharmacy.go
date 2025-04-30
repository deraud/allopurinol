package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type PharmacyServiceItf interface {
	GetPharmaciesService(ctx context.Context, options *entity.PharmacyOptions) ([]*entity.Pharmacy, *entity.PharmacyOptions, error)
	CreatePharmacyService(ctx context.Context, p *entity.Pharmacy) (*entity.Pharmacy, error)
	UpdatePharmacyService(ctx context.Context, p *entity.Pharmacy) (*entity.Pharmacy, error)
	GetPharmacyService(ctx context.Context, id int64) (*entity.Pharmacy, error)
	DeletePharmacyService(ctx context.Context, id int64) error
	UpdatePharmacyFromPharmacistService(ctx context.Context, pharmacistId int64, p *entity.Pharmacy) (*entity.Pharmacy, error)
	GetPharmacyFromPharmacistService(ctx context.Context, pharmacistCredId int64) (*entity.Pharmacy, error)
}

type PharmacyServiceImpl struct {
	transactorRepo repo.TransactorItf
	pharmacyRepo   repo.PharmacyRepoItf
	pharmacistRepo repo.PharmacistRepoItf
	partnerRepo    repo.PartnerRepoItf
	addressRepo    repo.AddressRepoItf
	shippingRepo   repo.ShippingRepoItf
}

func NewPharmacyService(transactorRepo repo.TransactorItf, pharmacyRepo repo.PharmacyRepoItf, pharmacistRepo repo.PharmacistRepoItf, partnerRepo repo.PartnerRepoItf, addressRepo repo.AddressRepoItf, shippingRepo repo.ShippingRepoItf) PharmacyServiceItf {
	return &PharmacyServiceImpl{
		transactorRepo: transactorRepo,
		pharmacyRepo:   pharmacyRepo,
		pharmacistRepo: pharmacistRepo,
		partnerRepo:    partnerRepo,
		addressRepo:    addressRepo,
		shippingRepo:   shippingRepo,
	}
}
