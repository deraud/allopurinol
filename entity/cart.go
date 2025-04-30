package entity

type CartItem struct {
	UserId    int64
	ProductId int64
	Quantity  int

	Product Product
}
