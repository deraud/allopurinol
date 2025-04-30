package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandlerImpl) GetOrders(ctx *gin.Context) {
	var param dto.OrderOptionsRequest
	err := ctx.ShouldBind(&param)
	if err != nil {
		ctx.Error(err)
		return
	}

	orders, options, err := h.orderService.GetOrdersService(ctx, mapper.OrderOptionsRequestToEntity(&param))
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.OrderGetResponse, len(orders))
	for i, o := range orders {
		res[i] = mapper.OrderToGetResponseDto(o)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.OrderOptionsToDto(options)},
	})
}

func (h *AdminHandlerImpl) GetOrder(ctx *gin.Context) {
	var param dto.OrderGetDetailRequest
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.Error(err)
		return
	}

	order, err := h.orderService.GetOrderService(ctx, param.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.OrderToGetDetailResponseDto(order),
	})
}
