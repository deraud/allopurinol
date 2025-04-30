package entity

type Province struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	Id         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}

type District struct {
	Id     string `json:"id"`
	CityId string `json:"regency_id"`
	Name   string `json:"name"`
}

type Subdistrict struct {
	Id         string `json:"id"`
	DistrictId string `json:"district_id"`
	Name       string `json:"name"`
}

type Coordinate struct {
	Address   AddressInfo
	Latitude  float64
	Longtiude float64
}

type AddressInfo struct {
	Province    string
	City        string
	District    string
	Subdistrict string
}

type AddressDrop struct {
	Coordinate  CoordinateInfo
	Province    string
	City        string
	District    string
	Subdistrict string
}

type CoordinateInfo struct {
	Latitude  float64
	Longtiude float64
}
