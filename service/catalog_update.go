package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
)

func (s *CatalogServiceImpl) UpdateCatalogService(ctx context.Context, pharmacistCredId int64, c *entity.Catalog) (*entity.Catalog, error) {
	exist, err := s.catalogRepo.IsIdExist(ctx, c.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, sentinel.ErrCatalogNotFound
	}

	pharmacyId, err := s.catalogRepo.GetPharmacyId(ctx, c.Id)
	if err != nil {
		return nil, err
	}

	assigned, err := s.pharmacistRepo.IsAssignedToSpecificPharmacy(ctx, pharmacistCredId, pharmacyId)
	if err != nil {
		return nil, err
	}
	if !assigned {
		return nil, sentinel.ErrPharmacistNoAccess
	}

	catalog, err := s.catalogRepo.UpdateCatalog(ctx, c)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}
