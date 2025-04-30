package dto

import (
	"github.com/shopspring/decimal"
)

type OrderGroupCreateRequest struct {
	AddressId            int64                  `json:"address_id" binding:"required,min=1"`
	PaymentMethodId      int64                  `json:"payment_method_id" binding:"required,min=1"`
	OrderPharmacyDetails []*OrderPharmacyDetail `json:"order_pharmacy_details" binding:"required,dive"`
}

type OrderPharmacyDetail struct {
	PharmacyId        int64 `json:"pharmacy_id" binding:"required,min=1"`
	LogisticPartnerId int64 `json:"logistic_partner_id" binding:"required,min=1"`
}

type OrderCreateResponse struct {
	Id                 int64                      `json:"id"`
	AddressId          int64                      `json:"address_id"`
	PaymentMethodId    int64                      `json:"payment_method_id"`
	PharmacyId         int64                      `json:"pharmacy_id"`
	LogisticPartnerId  int64                      `json:"logistic_partner_id"`
	TotalPriceProduct  decimal.Decimal            `json:"total_price_product"`
	TotalPriceShipping decimal.Decimal            `json:"total_price_shipping"`
	OrderItems         []*OrderItemCreateResponse `json:"order_items"`
}

type OrderItemCreateResponse struct {
	Id        int64 `json:"id"`
	CatalogId int64 `json:"catalog_id"`
	Quantity  int   `json:"quantity"`
}

type OrderGroupCreateResponse struct {
	Id     int64                  `json:"id"`
	Orders []*OrderCreateResponse `json:"orders"`
}

type OrderGroupUpdateRequest struct {
	Id    int64  `json:"id" binding:"required,min=1"`
	Proof string `json:"proof" binding:"required"`
}

type OrderGroupUpdateResponse struct {
	Id    int64  `json:"id"`
	Proof string `json:"proof"`
}

type PendingOrderOptionsRequest struct {
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=created_at"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
	Page      string `form:"page" binding:"omitempty,positive"`
	Limit     string `form:"limit" binding:"omitempty,positive"`
}

type PendingOrderOptionsResponse struct {
	Sort     SortOption `json:"sort"`
	Page     int        `json:"page"`
	Limit    int        `json:"limit"`
	TotalRow int        `json:"total_row"`
}

type PendingOrderGroupGetResponse struct {
	Id           int64                   `json:"id,omitempty"`
	Orders       []*PendingOrderResponse `json:"orders,omitempty"`
	ShippingCost string                  `json:"total_shipping_cost,omitempty"`
	TotalPrice   string                  `json:"total_price,omitempty"`
	UserAddress  string                  `json:"address,omitempty"`
	CreatedAt    string                  `json:"created_at,omitempty"`
}

type PendingOrderResponse struct {
	Id           int64                     `json:"id,omitempty"`
	PharmacyId   int64                     `json:"pharmacy_id,omitempty"`
	PharmacyName string                    `json:"pharmacy_name,omitempty"`
	ShippingCost string                    `json:"shipping_cost,omitempty"`
	Catalogs     []*PendingCatalogResponse `json:"catalogs,omitempty"`
}

type PendingCatalogResponse struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
	Price    string `json:"price,omitempty"`
}

type UserOrderUpdateRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type UserOrderUpdateResponse struct {
	Id       int64 `json:"id"`
	StatusId int64 `json:"status_id"`
}

type PharmacyOrderOptionsRequest struct {
	Page  string `form:"page" binding:"omitempty,positive"`
	Limit string `form:"limit" binding:"omitempty,positive"`
}

type PharmacyOrderOptionsResponse struct {
	Page     int `json:"page"`
	Limit    int `json:"limit"`
	TotalRow int `json:"total_row"`
}

type PharmacyOrderGetResponse struct {
	Id                int64  `json:"id"`
	Status            string `json:"status"`
	ProductCount      int    `json:"product_count"`
	TotalPriceProduct string `json:"total_price_product"`
}

type PharmacyOrderGetDetailRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type PharmacyOrderGetDetailResponse struct {
	Id                 int64                      `json:"id"`
	User               User                       `json:"user"`
	Address            AddressOrderDetail         `json:"address"`
	OrderStatus        OrderStatus                `json:"order_status"`
	LogisticPartner    LogisticPartnerGetResponse `json:"logistic_partner"`
	PaymentMethod      PaymentMethod              `json:"payment_method"`
	TotalPriceShipping string                     `json:"total_price_shipping"`
	TotalPriceProduct  string                     `json:"total_price_product"`
	OrderItems         []OrderItem                `json:"order_items"`
}

type OrderStatus struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type PharmacyOrderUpdateRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type PharmacyOrderUpdateResponse struct {
	Id       int64 `json:"id"`
	StatusId int64 `json:"status_id"`
}

type OrderOptionsRequest struct {
	PharmacyId string `form:"pharmacy_id" binding:"omitempty,positive"`
	StatusId   string `form:"status_id" binding:"omitempty,oneof=1 2 3 4 5"`
	Page       string `form:"page" binding:"omitempty,positive"`
	Limit      string `form:"limit" binding:"omitempty,positive"`
}

type OrderOptionsResponse struct {
	PharmacyId *int64 `json:"pharmacy_id"`
	StatusId   *int64 `json:"status_id"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalRow   int    `json:"total_row"`
}

type OrderGetResponse struct {
	Id       int64               `json:"id"`
	Pharmacy PharmacyOrderDetail `json:"pharmacy"`
	Status   OrderStatus         `json:"status"`
}

type OrderGetDetailRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type OrderGetDetailResponse struct {
	Id                 int64                      `json:"id"`
	Pharmacy           PharmacyOrderDetail        `json:"pharmacy"`
	User               User                       `json:"user"`
	Address            AddressOrderDetail         `json:"address"`
	OrderStatus        OrderStatus                `json:"order_status"`
	LogisticPartner    LogisticPartnerGetResponse `json:"logistic_partner"`
	PaymentMethod      PaymentMethod              `json:"payment_method"`
	TotalPriceShipping string                     `json:"total_price_shipping"`
	TotalPriceProduct  string                     `json:"total_price_product"`
	OrderItems         []OrderItem                `json:"order_items"`
}

type OrderItem struct {
	Id           int64  `json:"id"`
	Quantity     int    `json:"quantity"`
	CatalogId    int64  `json:"catalog_id"`
	Price        string `json:"price"`
	ProductId    int64  `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
}

type UserOrderOptionsRequest struct {
	FilterValue string `form:"filter" binding:"required,oneof=processed sent order_confirmed canceled verifying"`
	Page        string `form:"page" binding:"omitempty,positive"`
	Limit       string `form:"limit" binding:"omitempty,positive"`
}

type UserOrderOptionsResponse struct {
	Filter   FilterOption `json:"filter"`
	Page     int          `json:"page"`
	Limit    int          `json:"limit"`
	TotalRow int          `json:"total_row"`
}

type UserOrderGetResponse struct {
	Id                 int64        `json:"id"`
	Status             string       `json:"status"`
	PharmacyName       string       `json:"pharmacy_name"`
	Address            string       `json:"address"`
	CreatedAt          string       `json:"created_at"`
	TotalPriceShipping string       `json:"total_price_shipping"`
	TotalPriceProduct  string       `json:"total_price_product"`
	OrderItems         []*OrderItem `json:"order_items"`
}
