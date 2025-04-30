package entity

import (
	"github.com/shopspring/decimal"
)

type Shipping struct {
	PharmacyId     int64
	PharmacyLat    float64
	PharmacyLong   float64
	PharmacyPostal string
	UserId         int64
	UserLat        float64
	UserLong       float64
	UserPostal     string
	Catalogs       []*Catalog
	TotalWeight    decimal.Decimal
	ShippingCosts  []*ShippingCost
}

type ShippingCost struct {
	MethodId int64
	Method   string
	Cost     decimal.Decimal
}
