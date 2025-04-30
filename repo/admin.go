package repo

import (
	"context"
	"database/sql"
	"healthcare-allopurinol/entity"
)

type AdminRepoItf interface {
	GetUsers(ctx context.Context, options *entity.UserInfoOptions) ([]*entity.UserInfo, error)
	GetUsersCount(ctx context.Context, options *entity.UserInfoOptions) (int, error)
	GetPharmacyCount(ctx context.Context) (int64, error)
	GetPharmacistCount(ctx context.Context) (int64, error)
	GetUserCount(ctx context.Context) (int64, error)
	GetSalesReport(ctx context.Context, options *entity.ReportOptionsRequest, count int64) ([]*entity.Report, error)
	GetReportCountBySearchValue(ctx context.Context, options *entity.ReportOptionsRequest) (int64, error)
}

type AdminRepoImpl struct {
	db *sql.DB
}

func NewAdminRepo(database *sql.DB) AdminRepoItf {
	return &AdminRepoImpl{
		db: database,
	}
}
