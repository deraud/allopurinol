package entity

type Address struct {
	Id          int64
	UserId      *int64
	PharmacyId  *int64
	Province    string
	City        string
	District    string
	Subdistrict string
	PostalCode  string
	Name        string
	PhoneNumber string
	IsActive    bool
	Latitude    float64
	Longitude   float64
}
