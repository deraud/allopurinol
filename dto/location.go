package dto

type AddressDrop struct {
	Province    string `json:"province" binding:"required"`
	City        string `json:"city" binding:"required"`
	District    string `json:"district" binding:"required"`
	Subdistrict string `json:"subdistrict" binding:"required"`
}

type CoordinateRequest struct {
	Address AddressDrop `form:"address" binding:"required"`
}

type CoordinateResponse struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

type Coordinate struct {
	Latitude  float64 `json:"Latitude" binding:"required"`
	Longitude float64 `json:"Longitude" binding:"required"`
}

type AddressRequest struct {
	Coordinate Coordinate `form:"coordinate" binding:"required"`
}

type AddressResponse struct {
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	Subdistrict string `json:"subdistrict"`
}

type Province struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type City struct {
	Id         int64  `json:"id"`
	ProvinceId int64  `json:"province_id"`
	Name       string `json:"name"`
}

type District struct {
	Id     int64  `json:"id"`
	CityId int64  `json:"city_id"`
	Name   string `json:"name"`
}

type Subdistrict struct {
	Id         int64  `json:"id"`
	DistrictId int64  `json:"district_id"`
	Name       string `json:"name"`
}

type LocationCityRequest struct {
	ProvinceId string `uri:"province_id" binding:"required,positive"`
}

type LocationDistrictRequest struct {
	CityId string `uri:"city_id" binding:"required,positive"`
}

type LocationSubdistrictRequest struct {
	DistrictId string `uri:"district_id" binding:"required,positive"`
}
