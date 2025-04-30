package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *CatalogServiceImpl) CreateCatalogService(ctx context.Context, pharmacistCredId int64, c *entity.Catalog) (*entity.Catalog, error) {
	exist, err := s.pharmacyRepo.IsIdExist(ctx, c.PharmacyId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrPharmacyNotFound
	}

	assigned, err := s.pharmacistRepo.IsAssignedToSpecificPharmacy(ctx, pharmacistCredId, c.PharmacyId)
	if err != nil {
		return nil, err
	}
	if !assigned {
		return nil, sentinel.ErrPharmacistNoAccess
	}

	exist, err = s.productRepo.IsProductIdExist(ctx, c.ProductId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrProductNotFound
	}

	exist, err = s.catalogRepo.IsCatalogExist(ctx, c.PharmacyId, c.ProductId)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, sentinel.ErrCatalogCreated
	}

	catalog, err := s.catalogRepo.InsertCatalog(ctx, c)
	if err != nil {
		return nil, err
	}

	return catalog, err
}
