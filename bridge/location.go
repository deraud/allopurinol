package bridge

import (
	"encoding/json"
	"fmt"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
	"io"
	"net/http"

	nominatim "github.com/doppiogancio/go-nominatim"
)

type LocationBridgeItf interface {
	PostCoordinate(address entity.Coordinate) (*entity.Coordinate, error)
	PostAddress(coordinate entity.AddressDrop) (string, error)
	GetProvinces() ([]*entity.Province, error)
	GetCities(provinceId int64) ([]*entity.City, error)
	GetDistricts(cityId int64) ([]*entity.District, error)
	GetSubdistricts(districtId int64) ([]*entity.Subdistrict, error)
}

type LocationBridgeImpl struct {
	baseUrl string
}

func NewLocationBridge(baseUrl string) LocationBridgeItf {
	return &LocationBridgeImpl{
		baseUrl: baseUrl,
	}
}

func (b *LocationBridgeImpl) PostCoordinate(address entity.Coordinate) (*entity.Coordinate, error) {
	var coordinate entity.Coordinate
	add := fmt.Sprintf("%s,%s", address.Address.Subdistrict, address.Address.Province)
	coord, err := nominatim.Geocode(add)
	if err != nil {
		return nil, err
	}
	coordinate.Latitude = coord.Latitude
	coordinate.Longtiude = coord.Longitude
	return &coordinate, nil
}

func (b *LocationBridgeImpl) PostAddress(coordinate entity.AddressDrop) (string, error) {

	latitude := coordinate.Coordinate.Latitude
	longitude := coordinate.Coordinate.Longtiude

	location, err := nominatim.ReverseGeocode(
		latitude,  // Latitude
		longitude, // Longitude
		"id",      // Language (en,id)
	)

	address := location.DisplayName // full address in one string

	if err != nil {
		fmt.Println(err)
	}

	return address, nil
}

func (b *LocationBridgeImpl) GetProvinces() ([]*entity.Province, error) {
	url := fmt.Sprintf("%s/provinces.json", b.baseUrl)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var provinces []*entity.Province
	err = json.Unmarshal(body, &provinces)
	if err != nil {
		return nil, err
	}
	return provinces, nil
}

func (b *LocationBridgeImpl) GetCities(provinceId int64) ([]*entity.City, error) {
	url := fmt.Sprintf("%s/regencies/%d.json", b.baseUrl, provinceId)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil, sentinel.ErrProvinceNotFound
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cities []*entity.City
	err = json.Unmarshal(body, &cities)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (b *LocationBridgeImpl) GetDistricts(cityId int64) ([]*entity.District, error) {
	url := fmt.Sprintf("%s/districts/%d.json", b.baseUrl, cityId)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil, sentinel.ErrCityNotFound
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var districts []*entity.District
	err = json.Unmarshal(body, &districts)
	if err != nil {
		return nil, err
	}
	return districts, nil
}

func (b *LocationBridgeImpl) GetSubdistricts(districtId int64) ([]*entity.Subdistrict, error) {
	url := fmt.Sprintf("%s/villages/%d.json", b.baseUrl, districtId)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil, sentinel.ErrDistrictNotFound
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var subdistricts []*entity.Subdistrict
	err = json.Unmarshal(body, &subdistricts)
	if err != nil {
		return nil, err
	}
	return subdistricts, nil
}
