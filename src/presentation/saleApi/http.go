package saleApi

import (
	"net/http"

	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
	"github.com/caiiomp/vehicle-platform-sales/src/core/responses"
	"github.com/gin-gonic/gin"
)

type saleApi struct {
	saleService interfaces.SaleService
}

func RegisterSaleRoutes(app *gin.Engine, saleService interfaces.SaleService) {
	service := saleApi{
		saleService: saleService,
	}

	app.GET("/sales", service.search)
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
