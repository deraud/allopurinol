package dto

type UserInfoRequest struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type UserInfoOptionsRequest struct {
	SearchBy    string `form:"search_by" binding:"omitempty,oneof=email"`
	SearchValue string `form:"search_value"`
	Page        string `form:"page" binding:"omitempty,positive"`
	Limit       string `form:"limit" binding:"omitempty,positive"`
}

type UserInfoOptionsResponse struct {
	Search   SearchOptions `json:"search"`
	Page     int           `json:"page"`
	Limit    int           `json:"limit"`
	TotalRow int           `json:"total_row"`
}
