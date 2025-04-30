package entity

type CheckoutDetails struct {
	AddressId       int64
	PaymentMethodId int64
	OrderDetails    []*CheckoutPharmacyDetails
}

type CheckoutPharmacyDetails struct {
	PharmacyId        int64
	LogisticPartnerId int64
}
