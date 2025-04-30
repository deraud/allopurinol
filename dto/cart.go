package dto

type CartItemRequest struct {
	ProductId int64 `json:"product_id" binding:"required,gt=0"`
}

type CartItemSetRequest struct {
	ProductId int64 `json:"product_id" binding:"required,gt=0"`
	Quantity  int   `json:"quantity" binding:"required,gt=0"`
}

type CartItemResponse struct {
	Id       int64 `json:"id,omitempty"`
	Quantity int   `json:"quantity,omitempty"`
}
