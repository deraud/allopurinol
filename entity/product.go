package entity

import (
	"github.com/shopspring/decimal"
)

type ProductCategory struct {
	Id   int64
	Name string
}

type ProductCategoryOptions struct {
	SearchBy    string
	SearchValue string
	SortBy      string
	SortOrder   string
	Page        int
	Limit       int
	TotalRows   int
}

type Product struct {
	Id               int64
	ClassificationId int64
	Classification   string
	FormId           *int64
	Form             *string
	ManufacturerId   int64
	Manufacturer     string
	CategoryIds      []int64
	Categories       []string
	Name             string
	GenericName      string
	Description      string
	Stock            *int
	Usage            *int
	UnitInPack       *int
	SellingUnit      *string
	Weight           decimal.Decimal
	Height           decimal.Decimal
	Length           decimal.Decimal
	Width            decimal.Decimal
	Image            string
	ImageLink        string
	IsActive         bool
}

type ProductOptions struct {
	SearchBy    string
	SearchValue string
	FilterBy    string
	FilterValue string
	SortBy      string
	SortOrder   string
	Page        int
	Limit       int
	TotalRows   int
}

type ProductExtra struct {
	Id   int64
	Name string
}
