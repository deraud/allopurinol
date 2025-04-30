package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *ProductServiceImpl) GetProductCategoriesService(ctx context.Context, options *entity.ProductCategoryOptions) ([]*entity.ProductCategory, error) {
	count, err := s.productRepo.GetProductCategoriesCount(ctx, options)
	if err != nil {
		return nil, err
	}
	options.TotalRows = count

	productCats, err := s.productRepo.GetProductCategories(ctx, options)
	if err != nil {
		return nil, err
	}

	return productCats, nil
}

func (s *ProductServiceImpl) GetProductsService(ctx context.Context, options *entity.ProductOptions) ([]*entity.Product, error) {
	count, err := s.productRepo.GetProductsCount(ctx, options)
	if err != nil {
		return nil, err
	}
	options.TotalRows = count

	products, err := s.productRepo.GetProducts(ctx, options)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductServiceImpl) GetProductService(ctx context.Context, id int64) (*entity.Product, error) {
	exist, err := s.productRepo.IsProductIdExist(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrProductNotFound
	}

	product, err := s.productRepo.GetProductDetail(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductServiceImpl) GetProductClassificationsService(ctx context.Context) ([]*entity.ProductExtra, error) {
	extras, err := s.productRepo.GetProductClassifications(ctx)
	if err != nil {
		return nil, err
	}
	return extras, nil
}

func (s *ProductServiceImpl) GetProductFormsService(ctx context.Context) ([]*entity.ProductExtra, error) {
	extras, err := s.productRepo.GetProductForms(ctx)
	if err != nil {
		return nil, err
	}
	return extras, nil
}

func (s *ProductServiceImpl) GetProductManufacturersService(ctx context.Context) ([]*entity.ProductExtra, error) {
	extras, err := s.productRepo.GetProductManufacturers(ctx)
	if err != nil {
		return nil, err
	}
	return extras, nil
}
