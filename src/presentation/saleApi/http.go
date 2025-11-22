package saleApi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/responses"
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
// @Success 200 {array} responses.Sale
// @Failure 500 {object} responses.ErrorResponse
// @Router /sales [get]
func (ref *saleApi) search(ctx *gin.Context) {
	sales, err := ref.saleService.Search(ctx)
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
// @Param user body saleApi.saleWebhookRequest true "Body"
// @Success 204
// @Failure 400 {object} responses.ErrorResponse
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

	_, err := ref.saleService.UpdateStatusByPaymentID(ctx, request.PaymentID, request.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
