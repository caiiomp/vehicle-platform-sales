package main

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/caiiomp/vehicle-platform-sales/src/core/useCases/sale"
	"github.com/caiiomp/vehicle-platform-sales/src/core/useCases/vehicle"
	_ "github.com/caiiomp/vehicle-platform-sales/src/docs"
	"github.com/caiiomp/vehicle-platform-sales/src/presentation"
	"github.com/caiiomp/vehicle-platform-sales/src/presentation/saleApi"
	"github.com/caiiomp/vehicle-platform-sales/src/presentation/vehicleApi"
	"github.com/caiiomp/vehicle-platform-sales/src/repositories/mongodb/saleRepository"
	"github.com/caiiomp/vehicle-platform-sales/src/repositories/mongodb/vehicleRepository"
)

func main() {
	var (
		mongoURI      = os.Getenv("MONGO_URI")
		mongoDatabase = os.Getenv("MONGO_DATABASE")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("could not initialize mongodb client: %v", err)
	}

	if err = mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Collections
	vehiclesCollection := mongoClient.Database(mongoDatabase).Collection("vehicles")
	salesCollection := mongoClient.Database(mongoDatabase).Collection("sales")

	// Repositories
	vehicleRepository := vehicleRepository.NewVehicleRepository(vehiclesCollection)
	saleRepository := saleRepository.NewSaleRepository(salesCollection)

	// Services
	vehicleService := vehicle.NewVehicleService(vehicleRepository, saleRepository)
	saleService := sale.NewSaleService(saleRepository)

	app := presentation.SetupServer()

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	vehicleApi.RegisterVehicleRoutes(app, vehicleService)
	saleApi.RegisterSaleRoutes(app, saleService)

	if err = app.Run(":4002"); err != nil {
		log.Fatalf("coult not initialize http server: %v", err)
	}
}
