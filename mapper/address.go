package mapper

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
)

func AddressCreateRequestToEntity(dto *dto.AddressCreateRequest) *entity.Address {
	return &entity.Address{
		Province:    dto.Province,
		City:        dto.City,
		District:    dto.District,
		Subdistrict: dto.Subdistrict,
		PostalCode:  dto.PostalCode,
		Name:        dto.Name,
		PhoneNumber: dto.PhoneNumber,
		Latitude:    dto.Latitude,
		Longitude:   dto.Longitude,
	}
}

func AddressUpdateRequestToEntity(dto *dto.AddressUpdateRequest) *entity.Address {
	return &entity.Address{
		Province:    dto.Province,
		City:        dto.City,
		District:    dto.District,
		Subdistrict: dto.Subdistrict,
		PostalCode:  dto.PostalCode,
		Name:        dto.Name,
		PhoneNumber: dto.PhoneNumber,
		Latitude:    dto.Latitude,
		Longitude:   dto.Longitude,
	}
}

func AddressToDto(entity *entity.Address) *dto.Address {
	return &dto.Address{
		Id:          entity.Id,
		UserId:      entity.UserId,
		PharmacyId:  entity.PharmacyId,
		Province:    entity.Province,
		City:        entity.City,
		District:    entity.District,
		Subdistrict: entity.Subdistrict,
		PostalCode:  entity.PostalCode,
		Name:        entity.Name,
		PhoneNumber: entity.PhoneNumber,
		Latitude:    entity.Latitude,
		Longitude:   entity.Longitude,
	}
}
