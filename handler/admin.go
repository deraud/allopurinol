package handler

import (
	"healthcare-allopurinol/service"

	"github.com/gin-gonic/gin"
)

type AdminHandlerItf interface {
	GetUsersHandler(ctx *gin.Context)
	PostPharmacistHandler(ctx *gin.Context)
	PatchPharmacistHandler(ctx *gin.Context)
	DeletePharmacistHandler(ctx *gin.Context)
	GetPharmacistsHandler(ctx *gin.Context)
	GetPharmacistHandler(ctx *gin.Context)
	GetPharmaciesHandler(ctx *gin.Context)
	PostPharmacyHandler(ctx *gin.Context)
	PostProductHandler(ctx *gin.Context)
	GetProductsHandler(ctx *gin.Context)
	DeleteProductHandler(ctx *gin.Context)
	GetProductHandler(ctx *gin.Context)
	PutProductHandler(ctx *gin.Context)
	GetProductClassificationsHandler(ctx *gin.Context)
	GetProductFormsHandler(ctx *gin.Context)
	GetProductManufacturersHandler(ctx *gin.Context)
	PutPharmacyHandler(ctx *gin.Context)
	GetPharmacyHandler(ctx *gin.Context)
	DeletePharmacyHandler(ctx *gin.Context)
	PostPartnerHandler(ctx *gin.Context)
	PatchPartnerHandler(ctx *gin.Context)
	DeletePartnerHandler(ctx *gin.Context)
	GetPartnersHandler(ctx *gin.Context)
	GetPartnerHandler(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
	GetOrder(ctx *gin.Context)
	PostProductCategoryHandler(ctx *gin.Context)
	PatchProductCategoryHandler(ctx *gin.Context)
	DeleteProductCategoryHandler(ctx *gin.Context)
	GetProductCategoriesHandler(ctx *gin.Context)
	GetDashboardCountHandler(ctx *gin.Context)
	GetSalesReportHandler(ctx *gin.Context)
}

type AdminHandlerImpl struct {
	adminService      service.AdminServiceItf
	pharmacistService service.PharmacistServiceItf
	pharmacyService   service.PharmacyServiceItf
	productService    service.ProductServiceItf
	partnerService    service.PartnerServiceItf
	orderService      service.OrderServiceItf
}

func NewAdminHandler(adminService service.AdminServiceItf, pharmacistService service.PharmacistServiceItf, pharmacyService service.PharmacyServiceItf, productService service.ProductServiceItf, partnerService service.PartnerServiceItf, orderService service.OrderServiceItf) AdminHandlerItf {
	return &AdminHandlerImpl{
		adminService:      adminService,
		pharmacistService: pharmacistService,
		pharmacyService:   pharmacyService,
		productService:    productService,
		partnerService:    partnerService,
		orderService:      orderService,
	}
}
