package mapper

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"strconv"

	"github.com/shopspring/decimal"
)

func ProductCategoryOptionsToEntity(dto *dto.ProductCategoryOptionsRequest) *entity.ProductCategoryOptions {
	var (
		searchBy  = constant.PRODUCT_CATEGORY_DEFAULT_SEARCH_BY
		sortBy    = constant.PRODUCT_CATEGORY_DEFAULT_SORT_BY
		sortOrder = constant.PRODUCT_CATEGORY_DEFAULT_SORT_ORDER
		page      = constant.DEFAULT_PAGE
		limit     = constant.DEFAULT_LIMIT
	)

	if dto.SearchBy != "" {
		searchBy = dto.SearchBy
	}
	if dto.SortBy != "" {
		sortBy = dto.SortBy
	}
	if dto.SortOrder != "" {
		sortOrder = dto.SortOrder
	}
	if dto.Page != "" {
		page, _ = strconv.Atoi(dto.Page)
	}
	if dto.Limit != "" {
		limit, _ = strconv.Atoi(dto.Limit)
	}

	return &entity.ProductCategoryOptions{
		SearchBy:    searchBy,
		SearchValue: dto.SearchValue,
		SortBy:      sortBy,
		SortOrder:   sortOrder,
		Page:        page,
		Limit:       limit,
	}
}

func ProductCategoryToDto(entity *entity.ProductCategory) *dto.ProductCategoryGetResponse {
	return &dto.ProductCategoryGetResponse{
		Id:   entity.Id,
		Name: entity.Name,
	}
}

func ProductCategoryOptionsToDto(entity *entity.ProductCategoryOptions) *dto.ProductCategoryOptionsResponse {
	return &dto.ProductCategoryOptionsResponse{
		Search:   dto.SearchOptions{Column: entity.SearchBy, Value: entity.SearchValue},
		Sort:     dto.SortOption{Column: entity.SortBy, Order: entity.SortOrder},
		Page:     entity.Page,
		Limit:    entity.Limit,
		TotalRow: entity.TotalRows,
	}
}

func ProductCreateRequestToEntity(dto *dto.ProductCreateRequest) (*entity.Product, error) {
	weight, err := decimal.NewFromString(dto.Weight)
	if err != nil {
		return nil, err
	}
	height, err := decimal.NewFromString(dto.Height)
	if err != nil {
		return nil, err
	}
	length, err := decimal.NewFromString(dto.Length)
	if err != nil {
		return nil, err
	}
	width, err := decimal.NewFromString(dto.Width)
	if err != nil {
		return nil, err
	}

	if dto.UnitInPack != nil && *dto.UnitInPack == 0 {
		dto.UnitInPack = nil
	}
	if dto.SellingUnit != nil && *dto.SellingUnit == "" {
		dto.SellingUnit = nil
	}

	return &entity.Product{
		ClassificationId: dto.ClassificationId,
		FormId:           dto.FormId,
		ManufacturerId:   dto.ManufacturerId,
		CategoryIds:      dto.CategoryIds,
		Name:             dto.Name,
		GenericName:      dto.GenericName,
		Description:      dto.Description,
		UnitInPack:       dto.UnitInPack,
		SellingUnit:      dto.SellingUnit,
		Weight:           weight,
		Height:           height,
		Length:           length,
		Width:            width,
		Image:            dto.Image,
		ImageLink:        dto.Image,
		IsActive:         dto.IsActive,
	}, nil
}

func ProductCreateResponseToDto(entity *entity.Product) *dto.ProductCreateResponse {
	return &dto.ProductCreateResponse{
		Id:               entity.Id,
		ClassificationId: entity.ClassificationId,
		FormId:           entity.FormId,
		ManufacturerId:   entity.ManufacturerId,
		CategoryIds:      entity.CategoryIds,
		Name:             entity.Name,
		GenericName:      entity.GenericName,
		Description:      entity.Description,
		UnitInPack:       entity.UnitInPack,
		SellingUnit:      entity.SellingUnit,
		Weight:           entity.Weight.String(),
		Height:           entity.Height.String(),
		Length:           entity.Length.String(),
		Width:            entity.Width.String(),
		Image:            entity.ImageLink,
		IsActive:         entity.IsActive,
	}
}

func ProductOptionsToEntity(dto *dto.ProductOptionsRequest) *entity.ProductOptions {
	var (
		searchBy  = constant.PRODUCT_DEFAULT_SEARCH_BY
		sortBy    = constant.PRODUCT_DEFAULT_SORT_BY
		sortOrder = constant.PRODUCT_DEFAULT_SORT_ORDER
		page      = constant.DEFAULT_PAGE
		limit     = constant.DEFAULT_LIMIT
	)

	if dto.SearchBy != "" {
		searchBy = dto.SearchBy
	}
	if dto.SortBy != "" {
		sortBy = dto.SortBy
	}
	if dto.SortOrder != "" {
		sortOrder = dto.SortOrder
	}
	if dto.Page != "" {
		page, _ = strconv.Atoi(dto.Page)
	}
	if dto.Limit != "" {
		limit, _ = strconv.Atoi(dto.Limit)
	}

	return &entity.ProductOptions{
		SearchBy:    searchBy,
		SearchValue: dto.SearchValue,
		FilterBy:    dto.FilterBy,
		FilterValue: dto.FilterValue,
		SortBy:      sortBy,
		SortOrder:   sortOrder,
		Page:        page,
		Limit:       limit,
	}
}

