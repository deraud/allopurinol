package dto

type ReportOptionsRequest struct {
	SortOrder   string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	SearchValue string `form:"search_value"`
}

type GetReportResponse struct {
	Id    int64   `json:"pharmacy_id"`
	Name  string  `json:"pharmacy_name"`
	Sales float64 `json:"total_sales:"`
}
