package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandlerImpl) GetPharmaciesHandler(ctx *gin.Context) {
	var params dto.PharmacyOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacies, options, err := h.pharmacyService.GetPharmaciesService(ctx, mapper.PharmacyOptionsDtoToEntity(&params))
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.PharmacyGetResponse, len(pharmacies))
	for i, p := range pharmacies {
		res[i] = mapper.PharmacyToGetResponseDto(p)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.PharmacyOptionsEntityToDto(options)},
	})
}

func (h *AdminHandlerImpl) PostPharmacyHandler(ctx *gin.Context) {
	var body dto.PharmacyCreateRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacy, err := h.pharmacyService.CreatePharmacyService(ctx, mapper.PharmacyCreateRequestToEntity(&body))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PharmacyToCreateResponseDto(pharmacy),
	})
}

func (h *AdminHandlerImpl) PutPharmacyHandler(ctx *gin.Context) {
	var body dto.PharmacyUpdateRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacy, err := h.pharmacyService.UpdatePharmacyService(ctx, mapper.PharmacyUpdateRequestToEntity(&body))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PharmacyToUpdateResponseDto(pharmacy),
	})
}

func (h *AdminHandlerImpl) GetPharmacyHandler(ctx *gin.Context) {
	var params dto.PharmacyGetDetailRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacy, err := h.pharmacyService.GetPharmacyService(ctx, params.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PharmacyToGetDetailResponseDto(pharmacy),
	})
}

func (h *AdminHandlerImpl) DeletePharmacyHandler(ctx *gin.Context) {
	var params dto.PharmacyDeleteRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.pharmacyService.DeletePharmacyService(ctx, params.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
