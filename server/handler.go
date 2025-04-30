package server

import (
	"context"
	"database/sql"
	"healthcare-allopurinol/bridge"
	"healthcare-allopurinol/constant"
	"healthcare-allopurinol/handler"
	"healthcare-allopurinol/repo"
	"healthcare-allopurinol/service"
	"os"

	"github.com/resend/resend-go/v2"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/redis/go-redis/v9"
)

type ServerHandler struct {
	AdminHandler           handler.AdminHandlerItf
	AuthHandler            handler.AuthHandlerItf
	LocationHandler        handler.LocationHandlerItf
	LogisticPartnerHandler handler.LogisticPartnerHandlerItf
	PharmacistHandler      handler.PharmacistHandlerItf
	UserHandler            handler.UserHandlerItf
	ImageHandler           handler.ImageHandlerItf
}

func NewServerHandler(db *sql.DB, ctx context.Context, rdb *redis.Client, cld *cloudinary.Cloudinary) *ServerHandler {
	transactor := repo.NewTransactor(db)
	redisRepo := repo.NewRedisRepo(rdb, ctx)
	cloudinaryBridge := bridge.NewCloudinaryBridge(cld, ctx)
	rajaOngkirBridge := bridge.NewROBridge(constant.RAJA_ONGKIR_HOST, os.Getenv("RAJA_ONGKIR_KEY"))
	mailBridge := bridge.NewMailBridge(resend.NewClient(os.Getenv("MAIL_API_KEY")))

	pharmacistRepo := repo.NewPharmacistRepo(db)
	userRepo := repo.NewUserRepo(db)

	partnerRepo := repo.NewPartnerRepo(db)
	addressRepo := repo.NewAddressRepo(db)
	pharmacyRepo := repo.NewPharmacyRepo(db)
	shippingRepo := repo.NewShippingRepo(db)
	orderRepo := repo.NewOrderRepo(db)

	rajaOngkirRepo := repo.NewRORepo(db)

	authRepo := repo.NewAuthRepo(db)
	authService := service.NewAuthService(authRepo, userRepo, transactor, redisRepo, mailBridge)
	authHandler := handler.NewAuthHandler(authService)

	pharmacistService := service.NewPharmacistService(transactor, authRepo, pharmacistRepo, pharmacyRepo, mailBridge)
	pharmacyService := service.NewPharmacyService(transactor, pharmacyRepo, pharmacistRepo, partnerRepo, addressRepo, shippingRepo)

	productRepo := repo.NewProductRepo(db)
	productService := service.NewProductService(transactor, productRepo)

	partnerService := service.NewPartnerService(partnerRepo)

	locationBridge := bridge.NewLocationBridge(os.Getenv("EMSIFA_HOST"))
	locationService := service.NewLocationService(locationBridge)
	locationHandler := handler.NewLocationHandler(locationService)

	logisticPartnerService := service.NewLogisticPartnerService(shippingRepo)
	logisticPartnerHandler := handler.NewLogisticPartnerHandler(logisticPartnerService)

	cartRepo := repo.NewCartRepo(db)
	cartService := service.NewCartService(transactor, userRepo, productRepo, cartRepo)

	catalogRepo := repo.NewCatalogRepo(db)
	catalogService := service.NewCatalogService(productRepo, catalogRepo, pharmacistRepo, pharmacyRepo, addressRepo, userRepo, cartRepo, shippingRepo)

	orderService := service.NewOrderService(transactor, redisRepo, orderRepo, userRepo, catalogRepo, cartRepo, pharmacistRepo)
	checkoutService := service.NewCheckoutService(transactor, rajaOngkirBridge, shippingRepo, rajaOngkirRepo, pharmacyRepo, addressRepo, userRepo, redisRepo)

	adminRepo := repo.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService, pharmacistService, pharmacyService, productService, partnerService, orderService)

	pharmacistHandler := handler.NewPharmacistHandler(pharmacyService, catalogService, orderService)

	userService := service.NewUserService(transactor, userRepo)
	userHandler := handler.NewUserHandler(userService, catalogService, cartService, orderService, checkoutService)

	imageService := service.NewImageService(cloudinaryBridge)
	imageHandler := handler.NewImageHandler(imageService)

	return &ServerHandler{
		AdminHandler:           adminHandler,
		AuthHandler:            authHandler,
		LocationHandler:        locationHandler,
		LogisticPartnerHandler: logisticPartnerHandler,
		PharmacistHandler:      pharmacistHandler,
		UserHandler:            userHandler,
		ImageHandler:           imageHandler,
	}
}
