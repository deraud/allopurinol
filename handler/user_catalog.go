package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandlerImpl) GetCatalogs(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var params dto.AvailableCatalogOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	catalogs, options, err := h.catalogService.GetAvailableCatalogsService(ctx, mapper.AvailableCatalogOptionsToEntity(&params), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.AvailableCatalogGetResponse, len(catalogs))
	for i, c := range catalogs {
		res[i] = mapper.CatalogToAvailableCatalogGetResponseDto(c)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.AvailableCatalogOptionsToDto(options)},
	})
}

func (h *UserHandlerImpl) GetMostBoughtCatalogs(ctx *gin.Context) {
	catalogs, err := h.catalogService.GetMostBoughtCatalogsService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.AvailableCatalogGetResponse, len(catalogs))
	for i, c := range catalogs {
		res[i] = mapper.CatalogToAvailableCatalogGetResponseDto(c)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: res,
	})
}

func (h *UserHandlerImpl) GetCatalog(ctx *gin.Context) {
	var params dto.AvailableCatalogGetDetailRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	catalog, err := h.catalogService.GetAvailableCatalogService(ctx, params.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.CatalogToAvailableCatalogGetDetailResponseDto(catalog),
	})
}

func (h *UserHandlerImpl) GetCheckoutCatalog(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.CheckoutCatalogGetRequest
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

	available := make([]*dto.CheckoutCatalogGetResponse, len(pharmacies))
	for i, p := range pharmacies {
		available[i] = mapper.CatalogToCheckoutCatalogGetDetailResponseDto(p, catalogs[p.Id])
	}

	unavailable := make([]*dto.CheckoutCatalogDetail, len(unavailableCatalogs))
	for i, c := range unavailableCatalogs {
		unavailable[i] = mapper.CatalogToCheckoutCatalogDetailDto(c)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: gin.H{"available": available, "unavailable": unavailable},
	})
}
