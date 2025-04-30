package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *UserServiceImpl) UpdateAddressService(ctx context.Context, add *entity.UserAddress) (*entity.UserAddress, error) {
	u, err := s.userRepo.GetUserProfileByCredId(ctx, add.UserId)
	if err != nil {
		return nil, err
	}
	add.UserId = u.Id
	exist, _ := s.userRepo.IsUserAddressExist(ctx, add.UserId, *add.Id)
	if !exist {
		return nil, sentinel.ErrAddressNotFound
	}
	address, err := s.userRepo.UpdateAddress(ctx, add)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *UserServiceImpl) UpdateUserProfileService(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error) {
	id, err := s.userRepo.GetUserProfileByCredId(ctx, u.Id)
	if err != nil {
		return nil, err
	}
	u.Id = id.Id

	user, err := s.userRepo.UpdateUserProfile(ctx, u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateUserActivateAddressService(ctx context.Context, add *entity.UserAddress) (*entity.UserAddress, error) {
	u, err := s.userRepo.GetUserProfileByCredId(ctx, add.UserId)
	if err != nil {
		return nil, err
	}
	add.UserId = u.Id
	exist, _ := s.userRepo.IsUserAddressExist(ctx, add.UserId, *add.Id)
	if !exist {
		return nil, sentinel.ErrAddressNotFound
	}
	err = s.userRepo.DeactivateUserAddresses(ctx, add)
	if err != nil {
		return nil, err
	}
	address, err := s.userRepo.UpdateUserActivateAddress(ctx, add)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (s *UserServiceImpl) UpdateUserRemovePictureService(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error) {
	id, err := s.userRepo.GetUserProfileByCredId(ctx, u.Id)
	if err != nil {
		return nil, err
	}
	u.Id = id.Id
	user, err := s.userRepo.UpdateUserRemovePicture(ctx, u)
	if err != nil {
		return nil, err
	}
	return user, nil
}
