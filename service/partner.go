package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type PartnerServiceItf interface {
	CreatePartnerService(ctx context.Context, p *entity.Partner) (*entity.Partner, error)
	UpdatePartnerService(ctx context.Context, p *entity.Partner) (*entity.Partner, error)
	DeletePartnerService(ctx context.Context, id int64) error
	GetPartnersService(ctx context.Context, options *entity.PartnerOptions) ([]*entity.Partner, *entity.PartnerOptions, error)
	GetPartnerService(ctx context.Context, id int64) (*entity.Partner, error)
}

type PartnerServiceImpl struct {
	partnerRepo repo.PartnerRepoItf
}

func NewPartnerService(partnerRepo repo.PartnerRepoItf) PartnerServiceItf {
	return &PartnerServiceImpl{
		partnerRepo: partnerRepo,
	}
}
