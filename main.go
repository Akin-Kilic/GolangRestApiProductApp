package main

import (
	"context"
	"go-rest-api/common/app"
	"go-rest-api/common/postgresql"
	"go-rest-api/controller"
	"go-rest-api/persistence"
	"go-rest-api/service"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	productRepository := persistence.NewProductRepository(dbPool)

	productService := service.NewProductService(productRepository)

	productController := controller.NewProductcontroller(productService)

	productController.RegisterRouters(e)

	e.Start("localhost:8080")
}
