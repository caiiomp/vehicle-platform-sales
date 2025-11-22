package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	vehicleplatformpayments "github.com/caiiomp/vehicle-platform-sales/src/adapter/vehiclePlatformPayments"
	vehiclePlatformPaymentsHttpClient "github.com/caiiomp/vehicle-platform-sales/src/adapter/vehiclePlatformPayments/http"
	"github.com/caiiomp/vehicle-platform-sales/src/core/useCases/sale"
	"github.com/caiiomp/vehicle-platform-sales/src/core/useCases/vehicle"
	_ "github.com/caiiomp/vehicle-platform-sales/src/docs"
	"github.com/caiiomp/vehicle-platform-sales/src/presentation"
	"github.com/caiiomp/vehicle-platform-sales/src/presentation/saleApi"
	"github.com/caiiomp/vehicle-platform-sales/src/presentation/vehicleApi"
	salerepository "github.com/caiiomp/vehicle-platform-sales/src/repositories/postgres/saleRepository"
	vehiclerepository "github.com/caiiomp/vehicle-platform-sales/src/repositories/postgres/vehicleRepository"
)

func main() {
	var (
		apiPort = os.Getenv("API_PORT")

		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")

		vehiclePlatformPaymentsHost = os.Getenv("VEHICLE_PLATFORM_PAYMENTS_HOST")
		vehiclePlatformSalesHost    = os.Getenv("VEHICLE_PLATFORM_SALES_HOST")
	)

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("error to connect database: %s", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("error to ping database: %s", err)
	}

	// HTTP Clients
	httpClient := &http.Client{
		Timeout: time.Second * 3,
	}
	vehiclePlatformPaymentsHttpClient := vehiclePlatformPaymentsHttpClient.NewVehiclePlatformSalesHttpClient(httpClient, vehiclePlatformPaymentsHost, vehiclePlatformSalesHost)

	// Adapters
	vehiclePlatformPaymentsAdapter := vehicleplatformpayments.NewVehiclePlatformPaymentsAdapter(vehiclePlatformPaymentsHttpClient)

	// Repositories
	vehicleRepository := vehiclerepository.NewVehicleRepository(db)
	saleRepository := salerepository.NewSaleRepository(db)

	// Services
	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository, vehiclePlatformPaymentsAdapter)
	saleService := sale.NewSaleService(saleRepository, timeGenerator)

	app := presentation.SetupServer()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	vehicleApi.RegisterVehicleRoutes(app, vehicleService)
	saleApi.RegisterSaleRoutes(app, saleService)

	if apiPort == "" {
		apiPort = "8080"
	}

	if err = app.Run(":" + apiPort); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}

func timeGenerator() time.Time {
	return time.Now().UTC()
}
