package vehicleApi

import (
	"net/http"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/responses"
	"github.com/gin-gonic/gin"
)

type vehicleApi struct {
	vehicleService interfaces.VehicleService
}

func RegisterVehicleRoutes(app *gin.Engine, vehicleService interfaces.VehicleService) {
	service := vehicleApi{
		vehicleService: vehicleService,
	}

	app.POST("/vehicles", service.create)
	app.GET("/vehicles", service.search)
	app.GET("/vehicles/:vehicle_id", service.get)
	app.PATCH("/vehicles/:vehicle_id", service.update)
	app.POST("/vehicles/:vehicle_id/buy", service.buy)
}

// Create godoc
// @Summary Create Vehicle
// @Description Create a vehicle
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param user body vehicleApi.createVehicleRequest true "Body"
// @Success 201 {object} responses.Vehicle
// @Failure 204 {object} responses.ErrorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /vehicles [post]
func (ref *vehicleApi) create(ctx *gin.Context) {
	var request createVehicleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	vehicle, err := ref.vehicleService.Create(ctx, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusCreated, response)
}

// Create godoc
// @Summary Search vehicles
// @Description Seach vehicles
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param is_sold query boolean false "Filter vehicles by sold status"
// @Success 200 {array} responses.Vehicle
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /vehicles [get]
func (ref *vehicleApi) search(ctx *gin.Context) {
	var query vehicleQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	vehicles, err := ref.vehicleService.Search(ctx, query.IsSold)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response := make([]responses.Vehicle, len(vehicles))

	for i, vehicle := range vehicles {
		response[i] = responses.VehicleFromDomain(vehicle)
	}

	ctx.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Get Vehicle
// @Description Get a vehicle
// @Tags Vehicle
// @Accept json
// @Produce json
// @Success 200 {object} responses.Vehicle
// @Failure 204 {object} responses.ErrorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /vehicles/{vehicle_id} [get]
func (ref *vehicleApi) get(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	vehicle, err := ref.vehicleService.GetByID(ctx, uri.VehicleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Update Vehicle
// @Description Update a vehicle
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param user body vehicleApi.updateVehicleRequest false "Body"
// @Success 200 {object} responses.Vehicle
// @Failure 204 {object} responses.ErrorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /vehicles/{vehicle_id} [patch]
func (ref *vehicleApi) update(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var request updateVehicleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	vehicle, err := ref.vehicleService.Update(ctx, uri.VehicleID, *request.ToDomain())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Buy Vehicle
// @Description Buy a vehicle
// @Tags Vehicle
// @Accept json
// @Produce json
// @Param user body vehicleApi.buyVehicleRequest true "Body"
// @Success 200 {object} responses.Vehicle
// @Failure 204 {object} responses.ErrorResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /vehicles/{vehicle_id}/buy [post]
func (ref *vehicleApi) buy(ctx *gin.Context) {
	var uri vehicleURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var body buyVehicleRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	vehicle, err := ref.vehicleService.Buy(ctx, uri.VehicleID, body.BuyerDocumentNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if vehicle == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	response := responses.VehicleFromDomain(*vehicle)
	ctx.JSON(http.StatusOK, response)
}
