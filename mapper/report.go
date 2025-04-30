package mapper

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
)

func ReportOptionsDtoToEntity(dto *dto.ReportOptionsRequest) *entity.ReportOptionsRequest {
	return &entity.ReportOptionsRequest{
		SortOrder:   dto.SortOrder,
		SearchValue: dto.SearchValue,
	}
}

func ReportEntityToDto(entity *entity.Report) *dto.GetReportResponse {
	return &dto.GetReportResponse{
		Id:    entity.PharmacyId,
		Name:  entity.PharmacyName,
		Sales: entity.Sales,
	}
}
