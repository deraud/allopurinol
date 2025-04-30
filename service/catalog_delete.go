package service

import (
	"context"
	"healthcare-allopurinol/sentinel"
)

func (s *CatalogServiceImpl) DeleteCatalogService(ctx context.Context, pharmacistCredId int64, id int64) error {
	exist, err := s.catalogRepo.IsIdExist(ctx, id)
	if err != nil {
		return err
	}
	if !exist {
		return sentinel.ErrCatalogNotFound
	}

	pharmacyId, err := s.catalogRepo.GetPharmacyId(ctx, id)
	if err != nil {
		return err
	}

	assigned, err := s.pharmacistRepo.IsAssignedToSpecificPharmacy(ctx, pharmacistCredId, pharmacyId)
	if err != nil {
		return err
	}
	if !assigned {
		return sentinel.ErrPharmacistNoAccess
	}

	ordered, err := s.catalogRepo.IsOrdered(ctx, id)
	if err != nil {
		return err
	}
	if ordered {
		return sentinel.ErrCatalogDelete
	}

	err = s.catalogRepo.DeleteCatalog(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
