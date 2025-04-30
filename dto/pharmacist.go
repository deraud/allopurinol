package dto

type PharmacistCreateRequest struct {
	Name              string `json:"name" binding:"required"`
	SipaNumber        string `json:"sipa_number" binding:"required"`
	PhoneNumber       string `json:"phone_number" binding:"required,e164"`
	YearsOfExperience int    `json:"years_of_experience" binding:"required,min=0"`
	Email             string `json:"email" binding:"required,email"`
	Password          string `json:"password" binding:"required,password"`
}

type PharmacistCreateResponse struct {
	Id                int64  `json:"id"`
	PharmacyId        *int64 `json:"pharmacy_id"`
	Name              string `json:"name"`
	SipaNumber        string `json:"sipa_number"`
	PhoneNumber       string `json:"phone_number"`
	YearsOfExperience int    `json:"years_of_experience"`
	Email             string `json:"email"`
}

type PharmacistUpdateRequest struct {
	Id                int64  `json:"id" binding:"required"`
	PharmacyId        *int64 `json:"pharmacy_id" binding:"omitempty,min=1"`
	PhoneNumber       string `json:"phone_number" binding:"required,e164"`
	YearsOfExperience int    `json:"years_of_experience" binding:"required,min=0"`
}

type PharmacistUpdateResponse struct {
	Id                int64  `json:"id"`
	PharmacyId        *int64 `json:"pharmacy_id"`
	Name              string `json:"name"`
	SipaNumber        string `json:"sipa_number"`
	PhoneNumber       string `json:"phone_number"`
	YearsOfExperience int    `json:"years_of_experience"`
}

type PharmacistDeleteRequest struct {
	Id int64 `uri:"id" binding:"required"`
}

type PharmacistOptionsRequest struct {
	SearchBy      string `form:"search_by" binding:"omitempty,oneof=name email sipa_number phone_number" `
	SearchValue   string `form:"search_value"`
	SortBy        string `form:"sort_by" binding:"omitempty,oneof=name created_at"`
	SortOrder     string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	Assigned      string `form:"assigned" binding:"omitempty,boolean"`
	YearsExpStart string `form:"years_exp_start" binding:"omitempty,non_negative"`
	YearsExpEnd   string `form:"years_exp_end"  binding:"omitempty,non_negative"`
	Page          string `form:"page" binding:"omitempty,positive"`
	Limit         string `form:"limit" binding:"omitempty,positive"`
}

type PharmacistOptionsResponse struct {
	Search        SearchOptions `json:"search"`
	Sort          SortOption    `json:"sort"`
	Assigned      *bool         `json:"assigned"`
	YearsExpStart *int          `json:"years_exp_start"`
	YearsExpEnd   *int          `json:"years_exp_end"`
	Page          int           `json:"page"`
	Limit         int           `json:"limit"`
	TotalRow      int           `json:"total_row"`
}

type PharmacistGetResponse struct {
	Id                int64   `json:"id"`
	PharmacyId        *int64  `json:"pharmacy_id"`
	PharmacyName      *string `json:"pharmacy_name"`
	Name              string  `json:"name"`
	SipaNumber        string  `json:"sipa_number"`
	PhoneNumber       string  `json:"phone_number"`
	YearsOfExperience int     `json:"years_of_experience"`
	Email             string  `json:"email,omitempty"`
}

type PharmacistGetDetailRequest struct {
	Id int64 `uri:"id" binding:"required"`
}
