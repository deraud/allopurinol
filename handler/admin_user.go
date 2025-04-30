package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandlerImpl) GetUsersHandler(ctx *gin.Context) {
	var params dto.UserInfoOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	users, options, err := h.adminService.GetUsersService(ctx, mapper.UserInfoOptionsToEntity(&params))
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.UserInfoRequest, len(users))
	for i, u := range users {
		res[i] = mapper.UserInfoToDto(u)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: dto.ListData{Entries: res, PageInfo: mapper.UserInfoOptionsToDto(options)},
	})
}
