package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PartnerServiceImpl) CreatePartnerService(ctx context.Context, p *entity.Partner) (*entity.Partner, error) {
	valid := s.partnerRepo.IsYearFoundedValid(ctx, *p)
	if !valid {
		return nil, sentinel.ErrMoreThanCurrentYear
	}
	partner, err := s.partnerRepo.InsertPartner(ctx, p)
	if err != nil {
		return nil, err
	}

	return partner, nil
}
