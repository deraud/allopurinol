package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PharmacistServiceImpl) UpdatePharmacistService(ctx context.Context, p *entity.Pharmacist) (*entity.Pharmacist, error) {
	exist, err := s.pharmacistRepo.IsIdExist(ctx, p.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrPharmacistNotFound
	}

	id, err := s.pharmacistRepo.GetIdByPhoneNumber(ctx, p.Id, p.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if id != 0 {
		return nil, sentinel.ErrPhoneNumberRegistered
	}

	if p.PharmacyId != nil {
		exist, err := s.pharmacyRepo.IsIdExist(ctx, *p.PharmacyId)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, sentinel.ErrPharmacyNotFound
		}
	} else {
		pharmacist, _, err := s.pharmacistRepo.GetPharmacist(ctx, p.Id)
		if err != nil {
			return nil, err
		}
		if pharmacist.PharmacyId != nil {
			count, err := s.pharmacyRepo.GetPharmacistCount(ctx, *pharmacist.PharmacyId)
			if err != nil {
				return nil, err
			}
			if count <= 1 {
				return nil, sentinel.ErrPharmacyNoPharmacist
			}
		}
	}

	pharmacist, err := s.pharmacistRepo.UpdatePharmacist(ctx, p)
	if err != nil {
		return nil, err
	}

	return pharmacist, nil
}
