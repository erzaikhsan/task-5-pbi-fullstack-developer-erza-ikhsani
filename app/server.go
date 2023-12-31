package app

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/database"
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	route "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	/**
	@description Setup Database Connection
	*/
	db := database.Connection()
	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if helper.GodotEnv("GO_ENV") != "production" && helper.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if helper.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	/**
	@description Photo access
	*/
	router.Static("/public/photos", "./public/photos")
	router.GET("api/v1/photos/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		c.File("./public/photos/" + filename)
	})

	/**
	@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	/**
	@description Init All Route
	*/
	route.InitAuthRoutes(db, router)
	route.InitUserRoutes(db, router)
	route.InitPhotoRoutes(db, router)

	return router
}
