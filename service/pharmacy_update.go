package service

import (
	"context"
	"fmt"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PharmacyServiceImpl) UpdatePharmacyService(ctx context.Context, p *entity.Pharmacy) (*entity.Pharmacy, error) {
	exist, err := s.pharmacyRepo.IsIdExist(ctx, p.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrPharmacyNotFound
	}

	ids, err := s.shippingRepo.IsLogisticPartnerIdExistBulk(ctx, p.LogisticPartnerIds...)
	if len(ids) == 0 {
		return nil, sentinel.NewBadRequestError("logistic_partners[]", fmt.Sprintf("logistic_partner %v doesn't exist", ids))
	}
	if err != nil {
		return nil, err
	}

	if len(p.PharmacistIds) > 0 {
		assignedPharmacists, err := s.pharmacistRepo.IsAssignedToPharmacyBulk(ctx, p.Id, p.PharmacistIds...)
		if assignedPharmacists != nil {
			return nil, sentinel.NewBadRequestError("pharmacists[]", fmt.Sprintf("pharmacist [%v] had been assigned to another pharmacy", assignedPharmacists))
		}
		if err != nil {
			return nil, err
		}
	}

	var pharmacy *entity.Pharmacy
	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		pharmacy, err = s.pharmacyRepo.UpdatePharmacy(txCtx, p)
		if err != nil {
			return err
		}
		p.Address.PharmacyId = &pharmacy.Id

		pharmacy.Address, err = s.addressRepo.UpdateAddressPharmacy(txCtx, p.Address)
		if err != nil {
			return err
		}

		err = s.pharmacistRepo.UnassignPharmacistFromPharmacy(txCtx, pharmacy.Id)
		if err != nil {
			return err
		}

		err = s.pharmacistRepo.AssignPharmacistBulk(txCtx, pharmacy.Id, p.PharmacistIds...)
		if err != nil {
			return err
		}

		err = s.shippingRepo.DeleteShippingMethodExceptIds(txCtx, pharmacy.Id, p.LogisticPartnerIds...)
		if err != nil {
			return err
		}

		for _, logisticPartnerId := range p.LogisticPartnerIds {
			exist, err = s.shippingRepo.IsShippingMethodExist(txCtx, pharmacy.Id, logisticPartnerId)
			if err != nil {
				return err
			}

			if !exist {
				err = s.shippingRepo.InsertShippingMethod(txCtx, pharmacy.Id, logisticPartnerId)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return pharmacy, nil
}

func (s *PharmacyServiceImpl) UpdatePharmacyFromPharmacistService(ctx context.Context, pharmacistCredId int64, p *entity.Pharmacy) (*entity.Pharmacy, error) {
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
	p.Id = pharmacyId

	partnerId, err := s.pharmacyRepo.GetPartnerId(ctx, pharmacyId)
	if err != nil {
		return nil, err
	}

	active, err := s.partnerRepo.IsActive(ctx, partnerId)
	if err != nil {
		return nil, err
	}
	if !active && p.IsActive {
		return nil, sentinel.ErrPartnerInactive
	}

	ids, err := s.shippingRepo.IsLogisticPartnerIdExistBulk(ctx, p.LogisticPartnerIds...)
	if len(ids) == 0 {
		return nil, sentinel.NewBadRequestError("logistic_partners[]", fmt.Sprintf("logistic_partner %v doesn't exist", ids))
	}
	if err != nil {
		return nil, err
	}

	var pharmacy *entity.Pharmacy
	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		pharmacy, err = s.pharmacyRepo.UpdatePharmacy(txCtx, p)
		if err != nil {
			return err
		}
		p.Address.PharmacyId = &pharmacy.Id

		pharmacy.Address, err = s.addressRepo.UpdateAddressPharmacy(txCtx, p.Address)
		if err != nil {
			return err
		}

		err = s.shippingRepo.DeleteShippingMethodExceptIds(txCtx, pharmacy.Id, p.LogisticPartnerIds...)
		if err != nil {
			return err
		}

		for _, logisticPartnerId := range p.LogisticPartnerIds {
			exist, err := s.shippingRepo.IsShippingMethodExist(txCtx, pharmacy.Id, logisticPartnerId)
			if err != nil {
				return err
			}

			if !exist {
				err = s.shippingRepo.InsertShippingMethod(txCtx, pharmacy.Id, logisticPartnerId)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return pharmacy, nil
}
