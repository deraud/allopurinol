package mapper

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"strconv"
	"strings"
)

func OrderGroupCreateRequestToCheckoutDetailsEntity(dto *dto.OrderGroupCreateRequest) *entity.CheckoutDetails {
	orderPharmacyDetails := make([]*entity.CheckoutPharmacyDetails, len(dto.OrderPharmacyDetails))
	for i, detail := range dto.OrderPharmacyDetails {
		orderPharmacyDetails[i] = CheckoutPharmacyDetailToEntity(detail)
	}

	return &entity.CheckoutDetails{
		AddressId:       dto.AddressId,
		PaymentMethodId: dto.PaymentMethodId,
		OrderDetails:    orderPharmacyDetails,
	}
}

func CheckoutPharmacyDetailToEntity(dto *dto.OrderPharmacyDetail) *entity.CheckoutPharmacyDetails {
	return &entity.CheckoutPharmacyDetails{
		PharmacyId:        dto.PharmacyId,
		LogisticPartnerId: dto.LogisticPartnerId,
	}
}

func OrderGroupToCreateResponseDto(entity *entity.OrderGroup) *dto.OrderGroupCreateResponse {
	orders := make([]*dto.OrderCreateResponse, len(entity.Orders))
	for i, o := range entity.Orders {
		orders[i] = OrderToCreateResponseDto(o)
	}

	return &dto.OrderGroupCreateResponse{
		Id:     entity.Id,
		Orders: orders,
	}
}

func OrderToCreateResponseDto(entity *entity.Order) *dto.OrderCreateResponse {
	orderItems := make([]*dto.OrderItemCreateResponse, len(entity.OrderItems))
	for i, o := range entity.OrderItems {
		orderItems[i] = OrderItemToCreateResponseDto(o)
	}

	return &dto.OrderCreateResponse{
		Id:                 entity.Id,
		AddressId:          entity.AddressId,
		PaymentMethodId:    entity.PaymentMethodId,
		PharmacyId:         entity.PharmacyId,
		LogisticPartnerId:  entity.LogisticPartnerId,
		TotalPriceProduct:  entity.TotalPriceProduct,
		TotalPriceShipping: entity.TotalPriceShipping,
		OrderItems:         orderItems,
	}
}

func OrderItemToCreateResponseDto(entity *entity.OrderItem) *dto.OrderItemCreateResponse {
	return &dto.OrderItemCreateResponse{
		Id:        entity.Id,
		CatalogId: entity.CatalogId,
		Quantity:  entity.Quantity,
	}
}

func OrderGroupUpdateRequestEntity(dto *dto.OrderGroupUpdateRequest) *entity.OrderGroup {
	return &entity.OrderGroup{
		Id:    dto.Id,
		Proof: dto.Proof,
	}
}

func OrderGroupToUpdateResponseDto(entity *entity.OrderGroup) *dto.OrderGroupUpdateResponse {
	return &dto.OrderGroupUpdateResponse{
		Id:    entity.Id,
		Proof: entity.Proof,
	}
}

func PendingOrderOptionsToEntity(dto *dto.PendingOrderOptionsRequest) *entity.PendingOrderOptions {
	var (
		sortBy    = constant.PENDING_ORDER_DEFAULT_SORT_BY
		sortOrder = constant.PENDING_ORDER_DEFAULT_SORT_ORDER
		page      = constant.DEFAULT_PAGE
		limit     = constant.DEFAULT_LIMIT
	)

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

	return &entity.PendingOrderOptions{
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Page:      page,
		Limit:     limit,
	}
}

func PendingOrderOptionsToDto(entity *entity.PendingOrderOptions) *dto.PendingOrderOptionsResponse {
	return &dto.PendingOrderOptionsResponse{
		Sort:     dto.SortOption{Column: entity.SortBy, Order: entity.SortOrder},
		Page:     entity.Page,
		Limit:    entity.Limit,
		TotalRow: entity.TotalRows,
	}
}

func PendingOrderGroupToDto(entity *entity.PendingOrderGroup) *dto.PendingOrderGroupGetResponse {
	var orders []*dto.PendingOrderResponse
	for _, order := range entity.Order {
		orders = append(orders, PendingOrderToDto(order))
	}
	return &dto.PendingOrderGroupGetResponse{
		Id:           entity.Id,
		Orders:       orders,
		ShippingCost: entity.ShippingCost.String(),
		TotalPrice:   entity.TotalPrice.String(),
		UserAddress:  entity.UserAddress,
		CreatedAt:    entity.CreatedAt.Time.String(),
	}
}

func PendingOrderToDto(entity *entity.PendingOrder) *dto.PendingOrderResponse {
	var catalogs []*dto.PendingCatalogResponse
	for _, catalog := range entity.Catalogs {
		catalogs = append(catalogs, PendingCatalogToDto(catalog))
	}
	return &dto.PendingOrderResponse{
		Id:           entity.Id,
		PharmacyId:   entity.PharmacyId,
		PharmacyName: entity.PharmacyName,
		ShippingCost: entity.ShippingCost.String(),
		Catalogs:     catalogs,
	}
}

