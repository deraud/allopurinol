package dto

type AddressCreateRequest struct {
	Province    string  `json:"province" binding:"required"`
	City        string  `json:"city" binding:"required"`
	District    string  `json:"district" binding:"required"`
	Subdistrict string  `json:"subdistrict" binding:"required"`
	PostalCode  string  `json:"postal_code" binding:"required,postal_code"`
	Name        string  `json:"name" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,e164"`
	Latitude    float64 `json:"latitude" binding:"required,latitude"`
	Longitude   float64 `json:"longitude" binding:"required,longitude"`
}

type AddressUpdateRequest struct {
	Province    string  `json:"province" binding:"required"`
	City        string  `json:"city" binding:"required"`
	District    string  `json:"district" binding:"required"`
	Subdistrict string  `json:"subdistrict" binding:"required"`
	PostalCode  string  `json:"postal_code" binding:"required,postal_code"`
	Name        string  `json:"name" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,e164"`
	Latitude    float64 `json:"latitude" binding:"required,latitude"`
	Longitude   float64 `json:"longitude" binding:"required,longitude"`
}

type Address struct {
	Id          int64   `json:"id,omitempty"`
	UserId      *int64  `json:"user_id,omitempty"`
	PharmacyId  *int64  `json:"pharmacy_id,omitempty"`
	Province    string  `json:"province,omitempty"`
	City        string  `json:"city,omitempty"`
	District    string  `json:"district,omitempty"`
	Subdistrict string  `json:"subdistrict,omitempty"`
	PostalCode  string  `json:"postal_code,omitempty"`
	Name        string  `json:"name,omitempty"`
	PhoneNumber string  `json:"phone_number,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
}

type AddressOrderDetail struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
