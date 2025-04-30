package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *LocationHandlerImpl) GetProvincesHandler(ctx *gin.Context) {
	provinces, err := h.locationService.GetProvinces(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.Province, len(provinces))
	for i, p := range provinces {
		res[i] = mapper.ProvinceToDto(p)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: res,
	})
}
