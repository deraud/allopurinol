package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PharmacistServiceImpl) GetPharmacistsService(ctx context.Context, options *entity.PharmacistOptions) ([]*entity.Pharmacist, []string, *entity.PharmacistOptions, error) {
	count, err := s.pharmacistRepo.GetPharmacistsCount(ctx, options)
	if err != nil {
		return nil, nil, nil, err
	}
	options.TotalRows = count

	pharmacists, emails, err := s.pharmacistRepo.GetPharmacists(ctx, options)
	if err != nil {
		return nil, nil, nil, err
	}

	return pharmacists, emails, options, nil
}

func (s *PharmacistServiceImpl) GetPharmacistService(ctx context.Context, id int64) (*entity.Pharmacist, string, error) {
	exist, err := s.pharmacistRepo.IsIdExist(ctx, id)
	if err != nil {
		return nil, "", err
	}
	if !exist {
		return nil, "", sentinel.ErrPharmacistNotFound
	}

	pharmacist, email, err := s.pharmacistRepo.GetPharmacist(ctx, id)
	if err != nil {
		return nil, "", err
	}

	return pharmacist, email, nil
}
