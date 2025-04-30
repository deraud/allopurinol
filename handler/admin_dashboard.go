package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AdminHandlerImpl) GetDashboardCountHandler(ctx *gin.Context) {
	var dashboard dto.DashboardCount

	count, err := h.adminService.GetDashboardCountService(ctx, mapper.DashboardCountDtoToEntity(&dashboard))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: mapper.DashboardCountEntityToDto(count),
	})
}
