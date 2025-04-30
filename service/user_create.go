package service

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (s *UserServiceImpl) CreateAddressService(ctx context.Context, add *entity.UserAddress) (*entity.UserAddress, error) {
	u, err := s.userRepo.GetUserProfileByCredId(ctx, add.UserId)
	if err != nil {
		return nil, err
	}
	add.UserId = u.Id
	exist, _ := s.userRepo.IsUserAddressesExist(ctx, add.UserId)
	if !exist {
		add.IsActive = true
	}

	var address *entity.UserAddress

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		address, err = s.userRepo.InsertAddress(txCtx, add)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return address, nil
}
