package service

import (
	"context"
	"healthcare-allopurinol/sentinel"
)

func (s *ProductServiceImpl) DeleteProductCategoryService(ctx context.Context, id int64) error {
	exist, err := s.productRepo.IsIdExist(ctx, id)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrProductCategoryNotFound
	}

	exist, err = s.productRepo.IsCategoryInProductExist(ctx, id)
	if err != nil {
		return err
	}
	if exist {
		return sentinel.ErrProductCategoryIsUsed
	}

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.productRepo.DeleteProductCategory(txCtx, id)
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

func (s *ProductServiceImpl) DeleteProductService(ctx context.Context, id int64) error {
	exist, err := s.productRepo.IsProductIdExist(ctx, id)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrProductNotFound
	}

	exist, err = s.productRepo.IsProductBoughtExist(ctx, id)
	if err != nil {
		return err
	}
	if exist {
		return sentinel.ErrProductBought
	}

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.productRepo.DeleteProductCategoryMaps(txCtx, id)
		if err != nil {
			return err
		}

		catalogIds, err := s.productRepo.GetProductCatalogIds(txCtx, id)
		if err != nil {
			return err
		}

		if catalogIds != nil {
			err = s.productRepo.DeleteProductCarts(txCtx, catalogIds)
			if err != nil {
				return err
			}

			err = s.productRepo.DeleteProductCatalog(txCtx, id)
			if err != nil {
				return err
			}
		}

		err = s.productRepo.DeleteProduct(txCtx, id)
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
