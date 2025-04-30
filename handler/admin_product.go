package handler

import (
	"github.com/gin-gonic/gin"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/mapper"
	"net/http"
)

func (h *AdminHandlerImpl) PostProductCategoryHandler(ctx *gin.Context) {
	var body dto.ProductCategoryCreateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	productCat, err := h.productService.CreateProductCategoryService(ctx, &entity.ProductCategory{
		Name: body.Name,
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Data: dto.ProductCategoryCreateResponse{
			Id:   productCat.Id,
			Name: productCat.Name,
		},
	})
}

func (h *AdminHandlerImpl) PatchProductCategoryHandler(ctx *gin.Context) {
	var body dto.ProductCategoryUpdateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	productCat, err := h.productService.UpdateProductCategoryService(ctx, &entity.ProductCategory{
		Id:   body.Id,
		Name: body.Name,
	})
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ProductCategoryUpdateResponse{
			Id:   productCat.Id,
			Name: productCat.Name,
		},
	})
}

func (h *AdminHandlerImpl) DeleteProductCategoryHandler(ctx *gin.Context) {
	var body dto.ProductCategoryDeleteRequest
	err := ctx.ShouldBindUri(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.productService.DeleteProductCategoryService(ctx, body.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *AdminHandlerImpl) GetProductCategoriesHandler(ctx *gin.Context) {
	var params dto.ProductCategoryOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	options := mapper.ProductCategoryOptionsToEntity(&params)
	productCats, err := h.productService.GetProductCategoriesService(ctx, options)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.ProductCategoryGetResponse, len(productCats))
	for i, p := range productCats {
		res[i] = mapper.ProductCategoryToDto(p)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.ProductCategoryOptionsToDto(options)},
	})
}

func (h *AdminHandlerImpl) PostProductHandler(ctx *gin.Context) {
	var body dto.ProductCreateRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	product, err := mapper.ProductCreateRequestToEntity(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	product, err = h.productService.CreateProductService(ctx, product)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.ProductCreateResponseToDto(product),
	})
}

func (h *AdminHandlerImpl) GetProductsHandler(ctx *gin.Context) {
	var params dto.ProductOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	options := mapper.ProductOptionsToEntity(&params)
	products, err := h.productService.GetProductsService(ctx, options)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.ProductGetResponse, len(products))
	for i, p := range products {
		res[i] = mapper.ProductToDto(p)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.ProductOptionsToDto(options)},
	})
}

func (h *AdminHandlerImpl) DeleteProductHandler(ctx *gin.Context) {
	var body dto.ProductDeleteRequest
	err := ctx.ShouldBindUri(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.productService.DeleteProductService(ctx, body.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *AdminHandlerImpl) GetProductHandler(ctx *gin.Context) {
	var body dto.ProductGetRequest
	err := ctx.ShouldBindUri(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	product, err := h.productService.GetProductService(ctx, body.Id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.ProductToDto(product),
	})
}

func (h *AdminHandlerImpl) PutProductHandler(ctx *gin.Context) {
	var body dto.ProductUpdateRequest
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	product, err := mapper.ProductUpdateRequestToEntity(&body)
	if err != nil {
		ctx.Error(err)
		return
	}

	product, err = h.productService.UpdateProductService(ctx, product)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.ProductUpdateResponseToDto(product),
	})
}

func (h *AdminHandlerImpl) GetProductClassificationsHandler(ctx *gin.Context) {
	classifications, err := h.productService.GetProductClassificationsService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.ProductExtraResponse, len(classifications))
	for i, c := range classifications {
		res[i] = &dto.ProductExtraResponse{
			Id:   c.Id,
			Name: c.Name,
		}
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res},
	})
}

func (h *AdminHandlerImpl) GetProductFormsHandler(ctx *gin.Context) {
	forms, err := h.productService.GetProductFormsService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.ProductExtraResponse, len(forms))
	for i, f := range forms {
		res[i] = &dto.ProductExtraResponse{
			Id:   f.Id,
			Name: f.Name,
		}
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res},
	})
}

func (h *AdminHandlerImpl) GetProductManufacturersHandler(ctx *gin.Context) {
	manufacturers, err := h.productService.GetProductManufacturersService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.ProductExtraResponse, len(manufacturers))
	for i, m := range manufacturers {
		res[i] = &dto.ProductExtraResponse{
			Id:   m.Id,
			Name: m.Name,
		}
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res},
	})
}
