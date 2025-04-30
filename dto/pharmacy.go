package dto

type PharmacyOptionsRequest struct {
	SearchBy    string `form:"search_by" binding:"omitempty,oneof=name"`
	SearchValue string `form:"search_value"`
	Page        string `form:"page" binding:"omitempty,positive"`
	Limit       string `form:"limit" binding:"omitempty,positive"`
}

type PharmacyOptionsResponse struct {
	Search   SearchOptions `json:"search"`
	Page     int           `json:"page"`
	Limit    int           `json:"limit"`
	TotalRow int           `json:"total_row"`
}

type PharmacyGetResponse struct {
	Id       int64    `json:"id"`
	Partner  *Partner `json:"partner"`
	Name     string   `json:"name"`
	Logo     string   `json:"logo"`
	Address  string   `json:"address,omitempty"`
	IsActive bool     `json:"is_active"`
}

type PharmacyCreateRequest struct {
	PartnerId        int64                `form:"partner_id" binding:"required,min=1"`
	Address          AddressCreateRequest `form:"address" binding:"required"`
	Name             string               `form:"name" binding:"required"`
	Logo             string               `form:"logo" binding:"required"`
	Pharmacists      []int64              `form:"pharmacists[]" binding:"omitempty,dive,min=1"`
	LogisticPartners []int64              `form:"logistic_partners[]" binding:"required,dive,min=1"`
	IsActive         *bool                 `form:"is_active" binding:"required"`
}

type PharmacyCreateResponse struct {
	Id                 int64   `json:"id"`
	PartnerId          int64   `json:"partner_id"`
	AddressId          int64   `json:"address_id"`
	Name               string  `json:"name"`
	Logo               string  `json:"logo"`
	PharmacistIds      []int64 `json:"pharmacist_ids"`
	LogisticPartnerIds []int64 `json:"logistic_partner_ids"`
	IsActive           bool    `json:"is_active"`
}

type PharmacyUpdateRequest struct {
	Id               int64                `form:"id" binding:"required,min=1"`
	Address          AddressUpdateRequest `form:"address" binding:"required"`
	Name             string               `form:"name" binding:"required"`
	Logo             string               `form:"logo" binding:"required"`
	Pharmacists      []int64              `form:"pharmacists[]" binding:"omitempty,dive,min=1"`
	LogisticPartners []int64              `form:"logistic_partners[]" binding:"required,dive,min=1"`
	IsActive         *bool                `form:"is_active" binding:"required"`
}

type PharmacyUpdateResponse struct {
	Id                 int64   `json:"id"`
	PartnerId          int64   `json:"partner_id"`
	AddressId          int64   `json:"address_id"`
	Name               string  `json:"name"`
	Logo               string  `json:"logo"`
	PharmacistIds      []int64 `json:"pharmacist_ids"`
	LogisticPartnerIds []int64 `json:"logistic_partner_ids"`
	IsActive           bool    `json:"is_active"`
}

type PharmacyGetDetailRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type PharmacyGetDetailResponse struct {
	Id               int64                         `json:"id"`
	Partner          *Partner                      `json:"partner"`
	Address          *Address                      `json:"address"`
	Name             string                        `json:"name"`
	Logo             string                        `json:"logo"`
	Pharmacists      []*PharmacistGetResponse      `json:"pharmacists"`
	LogisticPartners []*LogisticPartnerGetResponse `json:"logistic_partners"`
	IsActive         bool                          `json:"is_active"`
}

type PharmacyDeleteRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type PharmacistPharmacyUpdateRequest struct {
	Address          AddressUpdateRequest `form:"address" binding:"required"`
	Name             string               `form:"name" binding:"required"`
	Logo             string               `form:"logo" binding:"required"`
	LogisticPartners []int64              `form:"logistic_partners[]" binding:"required,dive,min=1"`
	IsActive         bool                 `form:"is_active"`
}

type PharmacyOrderDetail struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type PharmacyCatalogGetDetailResponse struct {
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	Address *Address `json:"address"`
}

type PharmacyCheckoutDetail struct {
	Id               int64                         `json:"id"`
	Name             string                        `json:"name"`
	Address          *Address                      `json:"address"`
	LogisticPartners []*LogisticPartnerGetResponse `json:"logistic_partners"`
}