func ProductToDto(entity *entity.Product) *dto.ProductGetResponse {
	var categories []dto.ProductCategoryGetResponse
	for i, category := range entity.Categories {
		categories = append(categories, dto.ProductCategoryGetResponse{
			Id:   entity.CategoryIds[i],
			Name: category,
		})
	}
	var stock int
	if entity.Stock != nil {
		stock = *entity.Stock
	}
	var usage int
	if entity.Usage != nil {
		usage = *entity.Usage
	}
	var weight string
	if !entity.Weight.Equal(decimal.Zero) {
		weight = entity.Weight.String()
	}
	var height string
	if !entity.Height.Equal(decimal.Zero) {
		height = entity.Height.String()
	}
	var width string
	if !entity.Width.Equal(decimal.Zero) {
		width = entity.Width.String()
	}
	var length string
	if !entity.Length.Equal(decimal.Zero) {
		length = entity.Length.String()
	}
	var form *dto.ProductFormGetResponse
	if entity.FormId != nil && entity.Form != nil {
		form = &dto.ProductFormGetResponse{
			Id:   entity.FormId,
			Name: entity.Form,
		}
	}

	return &dto.ProductGetResponse{
		Id: entity.Id,
		Classification: dto.ProductExtraGetResponse{
			Id:   entity.ClassificationId,
			Name: entity.Classification,
		},
		Form: form,
		Manufacturer: dto.ProductExtraGetResponse{
			Id:   entity.ManufacturerId,
			Name: entity.Manufacturer,
		},
		Categories:  categories,
		Name:        entity.Name,
		GenericName: entity.GenericName,
		Stock:       stock,
		Usage:       usage,
		Description: entity.Description,
		UnitInPack:  entity.UnitInPack,
		SellingUnit: entity.SellingUnit,
		Weight:      weight,
		Height:      height,
		Width:       width,
		Length:      length,
		Image:       entity.ImageLink,
		IsActive:    entity.IsActive,
	}
}

func ProductOptionsToDto(entity *entity.ProductOptions) *dto.ProductOptionsResponse {
	return &dto.ProductOptionsResponse{
		Search:   dto.SearchOptions{Column: entity.SearchBy, Value: entity.SearchValue},
		Sort:     dto.SortOption{Column: entity.SortBy, Order: entity.SortOrder},
		Page:     entity.Page,
		Limit:    entity.Limit,
		TotalRow: entity.TotalRows,
	}
}

func ProductToCatalogGetDetailResponseDto(entity *entity.Product) *dto.ProductCatalogGetDetailResponse {
	var categories []dto.ProductCategoryGetResponse
	for i, category := range entity.Categories {
		categories = append(categories, dto.ProductCategoryGetResponse{
			Id:   entity.CategoryIds[i],
			Name: category,
		})
	}

	form := ""
	if entity.Form != nil {
		form = *entity.Form
	}

	return &dto.ProductCatalogGetDetailResponse{
		Id:             entity.Id,
		Name:           entity.Name,
		GenericName:    entity.GenericName,
		Manufacturer:   entity.Manufacturer,
		Classification: entity.Classification,
		Form:           form,
		Description:    entity.Description,
		UnitInPack:     entity.UnitInPack,
		SellingUnit:    entity.SellingUnit,
		Weight:         entity.Weight.String(),
		Height:         entity.Height.String(),
		Length:         entity.Length.String(),
		Width:          entity.Width.String(),
		Categories:     categories,
		Image:          entity.ImageLink,
	}
}

func ProductUpdateRequestToEntity(dto *dto.ProductUpdateRequest) (*entity.Product, error) {
	weight, err := decimal.NewFromString(dto.Weight)
	if err != nil {
		return nil, err
	}
	height, err := decimal.NewFromString(dto.Height)
	if err != nil {
		return nil, err
	}
	length, err := decimal.NewFromString(dto.Length)
	if err != nil {
		return nil, err
	}
	width, err := decimal.NewFromString(dto.Width)
	if err != nil {
		return nil, err
	}

	if dto.UnitInPack != nil && *dto.UnitInPack == 0 {
		dto.UnitInPack = nil
	}
	if dto.SellingUnit != nil && *dto.SellingUnit == "" {
		dto.SellingUnit = nil
	}

	return &entity.Product{
		Id:               dto.Id,
		ClassificationId: dto.ClassificationId,
		FormId:           dto.FormId,
		ManufacturerId:   dto.ManufacturerId,
		CategoryIds:      dto.CategoryIds,
		Name:             dto.Name,
		GenericName:      dto.GenericName,
		Description:      dto.Description,
		UnitInPack:       dto.UnitInPack,
		SellingUnit:      dto.SellingUnit,
		Weight:           weight,
		Height:           height,
		Length:           length,
		Width:            width,
		Image:            dto.Image,
		ImageLink:        dto.Image,
		IsActive:         dto.IsActive,
	}, nil
}

func ProductUpdateResponseToDto(entity *entity.Product) *dto.ProductUpdateResponse {
	return &dto.ProductUpdateResponse{
		Id:               entity.Id,
		ClassificationId: entity.ClassificationId,
		FormId:           entity.FormId,
		ManufacturerId:   entity.ManufacturerId,
		CategoryIds:      entity.CategoryIds,
		Name:             entity.Name,
		GenericName:      entity.GenericName,
		Description:      entity.Description,
		UnitInPack:       entity.UnitInPack,
		SellingUnit:      entity.SellingUnit,
		Weight:           entity.Weight.String(),
		Height:           entity.Height.String(),
		Length:           entity.Length.String(),
		Width:            entity.Width.String(),
		Image:            entity.ImageLink,
		IsActive:         entity.IsActive,
	}
}

func ProductToCheckoutCatalogDto(entity *entity.Product) *dto.ProductCheckoutCatalog {
	return &dto.ProductCheckoutCatalog{
		Id:    entity.Id,
		Name:  entity.Name,
		Image: entity.ImageLink,
	}
}
