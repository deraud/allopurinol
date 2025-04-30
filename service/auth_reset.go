package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthServiceImpl) ForgotPasswordService(ctx context.Context, email string) error {
	exist, err := s.redis.IsResetExist(email)
	if err != nil {
		return err
	}
	if exist {
		return sentinel.ErrEmailOnCooldown
	}

	exist, err = s.repo.IsEmailExist(ctx, email)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}

	token, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	err = s.mailBridge.SendResetEmail(ctx, email, token.String())
	if err != nil {
		return err
	}

	err = s.redis.SetReset(email, token.String())
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) ResetPasswordService(ctx context.Context, token, password string) error {
	email, err := s.redis.GetResetValue(token)
	if err != nil {
		return sentinel.ErrInvalidResetToken
	}

	exist, err := s.repo.IsEmailExist(ctx, email)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrAccountNotFound
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.transactor.WithTransaction(ctx, func(ctx context.Context) error {
		err = s.repo.UpdateCredential(ctx, &entity.Credential{
			Email:    email,
			Password: string(hashed),
		})
		if err != nil {
			return err
		}

		err = s.redis.DeleteReset(token, email)
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
