package saleApi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/responses"
	"github.com/caiiomp/vehicle-platform-sales/src/presentation/constants"
)

type saleApi struct {
	saleService interfaces.SaleService
}

func RegisterSaleRoutes(app *gin.Engine, saleService interfaces.SaleService) {
	service := saleApi{
		saleService: saleService,
	}

	app.GET("/sales", service.search)
	app.POST("/sales/webhook", service.webhook)
}

// Create godoc
// @Summary List sales
// @Description List sales
// @Tags Sale
// @Accept json
// @Produce json
// @Param status query string false "Filter sales by status"
// @Success 200 {array} responses.Sale
// @Failure 500 {object} responses.ErrorResponse
// @Router /sales [get]
func (ref *saleApi) search(ctx *gin.Context) {
	var query saleQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	sales, err := ref.saleService.SearchByStatus(ctx, query.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response := make([]responses.Sale, len(sales))

	for i, sale := range sales {
		response[i] = responses.SaleFromDomain(sale)
	}

	ctx.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Sale Webhook
// @Description Sale Webhook
// @Tags Sale
// @Accept json
// @Produce json
// @Param expected_webhook body saleApi.saleWebhookRequest true "Body"
// @Success 204
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /sales/webhook [post]
func (ref *saleApi) webhook(ctx *gin.Context) {
	var request saleWebhookRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	sale, err := ref.saleService.UpdateStatusByPaymentID(ctx, request.PaymentID, request.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if sale == nil {
		ctx.JSON(http.StatusNotFound, responses.ErrorResponse{
			Error: constants.SaleDoesNotExist,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
