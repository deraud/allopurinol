package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandlerImpl) PostPartnerHandler(ctx *gin.Context) {
	var body dto.PartnerCreateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	partner, err := h.partnerService.CreatePartnerService(
		ctx,
		mapper.PartnerCreateRequestDtoToEntityPartner(&body),
	)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.PartnerEntityToDtoPartnerCreateResponse(partner),
	})
}

func (h *AdminHandlerImpl) PatchPartnerHandler(ctx *gin.Context) {
	var body dto.PartnerUpdateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	partner, err := h.partnerService.UpdatePartnerService(ctx, mapper.PartnerUpdateRequestDtoToEntity(&body))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PartnerEntityToDtoPartnerUpdateResponse(partner),
	})
}

func (h *AdminHandlerImpl) DeletePartnerHandler(ctx *gin.Context) {
	var body dto.PartnerDeleteRequest
	err := ctx.ShouldBindUri(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.partnerService.DeletePartnerService(ctx, body.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *AdminHandlerImpl) GetPartnersHandler(ctx *gin.Context) {
	var params dto.PartnerOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	partners, options, err := h.partnerService.GetPartnersService(ctx, mapper.PartnerOptionsDtoToEntity(&params))

	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.PartnerGetResponse, len(partners))
	for i, p := range partners {
		res[i] = mapper.PartnerEntityToDto(p)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.PartnerOptionsEntityToDto(options)},
	})
}

func (h *AdminHandlerImpl) GetPartnerHandler(ctx *gin.Context) {
	var params dto.PartnerGetDetailRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	partner, err := h.partnerService.GetPartnerService(ctx, params.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.PartnerEntityToDto(partner),
	})
}
