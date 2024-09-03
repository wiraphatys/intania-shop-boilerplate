package main

import (
	"github.com/wiraphatys/intania-shop-boilerplate/cmd/servers"
	"github.com/wiraphatys/intania-shop-boilerplate/internal/products"
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/config"
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/database"
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/logger"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewGormDatabase(cfg)

	logger := logger.NewLogger(cfg)

	// product
	productRepo := products.NewProductRepository(db)
	productSvc := products.NewProductService(cfg, logger.Named("ProductSvc"), productRepo)
	productGrpcHandler := products.NewProductGrpcHandler(productSvc)

	grpcServer := servers.NewGrpcServer(cfg, logger, db)
	grpcServer.Start(
		productGrpcHandler,
	)
}
