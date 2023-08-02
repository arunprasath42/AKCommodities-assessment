package route

import (
	"web-api/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(router *gin.Engine) {

	/***BASEPATH OF AN API. NOTE:THIS SHOULDN'T BE CHANGED***/
	api := router.Group("/api")

	/***ADD THE ROUTES HERE***/
	api.GET("/insert-into-db", controllers.FetchHistoricalExchangeRates)
	api.GET("/fetchexchange-rates", controllers.GetExchangeRates)

	router.Run(viper.GetString("server.port"))
}
