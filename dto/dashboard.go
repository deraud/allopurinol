package dto

type DashboardCount struct {
	Pharmacy   int64 `json:"pharmacy"`
	Pharmacist int64 `json:"pharmacist"`
	User       int64 `json:"user"`
}
