package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *ProductServiceImpl) UpdateProductCategoryService(ctx context.Context, p *entity.ProductCategory) (*entity.ProductCategory, error) {
	exist, err := s.productRepo.IsIdExist(ctx, p.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrProductCategoryNotFound
	}

	exist, err = s.productRepo.IsCategoryExist(ctx, p.Name)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, sentinel.ErrProductCategoryExist
	}

	productCat, err := s.productRepo.UpdateProductCategory(ctx, p)
	if err != nil {
		return nil, err
	}

	return productCat, nil
}

func (s *ProductServiceImpl) UpdateProductService(ctx context.Context, p *entity.Product) (*entity.Product, error) {
	exist, err := s.productRepo.IsProductIdExist(ctx, p.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrProductNotFound
	}

	for _, category := range p.CategoryIds {
		exist, err = s.productRepo.IsIdExist(ctx, category)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, sentinel.ErrProductCategoryNotFound
		}
	}

	product, err := s.productRepo.UpdateProduct(ctx, p)
	if err != nil {
		return nil, err
	}

	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		err = s.productRepo.DeleteProductCategoryMaps(txCtx, product.Id)
		if err != nil {
			return err
		}

		err = s.productRepo.InsertProductCategoryMaps(txCtx, product.Id, product.CategoryIds)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return p, nil
}
