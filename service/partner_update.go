package service

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (s *PartnerServiceImpl) UpdatePartnerService(ctx context.Context, p *entity.Partner) (*entity.Partner, error) {

	partner, err := s.partnerRepo.UpdatePartner(ctx, p)
	if err != nil {
		return nil, err
	}
	if !partner.IsActive {
		_ = s.partnerRepo.DeactivatePharmacies(ctx, p)
	}
	return partner, nil
}
