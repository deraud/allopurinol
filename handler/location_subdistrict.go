package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *LocationHandlerImpl) GetSubdistrictsHandler(ctx *gin.Context) {
	var params dto.LocationSubdistrictRequest
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.Error(err)
		return
	}

	subdistricts, err := h.locationService.GetSubdistricts(ctx, mapper.SubdistrictRequestToDistrictId(&params))
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.Subdistrict, len(subdistricts))
	for i, s := range subdistricts {
		res[i] = mapper.SubdistrictToDto(s)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: res,
	})
}
