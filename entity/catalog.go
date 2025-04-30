package entity

import (
	"github.com/shopspring/decimal"
)

type Catalog struct {
	Id         int64
	PharmacyId int64
	ProductId  int64
	Stock      int
	Price      decimal.Decimal
	IsActive   bool

	Name           string
	GenericName    string
	Manufacturer   string
	Classification string
	Form           *string
	Description    string
	UnitInPack     *int
	SellingUnit    string
	Image          string

	Product  *Product
	Pharmacy *Pharmacy
	
	Quantity int
}

type CatalogOptions struct {
	SearchBy         string
	SearchValue      string
	SortBy           string
	SortOrder        string
	ManufacturerId   *int64
	ClassificationId *int64
	FormId           *int64
	IsActive         *bool
	Page             int
	Limit            int
	TotalRows        int
}

type AvailableCatalogOptions struct {
	SearchBy    string
	SearchValue string
	CategoryId  *int64
	Page        int
	Limit       int
	TotalRows   int
}