func PendingCatalogToDto(entity *entity.PendingCatalog) *dto.PendingCatalogResponse {
	return &dto.PendingCatalogResponse{
		Id:       entity.Id,
		Name:     entity.Name,
		Quantity: entity.Quantity,
		Price:    entity.Price.String(),
	}
}

func UserOrderUpdateRequestToEntity(dto *dto.UserOrderUpdateRequest) *entity.Order {
	return &entity.Order{
		Id: dto.Id,
	}
}

func OrderToUserOrderUpdateResponseDto(entity *entity.Order) *dto.UserOrderUpdateResponse {
	return &dto.UserOrderUpdateResponse{
		Id:       entity.Id,
		StatusId: entity.StatusId,
	}
}

func PharmacyOrderOptionRequestToEntity(dto *dto.PharmacyOrderOptionsRequest) *entity.PharmacyOrderOptions {
	var (
		page  int = constant.DEFAULT_PAGE
		limit int = constant.DEFAULT_LIMIT
	)

	if dto.Page != "" {
		page, _ = strconv.Atoi(dto.Page)
	}
	if dto.Limit != "" {
		limit, _ = strconv.Atoi(dto.Limit)
	}

	return &entity.PharmacyOrderOptions{
		Page:  page,
		Limit: limit,
	}
}

func PharmacyOrderOptionToDto(entity *entity.PharmacyOrderOptions) *dto.PharmacyOrderOptionsResponse {
	return &dto.PharmacyOrderOptionsResponse{
		Page:     entity.Page,
		Limit:    entity.Limit,
		TotalRow: entity.TotalRows,
	}
}

func OrderToPharmacyOrderGetResponseDto(entity *entity.Order) *dto.PharmacyOrderGetResponse {
	return &dto.PharmacyOrderGetResponse{
		Id:                entity.Id,
		Status:            entity.OrderStatus.Name,
		ProductCount:      len(entity.OrderItems),
		TotalPriceProduct: entity.TotalPriceProduct.String(),
	}
}

func OrderToPharmacyOrderGetDetailResponseDto(entity *entity.Order) *dto.PharmacyOrderGetDetailResponse {
	orderItems := make([]dto.OrderItem, len(entity.OrderItems))
	for i, o := range entity.OrderItems {
		orderItems[i] = *OrderItemToDto(o)
	}

	return &dto.PharmacyOrderGetDetailResponse{
		Id: entity.Id,
		User: dto.User{
			Id:   entity.User.Id,
			Name: entity.User.Name,
		},
		Address: dto.AddressOrderDetail{
			Id:   entity.Address.Id,
			Name: entity.Address.Name,
		},
		OrderStatus: dto.OrderStatus{
			Id:   entity.OrderStatus.Id,
			Name: entity.OrderStatus.Name,
		},
		LogisticPartner: dto.LogisticPartnerGetResponse{
			Id:   entity.LogisticPartner.Id,
			Name: entity.LogisticPartner.Name,
		},
		PaymentMethod: dto.PaymentMethod{
			Id:   entity.PaymentMethod.Id,
			Name: entity.PaymentMethod.Name,
		},
		TotalPriceShipping: entity.TotalPriceShipping.String(),
		TotalPriceProduct:  entity.TotalPriceProduct.String(),
		OrderItems:         orderItems,
	}
}

func PharmacyOrderUpdateRequestToEntity(dto *dto.PharmacyOrderUpdateRequest) *entity.Order {
	return &entity.Order{
		Id: dto.Id,
	}
}

func OrderToPharmacyOrderUpdateResponseDto(entity *entity.Order) *dto.PharmacyOrderUpdateResponse {
	return &dto.PharmacyOrderUpdateResponse{
		Id:       entity.Id,
		StatusId: entity.StatusId,
	}
}

func OrderOptionsRequestToEntity(dto *dto.OrderOptionsRequest) *entity.OrderOptions {
	var (
		pharmacyId *int64
		statusId   *int64
		page       int = constant.DEFAULT_PAGE
		limit      int = constant.DEFAULT_LIMIT
	)

	if dto.PharmacyId != "" {
		value, _ := strconv.ParseInt(dto.PharmacyId, 10, 64)
		pharmacyId = &value
	}
	if dto.StatusId != "" {
		value, _ := strconv.ParseInt(dto.StatusId, 10, 64)
		statusId = &value
	}
	if dto.Page != "" {
		page, _ = strconv.Atoi(dto.Page)
	}
	if dto.Limit != "" {
		limit, _ = strconv.Atoi(dto.Limit)
	}

	return &entity.OrderOptions{
		PharmacyId: pharmacyId,
		StatusId:   statusId,
		Page:       page,
		Limit:      limit,
	}
}

