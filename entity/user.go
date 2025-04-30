package entity

type User struct {
	Id             int64
	CredId         int64
	Name           string
	Gender         *bool
	ProfilePicture *string
	IsVerified     bool
}

type UserProfile struct {
	Id             int64
	CredId         int64
	Name           string
	ProfilePicture string
	Email          string
	Gender         *bool
	IsVerified     bool
}

type UserAddress struct {
	Id          *int64
	UserId      int64
	Name        string
	Province    string
	City        string
	District    string
	PostalCode  string
	Subdistrict string
	PhoneNumber string
	IsActive    bool
	Latitude    float64
	Longitude   float64
}

type UserAddressOptions struct {
	IsActive *string
}
