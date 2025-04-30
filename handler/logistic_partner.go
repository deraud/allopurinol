package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"healthcare-allopurinol/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogisticPartnerHandlerItf interface {
	GetLogisticPartners(ctx *gin.Context)
}

type LogisticPartnerHandlerImpl struct {
	logisticPartnerService service.LogisticPartnerServiceItf
}

func NewLogisticPartnerHandler(logisticPartnerService service.LogisticPartnerServiceItf) LogisticPartnerHandlerItf {
	return &LogisticPartnerHandlerImpl{
		logisticPartnerService: logisticPartnerService,
	}
}

func (h *LogisticPartnerHandlerImpl) GetLogisticPartners(ctx *gin.Context) {
	logisticPartners, err := h.logisticPartnerService.GetLogisticPartnerService(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	res := make([]*dto.LogisticPartnerGetResponse, len(logisticPartners))
	for i, lp := range logisticPartners {
		res[i] = mapper.LogisticPartnerToGetResponseDto(lp)
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Data: res,
	})
}
