package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *LocationHandlerImpl) GetCitiesHandler(ctx *gin.Context) {
	var params dto.LocationCityRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	cities, err := h.locationService.GetCities(ctx, mapper.CityRequestToProvinceId(&params))
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.City, len(cities))
	for i, c := range cities {
		res[i] = mapper.CityToDto(c)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: res,
	})
}
