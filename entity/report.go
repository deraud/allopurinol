package entity

type ReportOptionsRequest struct {
	SortOrder   string
	SearchValue string
}

type Report struct {
	PharmacyId   int64
	PharmacyName string
	Sales        float64
}

type SalesByNameOptions struct {
	SearchValue string
	MonthStart  string
	MonthEnd    string
}
