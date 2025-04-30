package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"healthcare-allopurinol/middleware"
)

func SetupRouter(ginRouter *gin.Engine, handler *ServerHandler) {
	ginRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))
	setupAdminRouting(ginRouter, handler)
	setupGeneralRouting(ginRouter, handler)
	setupAuthRouting(ginRouter, handler)
	setupPharmacistRouting(ginRouter, handler)
}

func setupAdminRouting(ginRouter *gin.Engine, handler *ServerHandler) {
	adminRouter := ginRouter.Group("/admin", middleware.AuthenticationMiddleware, middleware.AdminMiddleware)
	{
		adminRouter.GET("/users", handler.AdminHandler.GetUsersHandler)
		adminRouter.POST("/pharmacists", handler.AdminHandler.PostPharmacistHandler)
		adminRouter.PATCH("/pharmacists", handler.AdminHandler.PatchPharmacistHandler)
		adminRouter.DELETE("/pharmacists/:id", handler.AdminHandler.DeletePharmacistHandler)
		adminRouter.GET("/pharmacists", handler.AdminHandler.GetPharmacistsHandler)
		adminRouter.GET("/pharmacists/:id", handler.AdminHandler.GetPharmacistHandler)
		adminRouter.GET("/pharmacies", handler.AdminHandler.GetPharmaciesHandler)
		adminRouter.POST("/pharmacies", handler.AdminHandler.PostPharmacyHandler)
		adminRouter.POST("/product-categories", handler.AdminHandler.PostProductCategoryHandler)
		adminRouter.PATCH("/product-categories", handler.AdminHandler.PatchProductCategoryHandler)
		adminRouter.DELETE("/product-categories/:id", handler.AdminHandler.DeleteProductCategoryHandler)
		adminRouter.POST("/products", handler.AdminHandler.PostProductHandler)
		adminRouter.GET("/products", handler.AdminHandler.GetProductsHandler)
		adminRouter.DELETE("/products/:id", handler.AdminHandler.DeleteProductHandler)
		adminRouter.GET("/products/:id", handler.AdminHandler.GetProductHandler)
		adminRouter.PUT("/products", handler.AdminHandler.PutProductHandler)
		adminRouter.GET("/product-categories", handler.AdminHandler.GetProductCategoriesHandler)
		adminRouter.GET("/product-classifications", handler.AdminHandler.GetProductClassificationsHandler)
		adminRouter.GET("/product-forms", handler.AdminHandler.GetProductFormsHandler)
		adminRouter.GET("/product-manufacturers", handler.AdminHandler.GetProductManufacturersHandler)
		adminRouter.PUT("/pharmacies", handler.AdminHandler.PutPharmacyHandler)
		adminRouter.GET("/pharmacies/:id", handler.AdminHandler.GetPharmacyHandler)
		adminRouter.DELETE("/pharmacies/:id", handler.AdminHandler.DeletePharmacyHandler)
		adminRouter.POST("/partners", handler.AdminHandler.PostPartnerHandler)
		adminRouter.PATCH("/partners", handler.AdminHandler.PatchPartnerHandler)
		adminRouter.DELETE("/partners/:id", handler.AdminHandler.DeletePartnerHandler)
		adminRouter.GET("/partners", handler.AdminHandler.GetPartnersHandler)
		adminRouter.GET("/partners/:id", handler.AdminHandler.GetPartnerHandler)
		adminRouter.GET("/dashboard-count", handler.AdminHandler.GetDashboardCountHandler)
		adminRouter.GET("/report", handler.AdminHandler.GetSalesReportHandler)
		adminRouter.GET("/orders", handler.AdminHandler.GetOrders)
		adminRouter.GET("/orders/:id", handler.AdminHandler.GetOrder)
	}
}

