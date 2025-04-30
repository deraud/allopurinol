package handler

import (
	"healthcare-allopurinol/service"

	"github.com/gin-gonic/gin"
)

type PharmacistHandlerItf interface {
	PutPharmacy(ctx *gin.Context)
	GetPharmacy(ctx *gin.Context)
	PostCatalog(ctx *gin.Context)
	PatchCatalog(ctx *gin.Context)
	DeleteCatalog(ctx *gin.Context)
	GetCatalogs(ctx *gin.Context)
	GetCatalog(ctx *gin.Context)
	GetPharmacyOrders(ctx *gin.Context)
	GetPharmacyOrder(ctx *gin.Context)
	PatchPharmacyOrderShipped(ctx *gin.Context)
	PatchPharmacyOrderCanceled(ctx *gin.Context)
}

type PharmacistHandlerImpl struct {
	pharmacyService service.PharmacyServiceItf
	catalogService  service.CatalogServiceItf
	orderService    service.OrderServiceItf
}

func NewPharmacistHandler(pharmacyService service.PharmacyServiceItf, catalogService service.CatalogServiceItf, orderService service.OrderServiceItf) PharmacistHandlerItf {
	return &PharmacistHandlerImpl{
		pharmacyService: pharmacyService,
		catalogService:  catalogService,
		orderService:    orderService,
	}
}
