package service

import (
	"context"
	"github.com/google/uuid"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/sentinel"
)

func (s *AuthServiceImpl) VerifyService(ctx context.Context, credId int64) error {
	exist, err := s.redis.IsVerifyExist(credId)
	if err != nil {
		return err
	}
	if exist {
		return sentinel.ErrAccountOnCooldown
	}

	exist, err = s.repo.IsIdExist(ctx, credId)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrCredentialIdNotFound
	}

	cred, err := s.repo.GetCredentialById(ctx, credId)
	if err != nil {
		return err
	}

	if cred.RoleId != constant.ROLE_USER {
		return sentinel.ErrNonUserCannotVerify
	}

	user, err := s.userRepo.GetUserByCredId(ctx, credId)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return sentinel.ErrUserAlreadyVerified
	}

	token, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	err = s.mailBridge.SendVerifyEmail(ctx, cred.Email, user.Name, token.String())
	if err != nil {
		return err
	}

	err = s.redis.SetVerify(credId, token.String())
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) VerifyTokenService(ctx context.Context, token string) error {
	credId, err := s.redis.GetVerifyValue(token)
	if err != nil {
		return sentinel.ErrInvalidVerifyToken
	}

	exist, err := s.repo.IsIdExist(ctx, credId)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrCredentialIdNotFound
	}

	err = s.transactor.WithTransaction(ctx, func(ctx context.Context) error {
		err = s.repo.UpdateVerified(ctx, credId)
		if err != nil {
			return err
		}

		err = s.redis.DeleteVerify(token, credId)
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
