package mapper

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"strconv"
)

func PharmacistCreateRequestToCredentialEntity(dto *dto.PharmacistCreateRequest) *entity.Credential {
	return &entity.Credential{
		Email:    dto.Email,
		Password: dto.Password,
		RoleId:   constant.ROLE_PHARMACIST,
	}
}

func PharmacistCreateRequestToPharmacistEntity(dto *dto.PharmacistCreateRequest) *entity.Pharmacist {
	return &entity.Pharmacist{
		Name:              dto.Name,
		SipaNumber:        dto.SipaNumber,
		PhoneNumber:       dto.PhoneNumber,
		YearsOfExperience: dto.YearsOfExperience,
	}
}

func PharmacistToPharmacistCreateResponseDto(p *entity.Pharmacist, c *entity.Credential) *dto.PharmacistCreateResponse {
	return &dto.PharmacistCreateResponse{
		Id:                p.Id,
		PharmacyId:        p.PharmacyId,
		Name:              p.Name,
		SipaNumber:        p.SipaNumber,
		PhoneNumber:       p.PhoneNumber,
		YearsOfExperience: p.YearsOfExperience,
		Email:             c.Email,
	}
}

func PharmacistUpdateRequestToEntity(dto *dto.PharmacistUpdateRequest) *entity.Pharmacist {
	return &entity.Pharmacist{
		Id:                dto.Id,
		PharmacyId:        dto.PharmacyId,
		PhoneNumber:       dto.PhoneNumber,
		YearsOfExperience: dto.YearsOfExperience,
	}
}

func PharmacistToPharmacistUpdateResponseDto(entity *entity.Pharmacist) *dto.PharmacistUpdateResponse {
	return &dto.PharmacistUpdateResponse{
		Id:                entity.Id,
		PharmacyId:        entity.PharmacyId,
		Name:              entity.Name,
		SipaNumber:        entity.SipaNumber,
		PhoneNumber:       entity.PhoneNumber,
		YearsOfExperience: entity.YearsOfExperience,
	}
}

func PharmacistOptionsToEntity(dto *dto.PharmacistOptionsRequest) *entity.PharmacistOptions {
	var (
		searchBy      string = constant.PHARMACIST_DEFAULT_SEARCH_BY
		sortBy        string = constant.PHARMACIST_DEFAULT_SORT_BY
		sortOrder     string = constant.PHARMACIST_DEFAULT_SORT_ORDER
		assigned      *bool
		yearsExpStart *int
		yearsExpEnd   *int
		page          int = constant.DEFAULT_PAGE
		limit         int = constant.DEFAULT_LIMIT
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
	if dto.Assigned != "" {
		value, _ := strconv.ParseBool(dto.Assigned)
		assigned = &value
	}
	if dto.YearsExpStart != "" {
		value, _ := strconv.Atoi(dto.YearsExpStart)
		yearsExpStart = &value
	}
	if dto.YearsExpEnd != "" {
		value, _ := strconv.Atoi(dto.YearsExpEnd)
		yearsExpEnd = &value
	}
	if dto.Page != "" {
		page, _ = strconv.Atoi(dto.Page)
	}
	if dto.Limit != "" {
		limit, _ = strconv.Atoi(dto.Limit)
	}

	return &entity.PharmacistOptions{
		SearchBy:      searchBy,
		SearchValue:   dto.SearchValue,
		SortBy:        sortBy,
		SortOrder:     sortOrder,
		Assigned:      assigned,
		YearsExpStart: yearsExpStart,
		YearsExpEnd:   yearsExpEnd,
		Page:          page,
		Limit:         limit,
	}
}

func PharmacistOptionsToDto(entity *entity.PharmacistOptions) *dto.PharmacistOptionsResponse {
	return &dto.PharmacistOptionsResponse{
		Search:        dto.SearchOptions{Column: entity.SearchBy, Value: entity.SearchValue},
		Sort:          dto.SortOption{Column: entity.SortBy, Order: entity.SortOrder},
		Assigned:      entity.Assigned,
		YearsExpStart: entity.YearsExpStart,
		YearsExpEnd:   entity.YearsExpEnd,
		Page:          entity.Page,
		Limit:         entity.Limit,
		TotalRow:      entity.TotalRows,
	}
}

func PharmacistToDto(entity *entity.Pharmacist, email string) *dto.PharmacistGetResponse {
	return &dto.PharmacistGetResponse{
		Id:                entity.Id,
		PharmacyId:        entity.PharmacyId,
		PharmacyName:      entity.PharmacyName,
		Name:              entity.Name,
		Email:             email,
		SipaNumber:        entity.SipaNumber,
		PhoneNumber:       entity.PhoneNumber,
		YearsOfExperience: entity.YearsOfExperience,
	}
}
