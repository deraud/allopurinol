package service

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (s *LogisticPartnerServiceImpl) GetLogisticPartnerService(ctx context.Context) ([]*entity.LogisticPartner, error) {
	logisticPartners, err := s.shippingRepo.GetLogisticPartners(ctx)
	if err != nil {
		return nil, err
	}

	return logisticPartners, nil
}
