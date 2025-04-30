package mapper

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"strconv"

	"github.com/shopspring/decimal"
)

func CatalogCreateRequestDtoToEntity(dto *dto.CatalogCreateRequest) *entity.Catalog {
	price, _ := decimal.NewFromString(dto.Price)
	return &entity.Catalog{
		PharmacyId: dto.PharmacyId,
		ProductId:  dto.ProductId,
		Stock:      dto.Stock,
		Price:      price,
	}
}

func CatalogEntityToDto(entity *entity.Catalog) *dto.Catalog {
	return &dto.Catalog{
		Id:         entity.Id,
		PharmacyId: entity.PharmacyId,
		ProductId:  entity.ProductId,
		Stock:      entity.Stock,
		Price:      entity.Price,
		IsActive:   entity.IsActive,
	}
}

func CatalogUpdateRequestDtoToEntity(dto *dto.CatalogUpdateRequest) *entity.Catalog {
	price, _ := decimal.NewFromString(dto.Price)
	return &entity.Catalog{
		Id:       dto.Id,
		Stock:    dto.Stock,
		Price:    price,
		IsActive: *dto.IsActive,
	}
}

func CatalogOptionsToEntity(dto *dto.CatalogOptionsRequest) *entity.CatalogOptions {
	var (
		searchBy         string = constant.CATALOG_DEFAULT_SEARCH_BY
		sortBy           string = constant.CATALOG_DEFAULT_SORT_BY
		sortOrder        string = constant.CATALOG_DEFAULT_SORT_ORDER
		manufacturerId   *int64
		classificationId *int64
		formId           *int64
		isActive         *bool
		page             int = constant.DEFAULT_PAGE
		limit            int = constant.DEFAULT_LIMIT
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
	if dto.ManufacturerId != "" {
		value, _ := strconv.ParseInt(dto.ManufacturerId, 10, 64)
		manufacturerId = &value
	}
	if dto.ClassificationId != "" {
		value, _ := strconv.ParseInt(dto.ClassificationId, 10, 64)
		classificationId = &value
	}
	if dto.FormId != "" {
		value, _ := strconv.ParseInt(dto.FormId, 10, 64)
		formId = &value
	}
	if dto.IsActive != "" {
		value, _ := strconv.ParseBool(dto.IsActive)
		isActive = &value
	}
	if dto.Page != "" {
		page, _ = strconv.Atoi(dto.Page)
	}
	if dto.Limit != "" {
		limit, _ = strconv.Atoi(dto.Limit)
	}

	return &entity.CatalogOptions{
		SearchBy:         searchBy,
		SearchValue:      dto.SearchValue,
		SortBy:           sortBy,
		SortOrder:        sortOrder,
		ManufacturerId:   manufacturerId,
		ClassificationId: classificationId,
		FormId:           formId,
		IsActive:         isActive,
		Page:             page,
		Limit:            limit,
	}
}

func CatalogOptionsToDto(entity *entity.CatalogOptions) *dto.CatalogOptionsResponse {
	return &dto.CatalogOptionsResponse{
		Search:           dto.SearchOptions{Column: entity.SearchBy, Value: entity.SearchValue},
		Sort:             dto.SortOption{Column: entity.SortBy, Order: entity.SortOrder},
		ManufacturerId:   entity.ManufacturerId,
		ClassificationId: entity.ClassificationId,
		FormId:           entity.FormId,
		IsActive:         entity.IsActive,
		Page:             entity.Page,
		Limit:            entity.Limit,
		TotalRow:         entity.TotalRows,
	}
}

func CatalogToGetResponseDto(entity *entity.Catalog) *dto.CatalogGetResponse {
	return &dto.CatalogGetResponse{
		Id:             entity.Id,
		Name:           entity.Name,
		GenericName:    entity.GenericName,
		Manufacturer:   entity.Manufacturer,
		Classification: entity.Classification,
		Form:           entity.Form,
		Stock:          entity.Stock,
		IsActive:       entity.IsActive,
	}
}

func CatalogToGetDetailResponseDto(entity *entity.Catalog) *dto.CatalogGetDetailResponse {
	return &dto.CatalogGetDetailResponse{
		Id:             entity.Id,
		Stock:          entity.Stock,
		Price:          entity.Price,
		IsActive:       entity.IsActive,
		Name:           entity.Name,
		GenericName:    entity.GenericName,
		Manufacturer:   entity.Manufacturer,
		Classification: entity.Classification,
		Form:           entity.Form,
		Description:    entity.Description,
		UnitInPack:     entity.UnitInPack,
		SellingUnit:    entity.SellingUnit,
		Image:          entity.Image,
	}
}

func AvailableCatalogOptionsToEntity(dto *dto.AvailableCatalogOptionsRequest) *entity.AvailableCatalogOptions {
	var (
		searchBy   string = constant.CATALOG_DEFAULT_SEARCH_BY
		categoryId *int64
		page       int = constant.DEFAULT_PAGE
		limit      int = constant.DEFAULT_LIMIT
	)

	if dto.SearchBy != "" {
		searchBy = dto.SearchBy
	}
	if dto.CategoryId != "" {
		value, _ := strconv.ParseInt(dto.CategoryId, 10, 64)
		categoryId = &value
	}
	if dto.Page != "" {
		page, _ = strconv.Atoi(dto.Page)
	}
	if dto.Limit != "" {
		limit, _ = strconv.Atoi(dto.Limit)
	}

	return &entity.AvailableCatalogOptions{
		SearchBy:    searchBy,
		SearchValue: dto.SearchValue,
		CategoryId:  categoryId,
		Page:        page,
		Limit:       limit,
	}
}

func CatalogToAvailableCatalogGetResponseDto(entity *entity.Catalog) *dto.AvailableCatalogGetResponse {
	return &dto.AvailableCatalogGetResponse{
		Id:    entity.Id,
		Price: entity.Price,
		Stock: entity.Stock,
		Product: dto.ProductCatalogGetDetailResponse{
			Id:          entity.Product.Id,
			Name:        entity.Product.Name,
			Image:       entity.Product.ImageLink,
			SellingUnit: entity.Product.SellingUnit,
		},
	}
}

func AvailableCatalogOptionsToDto(entity *entity.AvailableCatalogOptions) *dto.AvailableCatalogOptionsResponse {
	return &dto.AvailableCatalogOptionsResponse{
		Search:   dto.SearchOptions{Column: entity.SearchBy, Value: entity.SearchValue},
		Page:     entity.Page,
		Limit:    entity.Limit,
		TotalRow: entity.TotalRows,
	}
}

func CatalogToAvailableCatalogGetDetailResponseDto(entity *entity.Catalog) *dto.AvailableCatalogGetDetailResponse {
	return &dto.AvailableCatalogGetDetailResponse{
		Id:       entity.Id,
		Price:    entity.Price,
		Stock:    entity.Stock,
		Product:  *ProductToCatalogGetDetailResponseDto(entity.Product),
		Pharmacy: *PharmacyToCatalogGetDetailResponseDto(entity.Pharmacy),
	}
}

func CatalogToCheckoutCatalogDetailDto(entity *entity.Catalog) *dto.CheckoutCatalogDetail {
	var price *decimal.Decimal
	if !entity.Price.Equal(decimal.NewFromInt(0)) {
		price = &entity.Price
	}
	return &dto.CheckoutCatalogDetail{
		Id:       entity.Id,
		Price:    price,
		Stock:    entity.Stock,
		Quantity: entity.Quantity,
		Product:  *ProductToCheckoutCatalogDto(entity.Product),
	}
}

func CatalogToCheckoutCatalogGetDetailResponseDto(pharmacy *entity.Pharmacy, catalogs []*entity.Catalog) *dto.CheckoutCatalogGetResponse {
	catalogsDto := make([]*dto.CheckoutCatalogDetail, len(catalogs))
	for i, c := range catalogs {
		catalogsDto[i] = CatalogToCheckoutCatalogDetailDto(c)
	}

	return &dto.CheckoutCatalogGetResponse{
		Pharmacy: *PharmacyToCheckoutDetailDto(pharmacy),
		Catalogs: catalogsDto,
	}
}
