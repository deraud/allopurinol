package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PharmacistHandlerImpl) GetPharmacyOrders(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.PharmacyOrderOptionsRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	orders, options, err := h.orderService.GetPharmacyOrdersService(ctx, mapper.PharmacyOrderOptionRequestToEntity(&body), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.PharmacyOrderGetResponse, len(orders))
	for i, o := range orders {
		res[i] = mapper.OrderToPharmacyOrderGetResponseDto(o)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.PharmacyOrderOptionToDto(options)},
	})
}

func (h *PharmacistHandlerImpl) GetPharmacyOrder(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var param dto.PharmacyOrderGetDetailRequest
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.Error(err)
		return
	}

	order, err := h.orderService.GetPharmacyOrderService(ctx, param.Id, content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.OrderToPharmacyOrderGetDetailResponseDto(order),
	})
}

func (h *PharmacistHandlerImpl) PatchPharmacyOrderShipped(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var param dto.PharmacyOrderUpdateRequest
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.Error(err)
		return
	}

	order, err := h.orderService.UpdatePharmacyOrderShippedService(ctx, mapper.PharmacyOrderUpdateRequestToEntity(&param), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.OrderToPharmacyOrderUpdateResponseDto(order),
	})
}

func (h *PharmacistHandlerImpl) PatchPharmacyOrderCanceled(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var param dto.PharmacyOrderUpdateRequest
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.Error(err)
		return
	}

	order, err := h.orderService.UpdatePharmacyOrderCanceledService(ctx, mapper.PharmacyOrderUpdateRequestToEntity(&param), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.OrderToPharmacyOrderUpdateResponseDto(order),
	})
}
