package mapper

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"strconv"
)

func PharmacyOptionsDtoToEntity(dto *dto.PharmacyOptionsRequest) *entity.PharmacyOptions {
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
	if dto.SearchBy == "" {
		dto.SearchBy = constant.PHARMACY_DEFAULT_SEARCH_BY
	}

	return &entity.PharmacyOptions{
		SearchBy:    dto.SearchBy,
		SearchValue: dto.SearchValue,
		Page:        page,
		Limit:       limit,
	}
}

func PharmacyOptionsEntityToDto(entity *entity.PharmacyOptions) *dto.PharmacyOptionsResponse {
	return &dto.PharmacyOptionsResponse{
		Search:   dto.SearchOptions{Column: entity.SearchBy, Value: entity.SearchValue},
		Page:     entity.Page,
		Limit:    entity.Limit,
		TotalRow: entity.TotalRows,
	}
}

func PharmacyToGetResponseDto(entity *entity.Pharmacy) *dto.PharmacyGetResponse {
	return &dto.PharmacyGetResponse{
		Id:       entity.Id,
		Partner:  PartnerToDto(entity.Partner),
		Logo:     entity.Logo,
		Name:     entity.Name,
		Address:  entity.Address.Name,
		IsActive: entity.IsActive,
	}
}

func PharmacyCreateRequestToEntity(dto *dto.PharmacyCreateRequest) *entity.Pharmacy {
	return &entity.Pharmacy{
		PartnerId:          dto.PartnerId,
		Address:            AddressCreateRequestToEntity(&dto.Address),
		Name:               dto.Name,
		Logo:               dto.Logo,
		PharmacistIds:      dto.Pharmacists,
		LogisticPartnerIds: dto.LogisticPartners,
		IsActive:           *dto.IsActive,
	}
}

func PharmacyUpdateRequestToEntity(dto *dto.PharmacyUpdateRequest) *entity.Pharmacy {
	return &entity.Pharmacy{
		Id:                 dto.Id,
		Address:            AddressUpdateRequestToEntity(&dto.Address),
		Name:               dto.Name,
		Logo:               dto.Logo,
		PharmacistIds:      dto.Pharmacists,
		LogisticPartnerIds: dto.LogisticPartners,
		IsActive:           *dto.IsActive,
	}
}

func PharmacyToCreateResponseDto(entity *entity.Pharmacy) *dto.PharmacyCreateResponse {
	return &dto.PharmacyCreateResponse{
		Id:                 entity.Id,
		PartnerId:          entity.PartnerId,
		AddressId:          entity.Address.Id,
		Name:               entity.Name,
		Logo:               entity.Logo,
		PharmacistIds:      entity.PharmacistIds,
		LogisticPartnerIds: entity.LogisticPartnerIds,
		IsActive:           entity.IsActive,
	}
}

func PharmacyToUpdateResponseDto(entity *entity.Pharmacy) *dto.PharmacyUpdateResponse {
	return &dto.PharmacyUpdateResponse{
		Id:                 entity.Id,
		PartnerId:          entity.PartnerId,
		AddressId:          entity.Address.Id,
		Name:               entity.Name,
		Logo:               entity.Logo,
		PharmacistIds:      entity.PharmacistIds,
		LogisticPartnerIds: entity.LogisticPartnerIds,
		IsActive:           entity.IsActive,
	}
}

func PharmacyToGetDetailResponseDto(pharmacy *entity.Pharmacy) *dto.PharmacyGetDetailResponse {
	pharmacistsDto := make([]*dto.PharmacistGetResponse, len(pharmacy.Pharmacists))
	for i, p := range pharmacy.Pharmacists {
		pharmacistsDto[i] = PharmacistToDto(p, "")
	}

	logisticPartnersDto := make([]*dto.LogisticPartnerGetResponse, len(pharmacy.LogisticPartners))
	for i, lp := range pharmacy.LogisticPartners {
		logisticPartnersDto[i] = LogisticPartnerToGetResponseDto(lp)
	}

	return &dto.PharmacyGetDetailResponse{
		Id:               pharmacy.Id,
		Partner:          PartnerToDto(pharmacy.Partner),
		Address:          AddressToDto(pharmacy.Address),
		Name:             pharmacy.Name,
		Logo:             pharmacy.Logo,
		Pharmacists:      pharmacistsDto,
		LogisticPartners: logisticPartnersDto,
		IsActive:         pharmacy.IsActive,
	}
}

func PharmacistPharmacyUpdateRequestToEntity(dto *dto.PharmacistPharmacyUpdateRequest) *entity.Pharmacy {
	return &entity.Pharmacy{
		Address:            AddressUpdateRequestToEntity(&dto.Address),
		Name:               dto.Name,
		Logo:               dto.Logo,
		LogisticPartnerIds: dto.LogisticPartners,
		IsActive:           dto.IsActive,
	}
}

func PharmacyToCatalogGetDetailResponseDto(pharmacy *entity.Pharmacy) *dto.PharmacyCatalogGetDetailResponse {
	return &dto.PharmacyCatalogGetDetailResponse{
		Id:      pharmacy.Id,
		Address: AddressToDto(pharmacy.Address),
		Name:    pharmacy.Name,
	}
}

func PharmacyToCheckoutDetailDto(entity *entity.Pharmacy) *dto.PharmacyCheckoutDetail {
	logisticPartnerDto := make([]*dto.LogisticPartnerGetResponse, len(entity.LogisticPartners))
	for i, lp := range entity.LogisticPartners {
		logisticPartnerDto[i] = LogisticPartnerToGetResponseDto(lp)
	}

	return &dto.PharmacyCheckoutDetail{
		Id:               entity.Id,
		Name:             entity.Name,
		Address:          AddressToDto(entity.Address),
		LogisticPartners: logisticPartnerDto,
	}
}
