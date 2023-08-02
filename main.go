package main

import (
	"web-api/config"
	"web-api/route"
	"web-api/src/repository"
	"web-api/utils/database"
	"web-api/utils/middleware"
	"web-api/utils/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	router := gin.Default()
	database.GetInstancemysql()
	repository.MySqlInit()
	validator.Init()
	router.Use(middleware.TracingMiddleware())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	route.SetupRoutes(router)

}
