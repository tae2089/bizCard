package router

import (
	"bizCard/api"
	"bizCard/application"
	"bizCard/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	bizcard := router.Group("/bizcard")
	user := router.Group("/user")
	SetupBizCardApi(bizcard)
	SetupUserApi(user)
	return router
}

func SetupBizCardApi(group *gin.RouterGroup) {
	group.POST("/register", api.RegisterBizCard)
	group.GET("/:uid", api.FindBizCard)
	group.PUT("/:uid", api.UpdateBizCard)
	group.DELETE("/:uid", api.DeleteBizCard)
}

func SetupUserApi(group *gin.RouterGroup) {
	group.POST("/register", api.RegisterUser)
}

func SetupService() {
	repository.RegisterRepositoryBeans()
	application.RegisterServiceBeans()
}
