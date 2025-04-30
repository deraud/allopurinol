package handler

import (
	"healthcare-allopurinol/service"

	"github.com/gin-gonic/gin"
)

type UserHandlerItf interface {
	GetCatalogs(ctx *gin.Context)
	GetCatalog(ctx *gin.Context)
	PostCartItemIncrement(ctx *gin.Context)
	PostCartItemDecrement(ctx *gin.Context)
	PostCartItemSet(ctx *gin.Context)
	PostCartItemRemove(ctx *gin.Context)
	GetCartItems(ctx *gin.Context)
	GetCheckoutCatalog(ctx *gin.Context)
	PostShippingCost(ctx *gin.Context)
	PostOrderGroup(ctx *gin.Context)
	PutOrderGroup(ctx *gin.Context)
	GetPendingOrders(ctx *gin.Context)
	PatchUserOrderCanceled(ctx *gin.Context)
	PatchUserOrderConfirmed(ctx *gin.Context)
	GetUserAddressesHandler(ctx *gin.Context)
	GetUserProfileHandler(ctx *gin.Context)
	PatchUserProfileHandler(ctx *gin.Context)
	PatchUserRemovePictureHandler(ctx *gin.Context)
	DeleteUserAddressHandler(ctx *gin.Context)
	PostAddressHandler(ctx *gin.Context)
	PatchAddressHandler(ctx *gin.Context)
	PatchUserActivateAddressHandler(ctx *gin.Context)
	GetUserOrders(ctx *gin.Context)
	GetMostBoughtCatalogs(ctx *gin.Context)
}

type UserHandlerImpl struct {
	service         service.UserServiceItf
	catalogService  service.CatalogServiceItf
	cartService     service.CartServiceItf
	orderService    service.OrderServiceItf
	checkoutService service.CheckoutServiceItf
}

func NewUserHandler(service service.UserServiceItf, catalogService service.CatalogServiceItf, cartService service.CartServiceItf, orderService service.OrderServiceItf, checkoutService service.CheckoutServiceItf) UserHandlerItf {
	return &UserHandlerImpl{
		service:         service,
		catalogService:  catalogService,
		cartService:     cartService,
		orderService:    orderService,
		checkoutService: checkoutService,
	}
}
