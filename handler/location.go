package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"healthcare-allopurinol/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LocationHandlerItf interface {
	PostCoordinateHandler(ctx *gin.Context)
	PostAddressHandler(ctx *gin.Context)
	GetProvincesHandler(ctx *gin.Context)
	GetCitiesHandler(ctx *gin.Context)
	GetDistrictsHandler(ctx *gin.Context)
	GetSubdistrictsHandler(ctx *gin.Context)
}

type LocationHandlerImpl struct {
	locationService service.LocationServiceItf
}

func NewLocationHandler(locationService service.LocationServiceItf) LocationHandlerItf {
	return &LocationHandlerImpl{
		locationService: locationService,
	}
}

func (h *LocationHandlerImpl) PostCoordinateHandler(ctx *gin.Context) {
	var address dto.CoordinateRequest
	err := ctx.ShouldBind(&address)
	if err != nil {
		ctx.Error(err)
		return
	}

	coordinate, err := h.locationService.PostCoordinate(ctx, mapper.CoordinateRequestDtoToEntity(&address))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.CoordinateEntityToDtoResponse(coordinate),
	})
}

func (h *LocationHandlerImpl) PostAddressHandler(ctx *gin.Context) {
	var coordinate dto.AddressRequest
	err := ctx.ShouldBind(&coordinate)
	if err != nil {
		ctx.Error(err)
		return
	}

	address, err := h.locationService.PostAddress(ctx, mapper.AddressRequestDtotoEntity(&coordinate))

	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Data: mapper.AddressEntityToDtoResponse(address),
	})
}
