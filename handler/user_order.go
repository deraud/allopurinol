package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandlerImpl) PostOrderGroup(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.OrderGroupCreateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacies, catalogs, unavailableCatalogs, err := h.catalogService.GetCheckoutCatalogsService(ctx, content.Id, body.AddressId)
	if err != nil {
		ctx.Error(err)
		return
	}
	if len(unavailableCatalogs) > 0 {
		ctx.Error(sentinel.ErrCartUnavailableItems)
		return
	}
	if len(pharmacies) == 0 {
		ctx.Error(sentinel.ErrCartEmpty)
		return
	}

	orderGroup, err := h.orderService.CreateOrderGroupService(ctx, pharmacies, catalogs, mapper.OrderGroupCreateRequestToCheckoutDetailsEntity(&body), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.OrderGroupToCreateResponseDto(orderGroup),
	})
}

func (h *UserHandlerImpl) PutOrderGroup(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.OrderGroupUpdateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	orderGroup, err := h.orderService.UpdateOrderGroupService(ctx, mapper.OrderGroupUpdateRequestEntity(&body), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.OrderGroupToUpdateResponseDto(orderGroup),
	})
}

func (h *UserHandlerImpl) GetPendingOrders(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var params dto.PendingOrderOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	options := mapper.PendingOrderOptionsToEntity(&params)
	groupOrders, err := h.orderService.GetPendingOrdersService(ctx, options, content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.PendingOrderGroupGetResponse, len(groupOrders))
	for i, p := range groupOrders {
		res[i] = mapper.PendingOrderGroupToDto(p)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.PendingOrderOptionsToDto(options)},
	})
}

func (h *UserHandlerImpl) PatchUserOrderCanceled(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var param dto.UserOrderUpdateRequest
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.Error(err)
		return
	}

	order, err := h.orderService.UpdateUserOrderCanceledService(ctx, mapper.UserOrderUpdateRequestToEntity(&param), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.OrderToUserOrderUpdateResponseDto(order),
	})
}

func (h *UserHandlerImpl) PatchUserOrderConfirmed(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var param dto.UserOrderUpdateRequest
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.Error(err)
		return
	}

	order, err := h.orderService.UpdateUserOrderConfirmedService(ctx, mapper.UserOrderUpdateRequestToEntity(&param), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.OrderToUserOrderUpdateResponseDto(order),
	})
}

func (h *UserHandlerImpl) GetUserOrders(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.UserOrderOptionsRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	orders, options, err := h.orderService.GetUserOrdersService(ctx, mapper.UserOrderOptionRequestToEntity(&body), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.UserOrderGetResponse, len(orders))
	for i, o := range orders {
		res[i] = mapper.OrderToUserOrderGetResponseDto(o)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.UserOrderOptionToDto(options)},
	})
}
