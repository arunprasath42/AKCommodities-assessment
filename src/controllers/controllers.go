// controllers/exchange_rates.go
package controllers

import (
	"net/http"
	"web-api/src/service"
	"web-api/utils/constant"
	"web-api/utils/response"

	"github.com/gin-gonic/gin"
)

// API to Read data from the database
func FetchHistoricalExchangeRates(c *gin.Context) {

	var service service.ExchangeRatesService
	exchangeRates, err := service.FetchHistoricalExchangeRates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(exchangeRates))
}

// GetExchangeRates -retrieves exchange rates based on query parameters
func GetExchangeRates(c *gin.Context) {
	baseCurrency := c.DefaultQuery("base", "")         // Extract base currency from query parameter
	specificCurrency := c.DefaultQuery("currency", "") // Extract specific currency from query parameter
	var service service.ExchangeRatesService

	exchangeRates, err := service.GetExchangeRates(baseCurrency, specificCurrency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(exchangeRates))
}
