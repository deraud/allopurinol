package service

import (
	"context"
	"healthcare-allopurinol/bridge"
	"healthcare-allopurinol/entity"
	"strings"
)

type LocationServiceItf interface {
	PostCoordinate(ctx context.Context, address *entity.Coordinate) (*entity.Coordinate, error)
	PostAddress(ctx context.Context, coordinate *entity.AddressDrop) (*entity.AddressDrop, error)
	GetProvinces(ctx context.Context) ([]*entity.Province, error)
	GetCities(ctx context.Context, provinceId int64) ([]*entity.City, error)
	GetDistricts(ctx context.Context, cityId int64) ([]*entity.District, error)
	GetSubdistricts(ctx context.Context, districtId int64) ([]*entity.Subdistrict, error)
}

type LocationServiceImpl struct {
	locationBridge bridge.LocationBridgeItf
}

func NewLocationService(locationBridge bridge.LocationBridgeItf) LocationServiceItf {
	return &LocationServiceImpl{
		locationBridge: locationBridge,
	}
}

func (s *LocationServiceImpl) PostCoordinate(ctx context.Context, address *entity.Coordinate) (*entity.Coordinate, error) {

	address.Address.Province = strings.ReplaceAll(address.Address.Province, " ", "%20")
	address.Address.City = strings.ReplaceAll(address.Address.City, " ", "%20")
	address.Address.District = strings.ReplaceAll(address.Address.District, " ", "%20")
	address.Address.Subdistrict = strings.ReplaceAll(address.Address.Subdistrict, " ", "%20")

	coordinate, err := s.locationBridge.PostCoordinate(*address)
	if err != nil {
		return nil, err
	}
	return coordinate, err
}

func (s *LocationServiceImpl) PostAddress(ctx context.Context, coordinate *entity.AddressDrop) (*entity.AddressDrop, error) {

	var address entity.AddressDrop

	location, err := s.locationBridge.PostAddress(*coordinate)
	if err != nil {
		return nil, err
	}

	stringSlice := strings.Split(location, ", ")
	province := stringSlice[len(stringSlice)-4]
	address.Province = province
	city := stringSlice[len(stringSlice)-5]
	address.City = city
	if len(stringSlice) > 5 {
		district := stringSlice[len(stringSlice)-6]
		address.District = district
	}
	if len(stringSlice) > 6 {
		subdistrict := stringSlice[len(stringSlice)-7]
		address.Subdistrict = subdistrict
	}
	if len(stringSlice) == 6 {
		subdistrict := stringSlice[len(stringSlice)-6]
		address.Subdistrict = subdistrict
	}

	return &address, err
}

func (s *LocationServiceImpl) GetProvinces(ctx context.Context) ([]*entity.Province, error) {
	provinces, err := s.locationBridge.GetProvinces()
	if err != nil {
		return nil, err
	}
	return provinces, nil
}

func (s *LocationServiceImpl) GetCities(ctx context.Context, provinceId int64) ([]*entity.City, error) {
	cities, err := s.locationBridge.GetCities(provinceId)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func (s *LocationServiceImpl) GetDistricts(ctx context.Context, cityId int64) ([]*entity.District, error) {
	districts, err := s.locationBridge.GetDistricts(cityId)
	if err != nil {
		return nil, err
	}

	return districts, nil
}

func (s *LocationServiceImpl) GetSubdistricts(ctx context.Context, districtId int64) ([]*entity.Subdistrict, error) {
	subdistricts, err := s.locationBridge.GetSubdistricts(districtId)
	if err != nil {
		return nil, err
	}

	return subdistricts, nil
}
