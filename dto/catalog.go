package dto

import "github.com/shopspring/decimal"

type CatalogCreateRequest struct {
	PharmacyId int64  `json:"pharmacy_id" binding:"required,min=1"`
	ProductId  int64  `json:"product_id" binding:"required,min=1"`
	Stock      int    `json:"stock" binding:"min=0"`
	Price      string `json:"price" binding:"required,positive_decimal"`
}

type Catalog struct {
	Id         int64           `json:"id"`
	PharmacyId int64           `json:"pharmacy_id"`
	ProductId  int64           `json:"product_id"`
	Stock      int             `json:"stock"`
	Price      decimal.Decimal `json:"price"`
	IsActive   bool            `json:"is_active"`
}

type CatalogUpdateRequest struct {
	Id       int64  `json:"id" binding:"required,min=1"`
	Stock    int    `json:"stock" binding:"min=0"`
	Price    string `json:"price" binding:"required,positive_decimal"`
	IsActive *bool  `json:"is_active" binding:"required"`
}

type CatalogDeleteRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type CatalogOptionsRequest struct {
	SearchBy         string `form:"search_by" binding:"omitempty,oneof=name generic_name" `
	SearchValue      string `form:"search_value"`
	SortBy           string `form:"sort_by" binding:"omitempty,oneof=name created_at stock"`
	SortOrder        string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	ManufacturerId   string `form:"manufacturer_id" binding:"omitempty,positive"`
	ClassificationId string `form:"classification_id" binding:"omitempty,positive"`
	FormId           string `form:"form_id"  binding:"omitempty,positive"`
	IsActive         string `form:"is_active" binding:"omitempty,boolean"`
	Page             string `form:"page" binding:"omitempty,positive"`
	Limit            string `form:"limit" binding:"omitempty,positive"`
}

type CatalogOptionsResponse struct {
	Search           SearchOptions `json:"search"`
	Sort             SortOption    `json:"sort"`
	ManufacturerId   *int64        `json:"manufacturer_id"`
	ClassificationId *int64        `json:"classification_id"`
	FormId           *int64        `json:"form_id"`
	IsActive         *bool         `json:"is_active"`
	Page             int           `json:"page"`
	Limit            int           `json:"limit"`
	TotalRow         int           `json:"total_row"`
}

type CatalogGetResponse struct {
	Id             int64   `json:"id"`
	Name           string  `json:"name"`
	GenericName    string  `json:"generic_name"`
	Manufacturer   string  `json:"manufacturer"`
	Classification string  `json:"classification"`
	Form           *string `json:"form"`
	Stock          int     `json:"stock"`
	IsActive       bool    `json:"is_active"`
}

type CatalogGetDetailRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type CatalogGetDetailResponse struct {
	Id             int64           `json:"id"`
	Stock          int             `json:"stock"`
	Price          decimal.Decimal `json:"price"`
	IsActive       bool            `json:"is_active"`
	Name           string          `json:"name"`
	GenericName    string          `json:"generic_name"`
	Manufacturer   string          `json:"manufacturer"`
	Classification string          `json:"classification"`
	Form           *string         `json:"form"`
	Description    string          `json:"description"`
	UnitInPack     *int             `json:"unit_in_pack"`
	SellingUnit    string          `json:"selling_unit"`
	Image          string          `json:"image"`
}

type AvailableCatalogOptionsRequest struct {
	SearchBy    string `form:"search_by" binding:"omitempty,oneof=name generic_name" `
	SearchValue string `form:"search_value"`
	CategoryId  string `form:"category_id" binding:"omitempty,positive"`
	Page        string `form:"page" binding:"omitempty,positive"`
	Limit       string `form:"limit" binding:"omitempty,positive"`
}

type AvailableCatalogOptionsResponse struct {
	Search   SearchOptions `json:"search"`
	Page     int           `json:"page"`
	Limit    int           `json:"limit"`
	TotalRow int           `json:"total_row"`
}

type AvailableCatalogGetResponse struct {
	Id      int64                           `json:"id"`
	Price   decimal.Decimal                 `json:"price"`
	Stock   int                             `json:"stock"`
	Product ProductCatalogGetDetailResponse `json:"product"`
}

type AvailableCatalogGetDetailRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type AvailableCatalogGetDetailResponse struct {
	Id       int64                            `json:"id"`
	Price    decimal.Decimal                  `json:"price"`
	Stock    int                              `json:"stock"`
	Product  ProductCatalogGetDetailResponse  `json:"product"`
	Pharmacy PharmacyCatalogGetDetailResponse `json:"pharmacy"`
}

type CheckoutCatalogGetRequest struct {
	AddressId int64 `json:"address_id" binding:"required,min=1"`
}

type CheckoutCatalogDetail struct {
	Id       int64                  `json:"id,omitempty"`
	Price    *decimal.Decimal       `json:"price,omitempty"`
	Stock    int                    `json:"stock,omitempty"`
	Quantity int                    `json:"quantity,omitempty"`
	Product  ProductCheckoutCatalog `json:"product,omitempty"`
}

type CheckoutCatalogGetResponse struct {
	Pharmacy PharmacyCheckoutDetail   `json:"pharmacy"`
	Catalogs []*CheckoutCatalogDetail `json:"catalogs"` // + quantity
}
