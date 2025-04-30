package dto

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserGetDetailRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type UserProfileResponse struct {
	Id             int64                       `json:"id"`
	Name           string                      `json:"name"`
	ProfilePicture string                      `json:"profile_picture"`
	Email          string                      `json:"email"`
	Gender         bool                        `json:"gender"`
	IsVerified     bool                        `json:"is_verified"`
	Address        []*UserGetAddressesResponse `json:"address"`
}

type UserUpdateProfileRequest struct {
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	Gender         *bool  `json:"gender"`
}

type UserUpdateProfileResponse struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	Gender         bool   `json:"gender"`
	IsVerified     bool   `json:"is_verified"`
}

type UserGetAddressesResponse struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Province    string  `json:"province"`
	City        string  `json:"city"`
	District    string  `json:"district"`
	Subdistrict string  `json:"subdistrict"`
	PhoneNumber string  `json:"phone_number"`
	PostalCode  string  `json:"postal_code"`
	IsActive    bool    `json:"is_active"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type UserGetAddressesOptionsRequest struct {
	IsActive string `form:"is_active"`
}

type UserAddressDeleteRequest struct {
	AddressId int64 `uri:"address_id" binding:"required,min=1"`
}

type UserAddressCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Province    string  `json:"province" binding:"required"`
	City        string  `json:"city" binding:"required"`
	District    string  `json:"district" binding:"required"`
	Subdistrict string  `json:"subdistrict" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,e164"`
	PostalCode  string  `json:"postal_code" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required,latitude"`
	Longitude   float64 `json:"longitude" binding:"required,longitude"`
}

type UserAddressUpdateRequest struct {
	AddressId   int64   `json:"id" binding:"required,min=1"`
	Name        string  `json:"name" binding:"required"`
	Province    string  `json:"province" binding:"required"`
	City        string  `json:"city" binding:"required"`
	District    string  `json:"district" binding:"required"`
	Subdistrict string  `json:"subdistrict" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,e164"`
	PostalCode  string  `json:"postal_code" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required,latitude"`
	Longitude   float64 `json:"longitude" binding:"required,longitude"`
}

type UserActivateAddressUpdateRequest struct {
	AddressId int64 `json:"id" binding:"required,min=1"`
}

type UserActivateAddressUpdateResponse struct {
	AddressId int64 `json:"id"`
	IsActive  bool  `json:"is_active"`
}
