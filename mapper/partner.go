package mapper

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"strconv"
)

func PartnerToDto(entity *entity.Partner) *dto.Partner {
	return &dto.Partner{
		Id:   entity.Id,
		Name: entity.Name,
	}
}

func PartnerCreateRequestDtoToEntityPartner(dto *dto.PartnerCreateRequest) *entity.Partner {
	return &entity.Partner{
		Name:                 dto.Name,
		YearFounded:          dto.YearFounded,
		ActiveDays:           dto.ActiveDays,
		OperationalHourStart: dto.OperationalHourStart,
		OperationalHourEnd:   dto.OperationalHourEnd,
		IsActive:             *dto.IsActive,
	}
}

func PartnerEntityToDtoPartnerCreateResponse(p *entity.Partner) *dto.PartnerCreateResponse {
	return &dto.PartnerCreateResponse{
		Id:                   p.Id,
		Name:                 p.Name,
		YearFounded:          p.YearFounded,
		ActiveDays:           p.ActiveDays,
		OperationalHourStart: p.OperationalHourStart,
		OperationalHourEnd:   p.OperationalHourEnd,
		IsActive:             p.IsActive,
	}
}

func PartnerUpdateRequestDtoToEntity(dto *dto.PartnerUpdateRequest) *entity.Partner {
	return &entity.Partner{
		Id:                   dto.Id,
		ActiveDays:           dto.ActiveDays,
		OperationalHourStart: dto.OperationalHourStart,
		OperationalHourEnd:   dto.OperationalHourEnd,
		IsActive:             *dto.IsActive,
	}
}

func PartnerEntityToDtoPartnerUpdateResponse(entity *entity.Partner) *dto.PartnerUpdateResponse {
	return &dto.PartnerUpdateResponse{
		Id:                   entity.Id,
		Name:                 entity.Name,
		YearFounded:          entity.YearFounded,
		ActiveDays:           entity.ActiveDays,
		OperationalHourStart: entity.OperationalHourStart,
		OperationalHourEnd:   entity.OperationalHourEnd,
		IsActive:             entity.IsActive,
	}
}

func PartnerOptionsDtoToEntity(dto *dto.PartnerOptionsRequest) *entity.PartnerOptions {
	var (
		searchBy  string = constant.PARTNER_DEFAULT_SEARCH_BY
		sortBy    string = constant.PARTNER_DEFAULT_SORT_BY
		sortOrder string = constant.PARTNER_DEFAULT_SORT_ORDER
		page      int    = constant.DEFAULT_PAGE
		limit     int    = constant.DEFAULT_LIMIT
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

	return &entity.PartnerOptions{
		SearchBy:    searchBy,
		SearchValue: dto.SearchValue,
		SortBy:      sortBy,
		SortOrder:   sortOrder,
		Page:        page,
		Limit:       limit,
	}
}

func PartnerOptionsEntityToDto(entity *entity.PartnerOptions) *dto.PartnerOptionsResponse {
	return &dto.PartnerOptionsResponse{
		Search:   dto.SearchOptions{Column: entity.SearchBy, Value: entity.SearchValue},
		Sort:     dto.SortOption{Column: entity.SortBy, Order: entity.SortOrder},
		Page:     entity.Page,
		Limit:    entity.Limit,
		TotalRow: entity.TotalRows,
	}
}

func PartnerEntityToDto(entity *entity.Partner) *dto.PartnerGetResponse {
	return &dto.PartnerGetResponse{
		Id:                   entity.Id,
		Name:                 entity.Name,
		ActiveDays:           entity.ActiveDays,
		YearFounded:          entity.YearFounded,
		OperationalHourStart: entity.OperationalHourStart,
		OperationalHourEnd:   entity.OperationalHourEnd,
		IsActive:             entity.IsActive,
	}
}
