package service

import (
	"context"
	"healthcare-allopurinol/entity"
)

func (s *AdminServiceImpl) GetSalesReportService(ctx context.Context, options *entity.ReportOptionsRequest) ([]*entity.Report, error) {

	count, err := s.adminRepo.GetReportCountBySearchValue(ctx, options)
	if err != nil {
		return nil, err
	}
	reports, err := s.adminRepo.GetSalesReport(ctx, options, count)
	if err != nil {
		return nil, err
	}
	return reports, err
}
