package router

import (
	"github.com/gin-gonic/gin"
	"main/api"
	"main/application"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/register", api.RegisterBizCard)
	setupService()
	return router
}

func setupService() {
	application.SetupBizCardService()
}
