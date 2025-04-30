package service

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (s *AdminServiceImpl) GetDashboardCountService(ctx context.Context, dashboard *entity.DashboardCount) (*entity.DashboardCount, error) {
	var err error
	dashboard.Pharmacy, err = s.adminRepo.GetPharmacyCount(ctx)
	if err != nil {
		return nil, err
	}
	dashboard.Pharmacist, err = s.adminRepo.GetPharmacistCount(ctx)
	if err != nil {
		return nil, err
	}
	dashboard.User, err = s.adminRepo.GetUserCount(ctx)
	if err != nil {
		return nil, err
	}
	return dashboard, nil
}
