package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *CartServiceImpl) IncrementCartItemService(ctx context.Context, cartItem *entity.CartItem, credId int64) error {
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

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		if !exist {
			err = s.cartRepo.InsertCartItem(txCtx, cartItem, 1)
		} else {
			err = s.cartRepo.UpdateCartItemIncrement(txCtx, cartItem)
		}
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

func (s *CartServiceImpl) DecrementCartItemService(ctx context.Context, cartItem *entity.CartItem, credId int64) error {
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
		onOne, err := s.cartRepo.IsCartItemOnOne(txCtx, cartItem)
		if err != nil {
			return err
		}

		if onOne {
			err = s.cartRepo.DeleteCartItem(txCtx, cartItem)
		} else {
			err = s.cartRepo.UpdateCartItemDecrement(txCtx, cartItem)
		}
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

func (s *CartServiceImpl) UpdateCartItemService(ctx context.Context, cartItem *entity.CartItem, value int, credId int64) error {
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

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		if !exist {
			err = s.cartRepo.InsertCartItem(txCtx, cartItem, value)
		} else {
			err = s.cartRepo.UpdateCartItem(txCtx, cartItem, value)
		}
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
