package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PharmacistServiceImpl) DeletePharmacistService(ctx context.Context, id int64) error {
	exist, err := s.pharmacistRepo.IsIdExist(ctx, id)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrPharmacistNotFound
	}

	isAssigned, err := s.pharmacistRepo.IsAssigned(ctx, id)
	if err != nil {
		return err
	}
	if isAssigned {
		return sentinel.ErrPharmacistAssigned
	}

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		credId, err := s.pharmacistRepo.DeletePharmacist(txCtx, id)
		if err != nil {
			return err
		}

		err = s.authRepo.DeleteCredential(txCtx, &entity.Credential{Id: credId})
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