func OrderOptionsToDto(entity *entity.OrderOptions) *dto.OrderOptionsResponse {
	return &dto.OrderOptionsResponse{
		PharmacyId: entity.PharmacyId,
		StatusId:   entity.StatusId,
		Page:       entity.Page,
		Limit:      entity.Limit,
		TotalRow:   entity.TotalRows,
	}
}

func OrderToGetResponseDto(entity *entity.Order) *dto.OrderGetResponse {
	return &dto.OrderGetResponse{
		Id: entity.Id,
		Pharmacy: dto.PharmacyOrderDetail{
			Id:   entity.Pharmacy.Id,
			Name: entity.Pharmacy.Name,
		},
		Status: dto.OrderStatus{
			Id:   entity.OrderStatus.Id,
			Name: entity.OrderStatus.Name,
		},
	}
}

func OrderToGetDetailResponseDto(entity *entity.Order) *dto.OrderGetDetailResponse {
	orderItems := make([]dto.OrderItem, len(entity.OrderItems))
	for i, o := range entity.OrderItems {
		orderItems[i] = *OrderItemToDto(o)
	}

	return &dto.OrderGetDetailResponse{
		Id: entity.Id,
		User: dto.User{
			Id:   entity.User.Id,
			Name: entity.User.Name,
		},
		Address: dto.AddressOrderDetail{
			Id:   entity.Address.Id,
			Name: entity.Address.Name,
		},
		Pharmacy: dto.PharmacyOrderDetail{
			Id:   entity.Pharmacy.Id,
			Name: entity.Pharmacy.Name,
		},
		OrderStatus: dto.OrderStatus{
			Id:   entity.OrderStatus.Id,
			Name: entity.OrderStatus.Name,
		},
		LogisticPartner: dto.LogisticPartnerGetResponse{
			Id:   entity.LogisticPartner.Id,
			Name: entity.LogisticPartner.Name,
		},
		PaymentMethod: dto.PaymentMethod{
			Id:   entity.PaymentMethod.Id,
			Name: entity.PaymentMethod.Name,
		},
		TotalPriceShipping: entity.TotalPriceShipping.String(),
		TotalPriceProduct:  entity.TotalPriceProduct.String(),
		OrderItems:         orderItems,
	}
}

func OrderItemToDto(entity *entity.OrderItem) *dto.OrderItem {
	return &dto.OrderItem{
		Id:           entity.Id,
		Quantity:     entity.Quantity,
		CatalogId:    entity.Catalog.Id,
		Price:        entity.Catalog.Price.String(),
		ProductId:    entity.Catalog.ProductId,
		ProductName:  entity.Catalog.Name,
		ProductImage: entity.Catalog.Image,
	}
}

func UserOrderOptionRequestToEntity(dto *dto.UserOrderOptionsRequest) *entity.UserOrderOptions {
	var (
		page  = constant.DEFAULT_PAGE
		limit = constant.DEFAULT_LIMIT
	)

	if dto.Page != "" {
		page, _ = strconv.Atoi(dto.Page)
	}
	if dto.Limit != "" {
		limit, _ = strconv.Atoi(dto.Limit)
	}
	dto.FilterValue = strings.ReplaceAll(dto.FilterValue, "_", " ")

	return &entity.UserOrderOptions{
		FilterBy:    "status",
		FilterValue: dto.FilterValue,
		Page:        page,
		Limit:       limit,
	}
}

func UserOrderOptionToDto(entity *entity.UserOrderOptions) *dto.UserOrderOptionsResponse {
	return &dto.UserOrderOptionsResponse{
		Filter: dto.FilterOption{
			Column: entity.FilterBy,
			Value:  entity.FilterValue,
		},
		Page:     entity.Page,
		Limit:    entity.Limit,
		TotalRow: entity.TotalRows,
	}
}

func OrderToUserOrderGetResponseDto(entity *entity.Order) *dto.UserOrderGetResponse {
	var items []*dto.OrderItem
	for _, o := range entity.OrderItems {
		items = append(items, &dto.OrderItem{
			Id:           o.Catalog.Id,
			Quantity:     o.Quantity,
			Price:        o.Price.String(),
			ProductName:  o.Catalog.Name,
			ProductImage: o.Catalog.Image,
		})
	}

	return &dto.UserOrderGetResponse{
		Id:                 entity.Id,
		Status:             entity.OrderStatus.Name,
		PharmacyName:       entity.PharmacyName,
		Address:            entity.AddressName,
		CreatedAt:          entity.CreatedAt.Time.String(),
		TotalPriceShipping: entity.TotalPriceShipping.String(),
		TotalPriceProduct:  entity.TotalPriceProduct.String(),
		OrderItems:         items,
	}
}
