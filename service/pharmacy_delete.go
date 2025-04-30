package service

import (
	"context"
	"healthcare-allopurinol/sentinel"
)

func (s *PharmacyServiceImpl) DeletePharmacyService(ctx context.Context, id int64) error {
	exist, err := s.pharmacyRepo.IsIdExist(ctx, id)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrPharmacyNotFound
	}

	pharmacists, err := s.pharmacistRepo.GetPharmacistsByPharmacy(ctx, id)
	if err != nil {
		return err
	}
	if len(pharmacists) != 0 {
		return sentinel.ErrPharmacyHasPharmacist
	}

	hasOrders, err := s.pharmacyRepo.HasOrders(ctx, id)
	if err != nil {
		return err
	}
	if hasOrders {
		return sentinel.ErrPharmacyHasOrder
	}

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.addressRepo.DeletePharmacyAddress(txCtx, id)
		if err != nil {
			return err
		}

		err = s.shippingRepo.DeleteShippingMethod(txCtx, id)
		if err != nil {
			return err
		}

		err = s.pharmacyRepo.DeletePharmacy(txCtx, id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
