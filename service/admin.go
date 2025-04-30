package service

import (
	"context"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/repo"
)

type AdminServiceItf interface {
	GetUsersService(ctx context.Context, options *entity.UserInfoOptions) ([]*entity.UserInfo, *entity.UserInfoOptions, error)
	GetDashboardCountService(ctx context.Context, dashboard *entity.DashboardCount) (*entity.DashboardCount, error)
	GetSalesReportService(ctx context.Context, options *entity.ReportOptionsRequest) ([]*entity.Report, error)
}

type AdminServiceImpl struct {
	adminRepo repo.AdminRepoItf
}

func NewAdminService(adminRepo repo.AdminRepoItf) AdminServiceItf {
	return &AdminServiceImpl{
		adminRepo: adminRepo,
	}
}
