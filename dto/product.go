package dto

type ProductCategoryCreateRequest struct {
	Name string `json:"name" binding:"required,max=255"`
}

type ProductCategoryCreateResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductCategoryUpdateRequest struct {
	Id   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,max=255"`
}

type ProductCategoryUpdateResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductCategoryDeleteRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

type ProductCategoryOptionsRequest struct {
	SearchBy    string `form:"search_by" binding:"omitempty,oneof=name" `
	SearchValue string `form:"search_value"`
	SortBy      string `form:"sort_by" binding:"omitempty,oneof=name created_at"`
	SortOrder   string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	Page        string `form:"page" binding:"omitempty,positive"`
	Limit       string `form:"limit" binding:"omitempty,positive"`
}

type ProductCategoryOptionsResponse struct {
	Search   SearchOptions `json:"search"`
	Sort     SortOption    `json:"sort"`
	Page     int           `json:"page"`
	Limit    int           `json:"limit"`
	TotalRow int           `json:"total_row"`
}

type ProductCategoryGetResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ProductCreateRequest struct {
	ClassificationId int64   `form:"classification_id" binding:"required,gt=0"`
	FormId           *int64  `form:"form_id" binding:"required_if=ClassificationId 2,required_if=ClassificationId 3,required_if=ClassificationId 4,omitempty,gt=0"`
	ManufacturerId   int64   `form:"manufacturer_id" binding:"required,gt=0"`
	CategoryIds      []int64 `form:"category_ids" binding:"required"`
	Name             string  `form:"name" binding:"required"`
	GenericName      string  `form:"generic_name" binding:"required"`
	Description      string  `form:"description" binding:"required"`
	UnitInPack       *int    `form:"unit_in_pack" binding:"required_if=ClassificationId 2,required_if=ClassificationId 3,required_if=ClassificationId 4,omitempty,gt=0"`
	SellingUnit      *string `form:"selling_unit" binding:"omitempty"`
	Weight           string  `form:"weight" binding:"required,positive_decimal"`
	Height           string  `form:"height" binding:"required,positive_decimal"`
	Length           string  `form:"length" binding:"required,positive_decimal"`
	Width            string  `form:"width" binding:"required,positive_decimal"`
	Image            string  `form:"image" binding:"required"`
	IsActive         bool    `form:"is_active"`
}

type ProductCreateResponse struct {
	Id               int64   `json:"id,omitempty"`
	ClassificationId int64   `json:"classification_id,omitempty"`
	FormId           *int64  `json:"form_id,omitempty"`
	ManufacturerId   int64   `json:"manufacturer_id,omitempty"`
	CategoryIds      []int64 `json:"category_ids,omitempty"`
	Name             string  `json:"name,omitempty"`
	GenericName      string  `json:"generic_name,omitempty"`
	Description      string  `json:"description,omitempty"`
	UnitInPack       *int    `json:"unit_in_pack,omitempty"`
	SellingUnit      *string `json:"selling_unit,omitempty"`
	Weight           string  `json:"weight,omitempty"`
	Height           string  `json:"height,omitempty"`
	Length           string  `json:"length,omitempty"`
	Width            string  `json:"width,omitempty"`
	Image            string  `json:"image,omitempty"`
	IsActive         bool    `json:"is_active,omitempty"`
}

type ProductOptionsRequest struct {
	SearchBy    string `form:"search_by" binding:"omitempty,oneof=name generic_name description" `
	SearchValue string `form:"search_value"`
	FilterBy    string `form:"filter_by" binding:"omitempty,oneof=manufacturer_id classification_id form_id is_active" `
	FilterValue string `form:"filter_value"`
	SortBy      string `form:"sort_by" binding:"omitempty,oneof=name created_at usage"`
	SortOrder   string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	Page        string `form:"page" binding:"omitempty,positive"`
	Limit       string `form:"limit" binding:"omitempty,positive"`
}

type ProductOptionsResponse struct {
	Search   SearchOptions `json:"search"`
	Sort     SortOption    `json:"sort"`
	Filter   FilterOption  `json:"filter"`
	Page     int           `json:"page"`
	Limit    int           `json:"limit"`
	TotalRow int           `json:"total_row"`
}

