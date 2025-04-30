package dto

import "healthcare-allopurinol/sentinel"

type Response struct {
	Data  any                     `json:"data,omitempty"`
	Error []sentinel.ErrorWrapper `json:"error,omitempty"`
}

type ListData struct {
	Entries  any `json:"entries"`
	PageInfo any `json:"page_info,omitempty"`
}

type ListUserProfileAddress struct {
	Profile any `json:"profile"`
	Address any `json:"address"`
}
