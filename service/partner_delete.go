package service

import (
	"context"
	"healthcare-allopurinol/sentinel"
)

func (s *PartnerServiceImpl) DeletePartnerService(ctx context.Context, id int64) error {
	exist, _ := s.partnerRepo.IsPartnerIdExist(ctx, id)
	if !exist {
		return sentinel.ErrPartnerNotFound
	}
	err := s.partnerRepo.DeletePartner(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
