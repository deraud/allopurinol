package mapper

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
)

func UserAddressesEntityToDto(entity *entity.UserAddress) *dto.UserGetAddressesResponse {
	return &dto.UserGetAddressesResponse{
		Id:          *entity.Id,
		Name:        entity.Name,
		Province:    entity.Province,
		City:        entity.City,
		District:    entity.District,
		Subdistrict: entity.Subdistrict,
		PhoneNumber: entity.PhoneNumber,
		PostalCode:  entity.PostalCode,
		IsActive:    entity.IsActive,
		Latitude:    entity.Latitude,
		Longitude:   entity.Longitude,
	}
}

func UserAddressesOptionsToEntity(dto *dto.UserGetAddressesOptionsRequest) *entity.UserAddressOptions {
	return &entity.UserAddressOptions{
		IsActive: &dto.IsActive,
	}
}

func UserProfileEntityToDto(entity *entity.UserProfile, address []*dto.UserGetAddressesResponse) *dto.UserProfileResponse {
	return &dto.UserProfileResponse{
		Id:             entity.Id,
		Name:           entity.Name,
		ProfilePicture: entity.ProfilePicture,
		Email:          entity.Email,
		Gender:         *entity.Gender,
		IsVerified:     entity.IsVerified,
		Address:        address,
	}
}

func UserUpdateProfileEntityToDto(entity *entity.UserProfile) *dto.UserUpdateProfileResponse {
	return &dto.UserUpdateProfileResponse{
		Id:             entity.Id,
		Name:           entity.Name,
		ProfilePicture: entity.ProfilePicture,
		Gender:         *entity.Gender,
		IsVerified:     entity.IsVerified,
	}
}

func UserUpdateProfileRequestToEntity(userId int64, user *dto.UserUpdateProfileRequest) *entity.UserProfile {

	return &entity.UserProfile{
		Id:             userId,
		Name:           user.Name,
		ProfilePicture: user.ProfilePicture,
		Gender:         user.Gender,
	}
}

func UserAddressDtoToEntity(userId int64, address *dto.UserAddressCreateRequest) *entity.UserAddress {
	return &entity.UserAddress{
		UserId:      userId,
		Name:        address.Name,
		Province:    address.Province,
		City:        address.City,
		District:    address.District,
		Subdistrict: address.Subdistrict,
		PhoneNumber: address.PhoneNumber,
		PostalCode:  address.PostalCode,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	}
}

func UserUpdateAddressDtoToEntity(userId int64, address *dto.UserAddressUpdateRequest) *entity.UserAddress {
	return &entity.UserAddress{
		Id:          &address.AddressId,
		UserId:      userId,
		Name:        address.Name,
		Province:    address.Province,
		City:        address.City,
		District:    address.District,
		Subdistrict: address.Subdistrict,
		PhoneNumber: address.PhoneNumber,
		PostalCode:  address.PostalCode,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	}
}

func UserActivateAddressDtoToEntity(userId int64, address *dto.UserActivateAddressUpdateRequest) *entity.UserAddress {
	return &entity.UserAddress{
		Id:     &address.AddressId,
		UserId: userId,
	}
}

func UserAddressEntityToActivateAddressResponse(address *entity.UserAddress) *dto.UserActivateAddressUpdateResponse {
	return &dto.UserActivateAddressUpdateResponse{
		AddressId: *address.Id,
		IsActive:  address.IsActive,
	}
}
