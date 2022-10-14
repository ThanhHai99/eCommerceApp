package main

import (
	"eCommerce/model"
	routers "eCommerce/route"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
)

import "eCommerce/config"

func main() {
	app := gin.New()

	appConfig := config.AppConfig{}
	_ = env.Parse(&appConfig)

	migrateDb()

	// Setup route
	router := app.Group("/api")
	routers.SetupRouter(router)
	// Set domain
	domain := fmt.Sprintf(":%s", appConfig.AppPort)
	// run app
	_ = app.Run(domain)
}

func migrateDb() {
	db, _ := model.Connect()
	defer db.Close()
	_, _ = db.DB().Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	_, _ = db.DB().Exec("CREATE SEQUENCE IF NOT EXISTS file_id")
	db.AutoMigrate(&model.Category{}, &model.CategoryLog{}, &model.Invoice{}, &model.Item{}, &model.ItemLog{}, &model.Order{}, &model.PriceLog{}, &model.Product{}, &model.Role{}, &model.Sale{}, &model.SaleItem{}, &model.SaleLog{}, &model.User{}, &model.Warehouse{}, &model.WarehouseLog{})
}
