package router

import (
	"bizCard/api"
	"bizCard/application"
	"bizCard/repository"
	"github.com/gin-gonic/gin"
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
	return router
}

func SetupService() {
	repository.RegisterRepositoryBeans()
	application.SetupBizCardService()
}
