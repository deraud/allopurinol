package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PharmacyServiceImpl) GetPharmaciesService(ctx context.Context, options *entity.PharmacyOptions) ([]*entity.Pharmacy, *entity.PharmacyOptions, error) {
	count, err := s.pharmacyRepo.GetPharmaciesCount(ctx, options)
	if err != nil {
		return nil, nil, err
	}
	options.TotalRows = count

	pharmacies, err := s.pharmacyRepo.GetPharmacies(ctx, options)
	if err != nil {
		return nil, nil, err
	}

	return pharmacies, options, nil
}

func (s *PharmacyServiceImpl) GetPharmacyService(ctx context.Context, id int64) (*entity.Pharmacy, error) {
	exist, err := s.pharmacyRepo.IsIdExist(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrPharmacyNotFound
	}

	pharmacy, err := s.pharmacyRepo.GetPharmacy(ctx, id)
	if err != nil {
		return nil, err
	}

	pharmacy.Pharmacists, err = s.pharmacistRepo.GetPharmacistsByPharmacy(ctx, id)
	if err != nil {
		return nil, err
	}

	pharmacy.LogisticPartners, err = s.shippingRepo.GetLogisticPartnersByPharmacy(ctx, id)
	if err != nil {
		return nil, err
	}

	return pharmacy, nil
}

func (s *PharmacyServiceImpl) GetPharmacyFromPharmacistService(ctx context.Context, pharmacistCredId int64) (*entity.Pharmacy, error) {
	pharmacistId, err := s.pharmacistRepo.GetIdByCredId(ctx, pharmacistCredId)
	if err != nil {
		return nil, err
	}

	assigned, err := s.pharmacistRepo.IsAssigned(ctx, pharmacistId)
	if err != nil {
		return nil, err
	}
	if !assigned {
		return nil, sentinel.ErrPharmacistNotAssigned
	}

	pharmacyId, err := s.pharmacistRepo.GetPharmacyIdByCredId(ctx, pharmacistCredId)
	if err != nil {
		return nil, err
	}

	pharmacy, err := s.pharmacyRepo.GetPharmacy(ctx, pharmacyId)
	if err != nil {
		return nil, err
	}

	pharmacy.Pharmacists, err = s.pharmacistRepo.GetPharmacistsByPharmacy(ctx, pharmacyId)
	if err != nil {
		return nil, err
	}

	pharmacy.LogisticPartners, err = s.shippingRepo.GetLogisticPartnersByPharmacy(ctx, pharmacyId)
	if err != nil {
		return nil, err
	}

	return pharmacy, nil
}
