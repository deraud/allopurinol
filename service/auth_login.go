package service

import (
	"context"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthServiceImpl) LoginService(ctx context.Context, email, password string) (string, error) {
	exist, err := s.repo.IsEmailExist(ctx, email)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", sentinel.ErrLogin
	}

	cred, err := s.repo.GetCredential(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(cred.Password), []byte(password))
	if err != nil {
		return "", sentinel.ErrLogin
	}

	isVerified := false
	if cred.RoleId == constant.ROLE_USER {
		isVerified, err = s.userRepo.IsUserVerified(ctx, cred.Id)
		if err != nil {
			return "", err
		}
	}

	token, err := utility.GenerateJWToken(utility.ClaimsContent{
		Id:   cred.Id,
		Role: cred.RoleId,
		IsVerified: isVerified,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
