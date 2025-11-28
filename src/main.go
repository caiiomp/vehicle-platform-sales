package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
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
		ctx = context.Background()

		environment = os.Getenv("ENVIRONMENT")
		apiPort     = os.Getenv("API_PORT")

		instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
		host                   = os.Getenv("DB_HOST")
		port                   = os.Getenv("DB_PORT")
		user                   = os.Getenv("DB_USER")
		password               = os.Getenv("DB_PASSWORD")
		dbname                 = os.Getenv("DB_NAME")

		vehiclePlatformPaymentsHost = os.Getenv("VEHICLE_PLATFORM_PAYMENTS_HOST")
		vehiclePlatformSalesHost    = os.Getenv("VEHICLE_PLATFORM_SALES_HOST")
	)

	db, err := getDb(ctx, environment, instanceConnectionName, host, port, user, password, dbname)
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

func getDb(ctx context.Context, environment, instanceConnectionName, host, port, user, password, dbname string) (*sql.DB, error) {
	var (
		db *sql.DB

		config *pgx.ConnConfig
		dialer *cloudsqlconn.Dialer
		opts   []cloudsqlconn.Option

		dataSourceName string
		err            error
	)

	switch environment {
	case "PROD":
		dataSourceName = fmt.Sprintf("user=%s password=%s database=%s", user, password, dbname)

		config, err = pgx.ParseConfig(dataSourceName)
		if err != nil {
			return nil, err
		}

		opts = append(opts, cloudsqlconn.WithLazyRefresh())

		dialer, err = cloudsqlconn.NewDialer(ctx, opts...)
		if err != nil {
			return nil, err
		}

		config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
			return dialer.Dial(ctx, instanceConnectionName)
		}

		dbUri := stdlib.RegisterConnConfig(config)
		db, err = sql.Open("pgx", dbUri)

	default:
		dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err = sql.Open("postgres", dataSourceName)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func timeGenerator() time.Time {
	return time.Now().UTC()
}
