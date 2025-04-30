package service

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (s *UserServiceImpl) GetUserProfile(ctx context.Context, id int64) (*entity.UserProfile, error) {
	u, err := s.userRepo.GetUserProfileByCredId(ctx, id)
	if err != nil {
		return nil, err
	}
	id = u.Id

	userProfile, err := s.userRepo.GetUserProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}

func (s *UserServiceImpl) GetUserAddresses(ctx context.Context, id int64, options *entity.UserAddressOptions) ([]*entity.UserAddress, error) {
	u, err := s.userRepo.GetUserProfileByCredId(ctx, id)
	if err != nil {
		return nil, err
	}
	id = u.Id
	count, err := s.userRepo.GetUserAddressesCount(ctx, id, options)
	if err != nil {
		return nil, err
	}
	userAdresses, err := s.userRepo.GetUserAddresses(ctx, id, count, options)
	if err != nil {
		return nil, err
	}

	return userAdresses, nil
}
