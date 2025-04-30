package service

import (
	"context"
	"fmt"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PharmacyServiceImpl) CreatePharmacyService(ctx context.Context, p *entity.Pharmacy) (*entity.Pharmacy, error) {
	exist, err := s.partnerRepo.IsIdExist(ctx, p.PartnerId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrPartnerNotFound
	}

	ids, err := s.shippingRepo.IsLogisticPartnerIdExistBulk(ctx, p.LogisticPartnerIds...)
	if len(ids) == 0 {
		return nil, sentinel.NewBadRequestError("logistic_partners[]", fmt.Sprintf("logistic_partner %v doesn't exist", ids))
	}
	if err != nil {
		return nil, err
	}

	if len(p.PharmacistIds) > 0 {
		assignedPharmacists, err := s.pharmacistRepo.IsAssignedBulk(ctx, p.PharmacistIds...)
		if assignedPharmacists != nil {
			return nil, sentinel.NewBadRequestError("pharmacists[]", fmt.Sprintf("pharmacist [%v] had been assigned to another pharmacy", assignedPharmacists))
		}
		if err != nil {
			return nil, err
		}
	}

	var pharmacy *entity.Pharmacy
	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		pharmacy, err = s.pharmacyRepo.InsertPharmacy(txCtx, p)
		if err != nil {
			return err
		}
		p.Address.PharmacyId = &pharmacy.Id

		pharmacy.Address, err = s.addressRepo.InsertAddress(txCtx, p.Address)
		if err != nil {
			return err
		}

		err = s.pharmacistRepo.AssignPharmacistBulk(txCtx, pharmacy.Id, p.PharmacistIds...)
		if err != nil {
			return err
		}

		err = s.shippingRepo.InsertShippingMethod(txCtx, pharmacy.Id, p.LogisticPartnerIds...)
		if err != nil {
			return err
		}
		pharmacy.PharmacistIds = p.PharmacistIds
		pharmacy.LogisticPartnerIds = p.LogisticPartnerIds
		return nil
	})
	if err != nil {
		return nil, err
	}

	return pharmacy, nil
}
