package dto

type ShippingRequest struct {
	PharmacyId int64 `json:"pharmacy_id" binding:"required,gt=0"`
	AddressId  int64 `json:"address_id" binding:"required,gt=0"`
}

type ShippingResponse struct {
	ShippingCosts []*ShippingCostResponse `json:"costs,omitempty"`
}

type ShippingCostResponse struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Cost string `json:"cost,omitempty"`
}
