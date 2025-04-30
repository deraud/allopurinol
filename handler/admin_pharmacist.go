package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandlerImpl) PostPharmacistHandler(ctx *gin.Context) {
	var body dto.PharmacistCreateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacist, credential, err := h.pharmacistService.CreatePharmacistService(
		ctx,
		mapper.PharmacistCreateRequestToCredentialEntity(&body),
		mapper.PharmacistCreateRequestToPharmacistEntity(&body),
	)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.PharmacistToPharmacistCreateResponseDto(pharmacist, credential),
	})
}

func (h *AdminHandlerImpl) PatchPharmacistHandler(ctx *gin.Context) {
	var body dto.PharmacistUpdateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacist, err := h.pharmacistService.UpdatePharmacistService(ctx, mapper.PharmacistUpdateRequestToEntity(&body))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PharmacistToPharmacistUpdateResponseDto(pharmacist),
	})
}

func (h *AdminHandlerImpl) DeletePharmacistHandler(ctx *gin.Context) {
	var body dto.PharmacistDeleteRequest
	err := ctx.ShouldBindUri(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.pharmacistService.DeletePharmacistService(ctx, body.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *AdminHandlerImpl) GetPharmacistsHandler(ctx *gin.Context) {
	var params dto.PharmacistOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacists, emails, options, err := h.pharmacistService.GetPharmacistsService(ctx, mapper.PharmacistOptionsToEntity(&params))
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.PharmacistGetResponse, len(pharmacists))
	for i, p := range pharmacists {
		res[i] = mapper.PharmacistToDto(p, emails[i])
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.PharmacistOptionsToDto(options)},
	})
}

func (h *AdminHandlerImpl) GetPharmacistHandler(ctx *gin.Context) {
	var params dto.PharmacistGetDetailRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacist, email, err := h.pharmacistService.GetPharmacistService(ctx, params.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PharmacistToDto(pharmacist, email),
	})
}