type ProductGetResponse struct {
	Id             int64                        `json:"id,omitempty"`
	Classification ProductExtraGetResponse      `json:"classification,omitempty"`
	Form           *ProductFormGetResponse      `json:"form,omitempty"`
	Manufacturer   ProductExtraGetResponse      `json:"manufacturer,omitempty"`
	Categories     []ProductCategoryGetResponse `json:"categories,omitempty"`
	Name           string                       `json:"name,omitempty"`
	GenericName    string                       `json:"generic_name,omitempty"`
	Pharmacies     []ProductPharmacyGetResponse `json:"pharmacies,omitempty"`
	Stock          int                          `json:"stock"`
	Usage          int                          `json:"usage"`
	Description    string                       `json:"description,omitempty"`
	UnitInPack     *int                         `json:"unit_in_pack,omitempty"`
	SellingUnit    *string                      `json:"selling_unit,omitempty"`
	Weight         string                       `json:"weight,omitempty"`
	Height         string                       `json:"height,omitempty"`
	Length         string                       `json:"length,omitempty"`
	Width          string                       `json:"width,omitempty"`
	Image          string                       `json:"image,omitempty"`
	IsActive       bool                         `json:"is_active"`
}

type ProductPharmacyGetResponse struct {
	Id          int64                          `json:"id,omitempty"`
	Name        string                         `json:"name,omitempty"`
	Location    string                         `json:"location,omitempty"`
	Stock       int64                          `json:"stock,omitempty"`
	Price       string                         `json:"price,omitempty"`
	Pharmacists []ProductPharmacistGetResponse `json:"pharmacists,omitempty"`
}

type ProductPharmacistGetResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductCatalogGetDetailResponse struct {
	Id             int64                        `json:"id,omitempty"`
	Name           string                       `json:"name,omitempty"`
	GenericName    string                       `json:"generic_name,omitempty"`
	Manufacturer   string                       `json:"manufacturer,omitempty"`
	Classification string                       `json:"classification,omitempty"`
	Form           string                       `json:"form,omitempty"`
	Description    string                       `json:"description,omitempty"`
	UnitInPack     *int                         `json:"unit_in_pack,omitempty"`
	SellingUnit    *string                      `json:"selling_unit,omitempty"`
	Weight         string                       `json:"weight,omitempty"`
	Height         string                       `json:"height,omitempty"`
	Length         string                       `json:"length,omitempty"`
	Width          string                       `json:"width,omitempty"`
	Categories     []ProductCategoryGetResponse `json:"categories,omitempty"`
	Image          string                       `json:"image,omitempty"`
}

type ProductExtraGetResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductFormGetResponse struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ProductDeleteRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

type ProductGetRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

type ProductUpdateRequest struct {
	Id               int64   `form:"id" binding:"required,gt=0"`
	ClassificationId int64   `form:"classification_id" binding:"required,gt=0"`
	FormId           *int64  `form:"form_id" binding:"required_if=ClassificationId 1,required_if=ClassificationId 2,required_if=ClassificationId 3,omitempty,gt=0"`
	ManufacturerId   int64   `form:"manufacturer_id" binding:"required,gt=0"`
	CategoryIds      []int64 `form:"category_ids" binding:"required"`
	Name             string  `form:"name" binding:"required"`
	GenericName      string  `form:"generic_name" binding:"required"`
	Description      string  `form:"description" binding:"required"`
	UnitInPack       *int    `form:"unit_in_pack" binding:"required_if=ClassificationId 1,required_if=ClassificationId 2,required_if=ClassificationId 3,omitempty,gt=0"`
	SellingUnit      *string `form:"selling_unit" binding:"omitempty"`
	Weight           string  `form:"weight" binding:"required,positive_decimal"`
	Height           string  `form:"height" binding:"required,positive_decimal"`
	Length           string  `form:"length" binding:"required,positive_decimal"`
	Width            string  `form:"width" binding:"required,positive_decimal"`
	Image            string  `form:"image" binding:"required"`
	IsActive         bool    `form:"is_active"`
}

type ProductUpdateResponse struct {
	Id               int64   `json:"id,omitempty"`
	ClassificationId int64   `json:"classification_id,omitempty"`
	FormId           *int64  `json:"form_id,omitempty"`
	ManufacturerId   int64   `json:"manufacturer_id,omitempty"`
	CategoryIds      []int64 `json:"category_ids,omitempty"`
	Name             string  `json:"name,omitempty"`
	GenericName      string  `json:"generic_name,omitempty"`
	Description      string  `json:"description,omitempty"`
	UnitInPack       *int    `json:"unit_in_pack,omitempty"`
	SellingUnit      *string `json:"selling_unit,omitempty"`
	Weight           string  `json:"weight,omitempty"`
	Height           string  `json:"height,omitempty"`
	Length           string  `json:"length,omitempty"`
	Width            string  `json:"width,omitempty"`
	Image            string  `json:"image,omitempty"`
	IsActive         bool    `json:"is_active,omitempty"`
}

type ProductExtraResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductCheckoutCatalog struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
