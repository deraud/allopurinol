package entity

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type OrderGroup struct {
	Id     int64
	UserId int64
	Proof  string

	Orders []*Order
}

type Order struct {
	Id                 int64
	UserId             int64
	AddressId          int64
	AddressName        string
	StatusId           int64
	PaymentMethodId    int64
	PharmacyId         int64
	PharmacyName       string
	LogisticPartnerId  int64
	OrderGroupId       int64
	CreatedAt          pgtype.Timestamp
	TotalPriceProduct  decimal.Decimal
	TotalPriceShipping decimal.Decimal

	OrderItems      []*OrderItem
	User            *User
	Address         *Address
	OrderStatus     *OrderStatus
	PaymentMethod   *PaymentMethod
	LogisticPartner *LogisticPartner
	Pharmacy        *Pharmacy
}

type OrderItem struct {
	Id        int64
	OrderId   int64
	CatalogId int64
	Quantity  int
	Price     decimal.Decimal

	Catalog *Catalog
}

type PendingOrderGroup struct {
	Id           int64
	Order        []*PendingOrder
	ShippingCost decimal.Decimal
	TotalPrice   decimal.Decimal
	UserAddress  string
	CreatedAt    pgtype.Timestamp
}

type PendingOrder struct {
	Id           int64
	PharmacyId   int64
	PharmacyName string
	ShippingCost decimal.Decimal
	Catalogs     []*PendingCatalog
}

type PendingCatalog struct {
	Id       int64
	Name     string
	Quantity int64
	Price    decimal.Decimal
}

type PendingOrderOptions struct {
	SortBy    string
	SortOrder string
	Page      int
	Limit     int
	TotalRows int
}

type PharmacyOrderOptions struct {
	Page      int
	Limit     int
	TotalRows int
}

type OrderStatus struct {
	Id   int64
	Name string
}

type OrderOptions struct {
	PharmacyId *int64
	StatusId   *int64
	Page       int
	Limit      int
	TotalRows  int
}

type UserOrderOptions struct {
	FilterBy    string
	FilterValue string
	Page        int
	Limit       int
	TotalRows   int
}
