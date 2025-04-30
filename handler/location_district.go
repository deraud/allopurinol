package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *LocationHandlerImpl) GetDistrictsHandler(ctx *gin.Context) {
	var params dto.LocationDistrictRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	districts, err := h.locationService.GetDistricts(ctx, mapper.DistrictRequestToCityId(&params))
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.District, len(districts))
	for i, d := range districts {
		res[i] = mapper.DistrictToDto(d)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: res,
	})
}
