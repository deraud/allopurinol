package mapper

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"strconv"
)

func AddressDtoToEntity(addressRequest *dto.AddressDrop) *entity.AddressInfo {
	return &entity.AddressInfo{
		Province:    addressRequest.Province,
		City:        addressRequest.City,
		District:    addressRequest.District,
		Subdistrict: addressRequest.Subdistrict,
	}
}

func CoordinateEntityToDtoResponse(coordinate *entity.Coordinate) *dto.CoordinateResponse {
	return &dto.CoordinateResponse{
		Latitude:  coordinate.Latitude,
		Longitude: coordinate.Longtiude,
	}
}

func CoordinateRequestDtoToEntity(coordinateRequest *dto.CoordinateRequest) *entity.Coordinate {
	return &entity.Coordinate{
		Address: *AddressDtoToEntity(&coordinateRequest.Address),
	}
}

func CoordinateDtoToEntity(coordinateRequest *dto.Coordinate) *entity.CoordinateInfo {
	return &entity.CoordinateInfo{
		Latitude:  coordinateRequest.Latitude,
		Longtiude: coordinateRequest.Longitude,
	}
}

func AddressEntityToDtoResponse(address *entity.AddressDrop) *dto.AddressResponse {
	return &dto.AddressResponse{
		Province:    address.Province,
		City:        address.City,
		District:    address.District,
		Subdistrict: address.Subdistrict,
	}
}

func AddressRequestDtotoEntity(addressRequest *dto.AddressRequest) *entity.AddressDrop {
	return &entity.AddressDrop{
		Coordinate: *CoordinateDtoToEntity(&addressRequest.Coordinate),
	}
}

func CityRequestToProvinceId(dto *dto.LocationCityRequest) int64 {
	provinceId, _ := strconv.ParseInt(dto.ProvinceId, 10, 64)
	return provinceId
}

func DistrictRequestToCityId(dto *dto.LocationDistrictRequest) int64 {
	cityId, _ := strconv.ParseInt(dto.CityId, 10, 64)
	return cityId
}

func SubdistrictRequestToDistrictId(dto *dto.LocationSubdistrictRequest) int64 {
	districtId, _ := strconv.ParseInt(dto.DistrictId, 10, 64)
	return districtId
}

func ProvinceToDto(entity *entity.Province) *dto.Province {
	id, _ := strconv.ParseInt(entity.Id, 10, 64)
	return &dto.Province{
		Id:   id,
		Name: entity.Name,
	}
}

func CityToDto(entity *entity.City) *dto.City {
	id, _ := strconv.ParseInt(entity.Id, 10, 64)
	provinceId, _ := strconv.ParseInt(entity.ProvinceId, 10, 64)
	return &dto.City{
		Id:         id,
		ProvinceId: provinceId,
		Name:       entity.Name,
	}
}

func DistrictToDto(entity *entity.District) *dto.District {
	id, _ := strconv.ParseInt(entity.Id, 10, 64)
	cityId, _ := strconv.ParseInt(entity.CityId, 10, 64)
	return &dto.District{
		Id:     id,
		CityId: cityId,
		Name:   entity.Name,
	}
}

func SubdistrictToDto(entity *entity.Subdistrict) *dto.Subdistrict {
	id, _ := strconv.ParseInt(entity.Id, 10, 64)
	districtId, _ := strconv.ParseInt(entity.DistrictId, 10, 64)
	return &dto.Subdistrict{
		Id:         id,
		DistrictId: districtId,
		Name:       entity.Name,
	}
}
