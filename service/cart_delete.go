package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *CartServiceImpl) DeleteCartItemService(ctx context.Context, cartItem *entity.CartItem, credId int64) error {
	user, err := s.userRepo.GetUserByCredId(ctx, credId)
	if err != nil {
		return err
	}
	cartItem.UserId = user.Id

	exist, err := s.productRepo.IsProductIdExist(ctx, cartItem.ProductId)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrProductNotFound
	}

	exist, err = s.cartRepo.IsCartItemExist(ctx, cartItem)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrCartItemNotFound
	}

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.cartRepo.DeleteCartItem(txCtx, cartItem)
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