func setupAuthRouting(ginRouter *gin.Engine, handler *ServerHandler) {
	authRouter := ginRouter.Group("/auth")
	{
		authRouter.POST("/login", handler.AuthHandler.LoginHandler)
		authRouter.POST("/register", handler.AuthHandler.RegisterHandler)
		authRouter.POST("/forgot-password", handler.AuthHandler.ForgotPasswordHandler)
		authRouter.POST("/reset-password", handler.AuthHandler.ResetPasswordHandler)
		verifyRouter := ginRouter.Group("/auth/verify", middleware.AuthenticationMiddleware)
		{
			verifyRouter.PATCH("", handler.AuthHandler.VerifyHandler)
		}
		verifyRouter.POST("/token", handler.AuthHandler.VerifyTokenHandler)
	}

	userRouter := ginRouter.Group("/user", middleware.AuthenticationMiddleware)
	{
		userRouter.GET("/catalogs", handler.UserHandler.GetCatalogs)
		userRouter.POST("/carts/increment", middleware.VerificationMiddleware, handler.UserHandler.PostCartItemIncrement)
		userRouter.POST("/carts/decrement", middleware.VerificationMiddleware, handler.UserHandler.PostCartItemDecrement)
		userRouter.POST("/carts/set", middleware.VerificationMiddleware, handler.UserHandler.PostCartItemSet)
		userRouter.POST("/carts/remove", middleware.VerificationMiddleware, handler.UserHandler.PostCartItemRemove)
		userRouter.GET("/carts", middleware.VerificationMiddleware, handler.UserHandler.GetCartItems)
		userRouter.POST("/checkout/catalogs", middleware.VerificationMiddleware, handler.UserHandler.GetCheckoutCatalog)
		userRouter.POST("/carts/shipping-cost", middleware.VerificationMiddleware, handler.UserHandler.PostShippingCost)
		userRouter.POST("/order-groups", middleware.VerificationMiddleware, handler.UserHandler.PostOrderGroup)
		userRouter.PUT("/order-groups", middleware.VerificationMiddleware, handler.UserHandler.PutOrderGroup)
		userRouter.GET("/orders/pending", handler.UserHandler.GetPendingOrders)
		userRouter.GET("/orders", handler.UserHandler.GetUserOrders)
		userRouter.PATCH("/orders/:id/cancel", handler.UserHandler.PatchUserOrderCanceled)
		userRouter.PATCH("/orders/:id/confirm", handler.UserHandler.PatchUserOrderConfirmed)
		userRouter.GET("/profile", handler.UserHandler.GetUserProfileHandler)
		userRouter.PATCH("/profile", handler.UserHandler.PatchUserProfileHandler)
		userRouter.PATCH("/profile/remove-picture", handler.UserHandler.PatchUserRemovePictureHandler)
		userRouter.GET("/addresses", handler.UserHandler.GetUserAddressesHandler)
		userRouter.DELETE("/addresses/:address_id", handler.UserHandler.DeleteUserAddressHandler)
		userRouter.POST("/addresses", handler.UserHandler.PostAddressHandler)
		userRouter.PATCH("/addresses", handler.UserHandler.PatchAddressHandler)
		userRouter.PATCH("/addresses/activate", handler.UserHandler.PatchUserActivateAddressHandler)
	}
	ginRouter.GET("user/catalogs/:id", handler.UserHandler.GetCatalog)
}

func setupGeneralRouting(ginRouter *gin.Engine, handler *ServerHandler) {
	generalRouter := ginRouter.Group("")
	{
		generalRouter.GET("/provinces", handler.LocationHandler.GetProvincesHandler)
		generalRouter.GET("/cities/:province_id", handler.LocationHandler.GetCitiesHandler)
		generalRouter.GET("/districts/:city_id", handler.LocationHandler.GetDistrictsHandler)
		generalRouter.GET("/subdistricts/:district_id", handler.LocationHandler.GetSubdistrictsHandler)
		generalRouter.GET("/logistic-partners", handler.LogisticPartnerHandler.GetLogisticPartners)
		generalRouter.POST("/images", handler.ImageHandler.PostHandler)
		generalRouter.POST("/user/coordinate", handler.LocationHandler.PostCoordinateHandler)
		generalRouter.POST("/user/address", handler.LocationHandler.PostAddressHandler)
		generalRouter.GET("/catalogs", handler.UserHandler.GetMostBoughtCatalogs)
		generalRouter.GET("/product-categories", handler.AdminHandler.GetProductCategoriesHandler)
		generalRouter.GET("/product-classifications", handler.AdminHandler.GetProductClassificationsHandler)
		generalRouter.GET("/product-forms", handler.AdminHandler.GetProductFormsHandler)
		generalRouter.GET("/product-manufacturers", handler.AdminHandler.GetProductManufacturersHandler)
	}
}

func setupPharmacistRouting(ginRouter *gin.Engine, handler *ServerHandler) {
	pharmacistRouter := ginRouter.Group("/pharmacist", middleware.AuthenticationMiddleware, middleware.PharmacistMiddleware)
	{
		pharmacistRouter.PUT("/pharmacy", handler.PharmacistHandler.PutPharmacy)
		pharmacistRouter.GET("/pharmacy", handler.PharmacistHandler.GetPharmacy)
		pharmacistRouter.POST("/catalogs", handler.PharmacistHandler.PostCatalog)
		pharmacistRouter.PATCH("/catalogs", handler.PharmacistHandler.PatchCatalog)
		pharmacistRouter.DELETE("/catalogs/:id", handler.PharmacistHandler.DeleteCatalog)
		pharmacistRouter.GET("/catalogs", handler.PharmacistHandler.GetCatalogs)
		pharmacistRouter.GET("/catalogs/:id", handler.PharmacistHandler.GetCatalog)
		pharmacistRouter.GET("/orders", handler.PharmacistHandler.GetPharmacyOrders)
		pharmacistRouter.GET("/orders/:id", handler.PharmacistHandler.GetPharmacyOrder)
		pharmacistRouter.PATCH("/orders/:id/send", handler.PharmacistHandler.PatchPharmacyOrderShipped)
		pharmacistRouter.PATCH("/orders/:id/cancel", handler.PharmacistHandler.PatchPharmacyOrderCanceled)
		pharmacistRouter.GET("/report", handler.AdminHandler.GetSalesReportHandler)
		pharmacistRouter.GET("/products", handler.AdminHandler.GetProductsHandler)
		pharmacistRouter.GET("/dashboard-count", handler.AdminHandler.GetDashboardCountHandler)
	}
}
