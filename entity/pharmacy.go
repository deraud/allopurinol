package entity

import "mime/multipart"

type Pharmacy struct {
	Id        int64
	PartnerId int64
	Name      string
	Logo      string
	IsActive  bool

	Partner *Partner
	Address *Address
	Pharmacists []*Pharmacist
	LogisticPartners []*LogisticPartner

	LogoFile           *multipart.File
	PharmacistIds      []int64
	LogisticPartnerIds []int64
}

type PharmacyOptions struct {
	SearchBy    string
	SearchValue string
	Page        int
	Limit       int
	TotalRows   int
}
