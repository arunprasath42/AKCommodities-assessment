// controllers/exchange_rates.go
package controllers

import (
	"net/http"
	"web-api/src/service"
	"web-api/utils/constant"
	"web-api/utils/response"

	"github.com/gin-gonic/gin"
)

// FetchHistoricalExchangeRates -inserts exchange rates into database from openexchangerates
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
	baseCurrency := c.DefaultQuery("base", "")
	specificCurrency := c.DefaultQuery("currency", "")
	var service service.ExchangeRatesService

	exchangeRates, err := service.GetExchangeRates(baseCurrency, specificCurrency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(exchangeRates))
}
