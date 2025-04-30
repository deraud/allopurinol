package sentinel

type ErrorWrapper struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ErrorWrapper) Error() string {
	return e.Message
}

type BadRequestError struct {
	ErrorWrapper
}

func (e BadRequestError) Error() string {
	return e.ErrorWrapper.Message
}

func NewBadRequestError(field string, message string) BadRequestError {
	err := BadRequestError{
		ErrorWrapper{
			Field:   field,
			Message: message,
		},
	}
	return err
}

type ForbiddenError struct {
	ErrorWrapper
}

func (e ForbiddenError) Error() string {
	return e.ErrorWrapper.Message
}

func NewForbiddenError(field string, message string) ForbiddenError {
	err := ForbiddenError{
		ErrorWrapper{
			Field:   field,
			Message: message,
		},
	}
	return err
}

type UnauthorizedError struct {
	ErrorWrapper
}

func (e UnauthorizedError) Error() string {
	return e.ErrorWrapper.Message
}

func NewUnauthorizedError(field string, message string) UnauthorizedError {
	err := UnauthorizedError{
		ErrorWrapper{
			Field:   field,
			Message: message,
		},
	}
	return err
}

var (
	ErrUnauthenticated = NewUnauthorizedError("token", "login required")
	ErrTokenExpired = NewUnauthorizedError("token", "token expired")
	ErrTokenInvalid = NewUnauthorizedError("token", "token invalid")
	
	ErrUnauthorized = NewForbiddenError("role", "permission denied")
	ErrUnverified   = NewForbiddenError("is_verified", "user not verified")

	ErrLogin                   = NewBadRequestError("email,password", "email or password is incorrect")
	ErrEmailOnCooldown         = NewBadRequestError("email", "email is on cooldown")
	ErrAccountNotFound         = NewBadRequestError("email", "account is not registered")
	ErrInvalidResetToken       = NewBadRequestError("reset_token", "invalid reset token")
	ErrAccountOnCooldown       = NewBadRequestError("credential_id", "account is on cooldown")
	ErrCredentialIdNotFound    = NewBadRequestError("credential_id", "credential id is unknown")
	ErrNonUserCannotVerify     = NewBadRequestError("credential_id", "non user cannot verify")
	ErrInvalidVerifyToken      = NewBadRequestError("verification_token", "invalid verify token")
	ErrAuth                    = NewBadRequestError("credential_id", "auth error")
	ErrEmailRegistered         = NewBadRequestError("email", "email already registered")
	ErrSipaNumberRegistered    = NewBadRequestError("sipa_number", "sipa number already registered")
	ErrPhoneNumberRegistered   = NewBadRequestError("phone_number", "phone number already registered")
	ErrPharmacistNotFound      = NewBadRequestError("id", "pharmacist doesn't exist")
	ErrUserAlreadyVerified     = NewBadRequestError("is_verified", "user is already verified")
	ErrPharmacistAssigned      = NewBadRequestError("id", "pharmacist is assigned to a pharmacy")
	ErrPharmacistNotAssigned   = NewBadRequestError("id", "pharmacist is not assigned to any pharmacy")
	ErrPharmacyNotFound        = NewBadRequestError("pharmacy_id", "pharmacy doesn't exist")
	ErrPharmacyNoPharmacist    = NewBadRequestError("pharmacy_id", "a pharmacy must have at least one pharmacist assigned to it")
	ErrPartnerNotFound         = NewBadRequestError("partner_id", "partner doesn't exist")
	ErrFileExtension           = NewBadRequestError("file", "extension not supported")
	ErrFileSize                = NewBadRequestError("file", "file size too big")
	ErrProvinceNotFound        = NewBadRequestError("province_id", "province doesn't exist")
	ErrCityNotFound            = NewBadRequestError("city_id", "city doesn't exist")
	ErrDistrictNotFound        = NewBadRequestError("district_id", "district doesn't exist")
	ErrPharmacyHasPharmacist   = NewBadRequestError("id", "pharmacy have pharmacist(s) assigned to it")
	ErrPharmacyHasOrder        = NewBadRequestError("id", "pharmacy have order(s)")
	ErrMoreThanCurrentYear     = NewBadRequestError("year_founded", "year founded must not be more than current year")
	ErrInvalidOperationalHour  = NewBadRequestError("operational_hour_start", "start operational hour must be earlier than end operational hour")
	ErrProductCategoryExist    = NewBadRequestError("product_category_id", "product category with the same name already exists")
	ErrProductCategoryNotFound = NewBadRequestError("product_category_id", "product category not found")
	ErrProductCategoryIsUsed   = NewBadRequestError("product_category_id", "product category is being used on one or more products")
	ErrProductReqNotFound      = NewBadRequestError("product_id", "product classification, form, manufacturer, or category doesn't exist")
	ErrProductExist            = NewBadRequestError("product_id", "product with the same name, generic name, and manufacturer already exists")
	ErrProductNotFound         = NewBadRequestError("product_id", "product not found")
	ErrProductBought           = NewBadRequestError("product_id", "product has been bought")
	ErrPharmacistNoAccess      = NewBadRequestError("pharmacy_id", "pharmacist is not assigned to this pharmacy")
	ErrCatalogCreated          = NewBadRequestError("catalog_id", "catalog already exist")
	ErrCatalogNotFound         = NewBadRequestError("id", "catalog doesn't exist")
	ErrCartItemNotFound        = NewBadRequestError("cart_item_id", "cart item not found")
	ErrAddressNotFound         = NewBadRequestError("address_id", "address doesn't exist")
	ErrAddressNotAssociated    = NewBadRequestError("address_id", "address is not associated with user")
	ErrCartUnavailableItems    = NewBadRequestError("cart", "cart has unavailable item(s)")
	ErrCartEmpty               = NewBadRequestError("cart", "cart is empty")
	ErrOrderGroupNotFound      = NewBadRequestError("id", "order group doesn't exist")
	ErrOrderGroupNotAssociated = NewBadRequestError("id", "order group is not associated with user")
	ErrAddressIsActive         = NewBadRequestError("is_active", "cannot delete active address")
	ErrPartnerInactive         = NewBadRequestError("is_active", "partner is inactive")
	ErrOrderNotFound           = NewBadRequestError("order_id", "order doesn't exist")
	ErrOrderCancelUser         = NewBadRequestError("status", "order can only be canceled before it is paid")
	ErrOrderUserNotAssociated  = NewBadRequestError("user", "user not associate with order")
	ErrStatusBackward          = NewBadRequestError("status", "status can't move backwards")
	ErrOrderCancel             = NewBadRequestError("status", "order can only be canceled before it is sent")
	ErrOrderConfirmUser        = NewBadRequestError("status", "order can only be confirmed when it is already sent")
	ErrCatalogDelete           = NewBadRequestError("id", "catalog can only be deleted when there is no order")
)
