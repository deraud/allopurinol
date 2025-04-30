package entity

type RajaOngkir struct {
	Meta ROMeta   `json:"meta"`
	Data []ROData `json:"data"`
}

type RajaOngkirCost struct {
	Meta ROMeta           `json:"meta"`
	Data []ROShippingData `json:"data"`
}

type ROMeta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type ROData struct {
	Id          int64  `json:"id"`
	Label       string `json:"label"`
	SubDistrict string `json:"subdistrict_name"`
	District    string `json:"district_name"`
	City        string `json:"city_name"`
	Province    string `json:"province_name"`
	ZipCode     string `json:"zip_code"`
}

type ROShippingData struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        int64  `json:"cost"`
	Etd         string `json:"etd"`
}
