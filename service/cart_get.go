package service

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (s *CartServiceImpl) GetCartItemsService(ctx context.Context, credId int64) ([]*entity.CartItem, error) {
	user, err := s.userRepo.GetUserByCredId(ctx, credId)
	if err != nil {
		return nil, err
	}

	items, err := s.cartRepo.GetCartItems(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	return items, nil
}
