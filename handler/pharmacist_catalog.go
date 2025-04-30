package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"healthcare-allopurinol/sentinel"
	"healthcare-allopurinol/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PharmacistHandlerImpl) PostCatalog(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.CatalogCreateRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacy, err := h.catalogService.CreateCatalogService(ctx, content.Id, mapper.CatalogCreateRequestDtoToEntity(&body))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.CatalogEntityToDto(pharmacy),
	})
}

func (h *PharmacistHandlerImpl) PatchCatalog(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var body dto.CatalogUpdateRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	pharmacy, err := h.catalogService.UpdateCatalogService(ctx, content.Id, mapper.CatalogUpdateRequestDtoToEntity(&body))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.CatalogEntityToDto(pharmacy),
	})
}

func (h *PharmacistHandlerImpl) DeleteCatalog(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var params dto.CatalogDeleteRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.catalogService.DeleteCatalogService(ctx, content.Id, params.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *PharmacistHandlerImpl) GetCatalogs(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var params dto.CatalogOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	catalogs, options, err := h.catalogService.GetCatalogsService(ctx, mapper.CatalogOptionsToEntity(&params), content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.CatalogGetResponse, len(catalogs))
	for i, c := range catalogs {
		res[i] = mapper.CatalogToGetResponseDto(c)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.CatalogOptionsToDto(options)},
	})
}

func (h *PharmacistHandlerImpl) GetCatalog(ctx *gin.Context) {
	c, ok := ctx.Get("userData")
	if !ok {
		ctx.Error(sentinel.ErrAuth)
	}
	content := c.(*utility.ClaimsContent)

	var params dto.CatalogDeleteRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	catalog, err := h.catalogService.GetCatalogService(ctx, params.Id, content.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.CatalogToGetDetailResponseDto(catalog),
	})
}
