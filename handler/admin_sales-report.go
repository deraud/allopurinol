package handler

import (
	"healthcare-allopurinol/dto"
	"healthcare-allopurinol/mapper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *AdminHandlerImpl) GetSalesReportHandler(ctx *gin.Context) {
	var params dto.ReportOptionsRequest
	err := ctx.ShouldBindQuery(&params)
	if err != nil {
		ctx.Error(err)
		return
	}
	reports, err := s.adminService.GetSalesReportService(ctx, mapper.ReportOptionsDtoToEntity(&params))
	if err != nil {
		ctx.Error(err)
		return
	}
	res := make([]*dto.GetReportResponse, len(reports))
	for i, reports := range reports {
		res[i] = mapper.ReportEntityToDto(reports)
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Data: res,
	})
}
