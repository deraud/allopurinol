package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *ProductServiceImpl) CreateProductCategoryService(ctx context.Context, p *entity.ProductCategory) (*entity.ProductCategory, error) {
	exist, err := s.productRepo.IsCategoryExist(ctx, p.Name)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, sentinel.ErrProductCategoryExist
	}

	var productCat *entity.ProductCategory
	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		productCat, err = s.productRepo.InsertProductCategory(txCtx, p)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return productCat, nil
}

func (s *ProductServiceImpl) CreateProductService(ctx context.Context, p *entity.Product) (*entity.Product, error) {
	exist, err := s.productRepo.IsProductRequirementsExist(ctx, p)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrProductReqNotFound
	}

	exist, err = s.productRepo.IsProductExist(ctx, p)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, sentinel.ErrProductExist
	}

	var product *entity.Product
	err = s.transactorRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		product, err = s.productRepo.InsertProduct(txCtx, p)
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

	return product, nil
}
