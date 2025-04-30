package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthServiceImpl) RegisterService(ctx context.Context, cred *entity.Credential, user *entity.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cred.Password = string(hashed)

	err = s.transactor.WithTransaction(ctx, func(ctx context.Context) error {
		exist, err := s.repo.IsEmailExist(ctx, cred.Email)
		if err != nil {
			return err
		}
		if exist {
			return sentinel.ErrEmailRegistered
		}

		credId, err := s.repo.InsertCredential(ctx, cred)
		if err != nil {
			return err
		}

		user.CredId = credId
		err = s.userRepo.InsertUser(ctx, user)
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
