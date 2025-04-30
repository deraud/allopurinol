package mapper

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
)

func LogisticPartnerToGetResponseDto(entity *entity.LogisticPartner) *dto.LogisticPartnerGetResponse {
	return &dto.LogisticPartnerGetResponse{
		Id:   entity.Id,
		Name: entity.Name,
	}
}
