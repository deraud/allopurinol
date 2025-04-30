package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *PharmacistServiceImpl) CreatePharmacistService(ctx context.Context, c *entity.Credential, p *entity.Pharmacist) (*entity.Pharmacist, *entity.Credential, error) {
	exist, err := s.authRepo.IsEmailExist(ctx, c.Email)
	if err != nil {
		return nil, nil, err
	}
	if exist {
		return nil, nil, sentinel.ErrEmailRegistered
	}

	exist, err = s.pharmacistRepo.IsSipaNumberExist(ctx, p.SipaNumber)
	if err != nil {
		return nil, nil, err
	}
	if exist {
		return nil, nil, sentinel.ErrSipaNumberRegistered
	}

	exist, err = s.pharmacistRepo.IsPhoneNumberExist(ctx, p.PhoneNumber)
	if err != nil {
		return nil, nil, err
	}
	if exist {
		return nil, nil, sentinel.ErrPhoneNumberRegistered
	}

	var (
		pharmacist *entity.Pharmacist
		credential entity.Credential
	)

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		rawPassword := c.Password

		hashed, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		c.Password = string(hashed)

		p.CredId, err = s.authRepo.InsertCredential(txCtx, c)
		if err != nil {
			return err
		}
		credential.Email = c.Email

		pharmacist, err = s.pharmacistRepo.InsertPharmacist(txCtx, p)
		if err != nil {
			return err
		}

		err = s.mailBridge.SendPharmacistEmail(ctx, c.Email, rawPassword)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return pharmacist, &credential, nil
}
