package main

import (
	"log"
	"net/http"

	_ "shop-api/docs" // сгенерируется swag init

	"shop-api/internal/config"
	"shop-api/internal/db"
	"shop-api/internal/handlers"
	"shop-api/internal/repository"
	"shop-api/internal/router"
)

// @title           Shop API
// @version         1.0
// @description     REST API магазина бытовой техники
// @termsOfService  http://swagger.io/terms/
// @contact.name    Shop API Support
// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http
func main() {
	cfg := config.Load()

	conn, err := db.Connect(cfg.DSN)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer conn.Close()

	clientRepo := repository.NewClientRepository(conn)
	productRepo := repository.NewProductRepository(conn)
	supplierRepo := repository.NewSupplierRepository(conn)
	imageRepo := repository.NewImageRepository(conn)

	h := router.Handlers{
		Client:   handlers.NewClientHandler(clientRepo),
		Product:  handlers.NewProductHandler(productRepo),
		Supplier: handlers.NewSupplierHandler(supplierRepo),
		Image:    handlers.NewImageHandler(imageRepo),
	}

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router.New(h),
	}

	log.Printf("listening on :%s (swagger: http://localhost:%s/swagger/index.html)", cfg.Port, cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
