package entity

type Pharmacist struct {
	Id                int64
	CredId            int64
	PharmacyId        *int64
	Name              string
	SipaNumber        string
	PhoneNumber       string
	YearsOfExperience int

	PharmacyName *string
}

type PharmacistOptions struct {
	SearchBy      string
	SearchValue   string
	SortBy        string
	SortOrder     string
	Assigned      *bool
	YearsExpStart *int
	YearsExpEnd   *int
	Page          int
	Limit         int
	TotalRows     int
}
