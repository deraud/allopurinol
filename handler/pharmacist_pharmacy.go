package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PharmacistHandlerImpl) PutPharmacy(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.PharmacistPharmacyUpdateRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacy, err := h.pharmacyService.UpdatePharmacyFromPharmacistService(ctx, content.Id, mapper.PharmacistPharmacyUpdateRequestToEntity(&body))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PharmacyToUpdateResponseDto(pharmacy),
	})
}

func (h *PharmacistHandlerImpl) GetPharmacy(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	pharmacy, err := h.pharmacyService.GetPharmacyFromPharmacistService(ctx, content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PharmacyToGetDetailResponseDto(pharmacy),
	})
}
