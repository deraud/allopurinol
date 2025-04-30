package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandlerImpl) GetUserAddressesHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)
	id := content.Id

	var params dto.UserGetAddressesOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}
	userAddresses, err := h.service.GetUserAddresses(ctx, id, mapper.UserAddressesOptionsToEntity(&params))

	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.UserGetAddressesResponse, len(userAddresses))
	for i, userAddresses := range userAddresses {
		res[i] = mapper.UserAddressesEntityToDto(userAddresses)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: res,
	})
}

func (h *UserHandlerImpl) PostAddressHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)
	id := content.Id
	var body dto.UserAddressCreateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}
	address, err := h.service.CreateAddressService(ctx, mapper.UserAddressDtoToEntity(id, &body))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.UserAddressesEntityToDto(address),
	})
}

func (h *UserHandlerImpl) PatchAddressHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)
	id := content.Id
	var body dto.UserAddressUpdateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}
	address, err := h.service.UpdateAddressService(ctx, mapper.UserUpdateAddressDtoToEntity(id, &body))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.UserAddressesEntityToDto(address),
	})
}

func (h *UserHandlerImpl) DeleteUserAddressHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)
	id := content.Id
	var body dto.UserAddressDeleteRequest
	err := ctx.ShouldBindUri(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.service.DeleteAddressService(ctx, id, body.AddressId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandlerImpl) GetUserProfileHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)
	id := content.Id
	var params dto.UserGetAddressesOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	userAddresses, err := h.service.GetUserAddresses(ctx, id, mapper.UserAddressesOptionsToEntity(&params))
	if err != nil {
		ctx.Error(err)
		return
	}
	resAddress := make([]*dto.UserGetAddressesResponse, len(userAddresses))
	for i, userAddresses := range userAddresses {
		resAddress[i] = mapper.UserAddressesEntityToDto(userAddresses)
	}

	userProfile, err := h.service.GetUserProfile(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.UserProfileEntityToDto(userProfile, resAddress),
	})
}

func (h *UserHandlerImpl) PatchUserProfileHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)
	id := content.Id
	var body dto.UserUpdateProfileRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}
	user, err := h.service.UpdateUserProfileService(ctx, mapper.UserUpdateProfileRequestToEntity(id, &body))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.UserUpdateProfileEntityToDto(user),
	})
}

func (h *UserHandlerImpl) PatchUserActivateAddressHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)
	id := content.Id
	var body dto.UserActivateAddressUpdateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}
	address, err := h.service.UpdateUserActivateAddressService(ctx, mapper.UserActivateAddressDtoToEntity(id, &body))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.UserAddressEntityToActivateAddressResponse(address),
	})
}

func (h *UserHandlerImpl) PatchUserRemovePictureHandler(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)
	id := content.Id
	var body dto.UserUpdateProfileRequest
	user, err := h.service.UpdateUserRemovePictureService(ctx, mapper.UserUpdateProfileRequestToEntity(id, &body))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.UserUpdateProfileEntityToDto(user),
	})

}
