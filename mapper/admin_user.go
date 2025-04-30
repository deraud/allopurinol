package mapper

import (
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/entity"
	"strconv"
)

func UserInfoToDto(userInfo *entity.UserInfo) *dto.UserInfoRequest {
	return &dto.UserInfoRequest{
		Id:    userInfo.Id,
		Email: userInfo.Email,
		Name:  userInfo.Name,
		Role:  userInfo.Role,
	}
}

func UserInfoOptionsToEntity(userInfoOptions *dto.UserInfoOptionsRequest) *entity.UserInfoOptions {
	var (
		page  int = constant.DEFAULT_PAGE
		limit int = constant.DEFAULT_LIMIT
	)

	if userInfoOptions.Page != "" {
		page, _ = strconv.Atoi(userInfoOptions.Page)
	}
	if userInfoOptions.Limit != "" {
		limit, _ = strconv.Atoi(userInfoOptions.Limit)
	}
	if userInfoOptions.SearchBy == "" {
		userInfoOptions.SearchBy = constant.USER_DEFAULT_SEARCH_BY
	}

	return &entity.UserInfoOptions{
		SearchBy:    userInfoOptions.SearchBy,
		SearchValue: userInfoOptions.SearchValue,
		Page:        page,
		Limit:       limit,
	}
}

func UserInfoOptionsToDto(userInfoOptions *entity.UserInfoOptions) *dto.UserInfoOptionsResponse {
	return &dto.UserInfoOptionsResponse{
		Search: dto.SearchOptions{
			Column: userInfoOptions.SearchBy,
			Value:  userInfoOptions.SearchValue,
		},
		Page:     userInfoOptions.Page,
		Limit:    userInfoOptions.Limit,
		TotalRow: userInfoOptions.TotalRows,
	}
}
