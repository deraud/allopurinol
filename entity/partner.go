package entity

type Partner struct {
	Id                   int64
	Name                 string
	YearFounded          int64
	ActiveDays           string
	OperationalHourStart string
	OperationalHourEnd   string
	IsActive             bool
}

type PartnerOptions struct {
	SearchBy    string
	SearchValue string
	SortBy      string
	SortOrder   string
	Page        int
	Limit       int
	TotalRows   int
}
