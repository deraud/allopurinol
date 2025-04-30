package dto

type Partner struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type PartnerCreateRequest struct {
	Name                 string `json:"name" binding:"required,max=255"`
	YearFounded          int64  `json:"year_founded" binding:"required,min=0"`
	ActiveDays           string `json:"active_days" binding:"required"`
	OperationalHourStart string `json:"operational_hour_start" binding:"required,datetime=15:04:05"`
	OperationalHourEnd   string `json:"operational_hour_end" binding:"required,datetime=15:04:05,gtcsfield=OperationalHourStart"`
	IsActive             *bool  `json:"is_active" binding:"required"`
}

type PartnerCreateResponse struct {
	Id                   int64  `json:"id"`
	Name                 string `json:"name"`
	YearFounded          int64  `json:"year_founded"`
	ActiveDays           string `json:"active_days"`
	OperationalHourStart string `json:"operational_hour_start"`
	OperationalHourEnd   string `json:"operational_hour_end"`
	IsActive             bool   `json:"is_active"`
}

type PartnerUpdateRequest struct {
	Id                   int64  `json:"id" binding:"required"`
	ActiveDays           string `json:"active_days" binding:"required_if=IsActive true"`
	OperationalHourStart string `json:"operational_hour_start" binding:"required_if=IsActive true,omitempty,datetime=15:04:05"`
	OperationalHourEnd   string `json:"operational_hour_end" binding:"required_if=IsActive true,omitempty,datetime=15:04:05,gtcsfield=OperationalHourStart"`
	IsActive             *bool  `json:"is_active" binding:"required"`
}

type PartnerUpdateResponse struct {
	Id                   int64  `json:"id"`
	Name                 string `jsong:"name"`
	YearFounded          int64  `json:"year_founded"`
	ActiveDays           string `json:"active_days"`
	OperationalHourStart string `json:"operational_hour_start"`
	OperationalHourEnd   string `json:"operational_hour_end"`
	IsActive             bool   `json:"is_active"`
}

type PartnerDeleteRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type PartnerOptionsRequest struct {
	SearchBy    string `form:"search_by" binding:"omitempty,oneof=name" `
	SearchValue string `form:"search_value"`
	SortBy      string `form:"sort_by" binding:"omitempty,oneof=name created_at"`
	SortOrder   string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	Page        string `form:"page" binding:"omitempty,positive"`
	Limit       string `form:"limit" binding:"omitempty,positive"`
}

type PartnerOptionsResponse struct {
	Search   SearchOptions `json:"search"`
	Sort     SortOption    `json:"sort"`
	Page     int           `json:"page"`
	Limit    int           `json:"limit"`
	TotalRow int           `json:"total_row"`
}

type PartnerGetResponse struct {
	Id                   int64  `json:"id"`
	Name                 string `json:"name"`
	YearFounded          int64  `json:"year_founded"`
	ActiveDays           string `json:"active_days"`
	OperationalHourStart string `json:"operational_hour_start"`
	OperationalHourEnd   string `json:"operational_hour_end"`
	IsActive             bool   `json:"is_active"`
}

type PartnerGetDetailRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}
