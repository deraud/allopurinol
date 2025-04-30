package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type UserServiceItf interface {
	GetUserAddresses(ctx context.Context, id int64, options *entity.UserAddressOptions) ([]*entity.UserAddress, error)
	GetUserProfile(ctx context.Context, id int64) (*entity.UserProfile, error)
	UpdateUserProfileService(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error)
	UpdateUserRemovePictureService(ctx context.Context, u *entity.UserProfile) (*entity.UserProfile, error)
	DeleteAddressService(ctx context.Context, userId int64, addressId int64) error
	CreateAddressService(ctx context.Context, add *entity.UserAddress) (*entity.UserAddress, error)
	UpdateAddressService(ctx context.Context, add *entity.UserAddress) (*entity.UserAddress, error)
	UpdateUserActivateAddressService(ctx context.Context, add *entity.UserAddress) (*entity.UserAddress, error)
}

type UserServiceImpl struct {
	transactorRepo repo.TransactorItf
	userRepo       repo.UserRepoItf
}

func NewUserService(transactorRepo repo.TransactorItf, userRepo repo.UserRepoItf) UserServiceItf {
	return &UserServiceImpl{
		transactorRepo: transactorRepo,
		userRepo:       userRepo,
	}
}
