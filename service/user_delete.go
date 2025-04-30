package service

import (
	"context"
	"healthcare-allopurinol/sentinel"
)

func (s *UserServiceImpl) DeleteAddressService(ctx context.Context, credId int64, addressId int64) error {
	u, err := s.userRepo.GetUserProfileByCredId(ctx, credId)
	if err != nil {
		return err
	}
	userId := u.Id
	exist, _ := s.userRepo.IsUserAddressExist(ctx, userId, addressId)
	if !exist {
		return sentinel.ErrAddressNotFound
	}
	active, _ := s.userRepo.IsUserAddressActive(ctx, userId, addressId)
	if active {
		return sentinel.ErrAddressIsActive
	}
	err = s.userRepo.DeleteAddress(ctx, userId, addressId)
	if err != nil {
		return err
	}
	return nil
}
