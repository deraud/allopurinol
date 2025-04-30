package handler

import (
	"github.com/gin-gonic/gin"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"
)

func postCartItem(h *UserHandlerImpl, ctx *gin.Context, increment bool) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.CartItemRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	if increment {
		err = h.cartService.IncrementCartItemService(ctx, &entity.CartItem{ProductId: body.ProductId}, content.Id)
	} else {
		err = h.cartService.DecrementCartItemService(ctx, &entity.CartItem{ProductId: body.ProductId}, content.Id)
	}
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandlerImpl) PostCartItemIncrement(ctx *gin.Context) {
	postCartItem(h, ctx, true)
}

func (h *UserHandlerImpl) PostCartItemDecrement(ctx *gin.Context) {
	postCartItem(h, ctx, false)
}

func (h *UserHandlerImpl) PostCartItemSet(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.CartItemSetRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.cartService.UpdateCartItemService(ctx, &entity.CartItem{ProductId: body.ProductId}, body.Quantity, content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandlerImpl) PostCartItemRemove(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.CartItemRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.cartService.DeleteCartItemService(ctx, &entity.CartItem{ProductId: body.ProductId}, content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandlerImpl) GetCartItems(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	items, err := h.cartService.GetCartItemsService(ctx, content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.CartItemResponse, len(items))
	for i, item := range items {
		res[i] = &dto.CartItemResponse{
			Id:       item.ProductId,
			Quantity: item.Quantity,
		}
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res},
	})
}

func (h *UserHandlerImpl) PostShippingCost(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.ShippingRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	_, catMap, _, err := h.catalogService.GetCheckoutCatalogsService(ctx, content.Id, body.AddressId)
	if err != nil {
		ctx.Error(err)
		return
	}
	catList, exist := catMap[body.PharmacyId]
	if !exist {
		ctx.Error(err)
		return
	}

	shipping, err := h.checkoutService.GetShippingCosts(ctx, &entity.Shipping{
		PharmacyId:    body.PharmacyId,
		ShippingCosts: []*entity.ShippingCost{},
		Catalogs:      catList,
	}, content.Id, body.AddressId)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.ShippingCostResponse, len(shipping.ShippingCosts))
	for i, item := range shipping.ShippingCosts {
		res[i] = &dto.ShippingCostResponse{
			Id:   item.MethodId,
			Name: item.Method,
			Cost: item.Cost.String(),
		}
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: &dto.ShippingResponse{
			ShippingCosts: res,
		},
	})
}
