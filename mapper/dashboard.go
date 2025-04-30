package mapper

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
)

func DashboardCountDtoToEntity(dashboard *dto.DashboardCount) *entity.DashboardCount {
	return &entity.DashboardCount{
		Pharmacy:   dashboard.Pharmacy,
		Pharmacist: dashboard.Pharmacist,
		User:       dashboard.User,
	}
}

func DashboardCountEntityToDto(dashboard *entity.DashboardCount) *dto.DashboardCount {
	return &dto.DashboardCount{
		Pharmacy:   dashboard.Pharmacy,
		Pharmacist: dashboard.Pharmacist,
		User:       dashboard.User,
	}
}
