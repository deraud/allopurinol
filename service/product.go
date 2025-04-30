package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type ProductServiceItf interface {
	CreateProductCategoryService(ctx context.Context, p *entity.ProductCategory) (*entity.ProductCategory, error)
	UpdateProductCategoryService(ctx context.Context, p *entity.ProductCategory) (*entity.ProductCategory, error)
	DeleteProductCategoryService(ctx context.Context, id int64) error
	GetProductCategoriesService(ctx context.Context, options *entity.ProductCategoryOptions) ([]*entity.ProductCategory, error)
	CreateProductService(ctx context.Context, p *entity.Product) (*entity.Product, error)
	GetProductsService(ctx context.Context, options *entity.ProductOptions) ([]*entity.Product, error)
	DeleteProductService(ctx context.Context, id int64) error
	GetProductService(ctx context.Context, id int64) (*entity.Product, error)
	UpdateProductService(ctx context.Context, p *entity.Product) (*entity.Product, error)
	GetProductClassificationsService(ctx context.Context) ([]*entity.ProductExtra, error)
	GetProductFormsService(ctx context.Context) ([]*entity.ProductExtra, error)
	GetProductManufacturersService(ctx context.Context) ([]*entity.ProductExtra, error)
}

type ProductServiceImpl struct {
	transactorRepo repo.TransactorItf
	productRepo    repo.ProductRepoItf
}

func NewProductService(transactorRepo repo.TransactorItf, productRepo repo.ProductRepoItf) ProductServiceItf {
	return &ProductServiceImpl{
		transactorRepo: transactorRepo,
		productRepo:    productRepo,
	}
}
