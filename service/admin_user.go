package service

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (s *AdminServiceImpl) GetUsersService(ctx context.Context, options *entity.UserInfoOptions) ([]*entity.UserInfo, *entity.UserInfoOptions, error) {
	count, err := s.adminRepo.GetUsersCount(ctx, options)
	if err != nil {
		return nil, nil, err
	}
	options.TotalRows = count

	users, err := s.adminRepo.GetUsers(ctx, options)
	if err != nil {
		return nil, nil, err
	}

	return users, options, nil
}
