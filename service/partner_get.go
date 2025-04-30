package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PartnerServiceImpl) GetPartnersService(ctx context.Context, options *entity.PartnerOptions) ([]*entity.Partner, *entity.PartnerOptions, error) {
	count, err := s.partnerRepo.GetPartnersCount(ctx, options)
	if err != nil {
		return nil, nil, err
	}
	options.TotalRows = count

	partners, err := s.partnerRepo.GetPartners(ctx, options)
	if err != nil {
		return nil, nil, err
	}

	return partners, options, nil
}

func (s *PartnerServiceImpl) GetPartnerService(ctx context.Context, id int64) (*entity.Partner, error) {
	exist, err := s.partnerRepo.IsPartnerIdExist(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrPartnerNotFound
	}

	partner, err := s.partnerRepo.GetPartner(ctx, id)
	if err != nil {
		return nil, err
	}

	return partner, nil
}
