package server

import (
	"context"
	"errors"
	"healthcare-allopurinol/middleware"
	"healthcare-allopurinol/utility"
	"healthcare-allopurinol/utility/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	// connecting to database
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal()
		}
	}()

	ctx := context.Background()
	rdb := ConnectRedis(ctx)
	cld := ConnectCloudinary(ctx)

	err = utility.InitValidator()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.SetLogger(logger.NewLogrusLogger())

	serverHandler := NewServerHandler(db, ctx, rdb, cld)

	ginRouter := gin.New()
	ginRouter.Use(gin.Recovery())
	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))
	ginRouter.Use(middleware.LoggerMiddleware)
	ginRouter.Use(middleware.PrometheusMiddleware)
	ginRouter.Use(middleware.ErrorMiddleware)
	SetupRouter(ginRouter, serverHandler)

	timeout, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	// starting server
	s := &http.Server{
		Addr:         ":" + os.Getenv("APP_PORT"),
		Handler:      ginRouter,
		ReadTimeout:  time.Duration(timeout) * time.Second,
		WriteTimeout: time.Duration(timeout) * time.Second,
	}

	gracefulShutdown(s)
}

func gracefulShutdown(s *http.Server) {
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server exited gracefully")
}
