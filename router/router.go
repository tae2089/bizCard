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
	bizcard = SetupBizCardApi(bizcard)
	return router
}

func SetupBizCardApi(group *gin.RouterGroup) *gin.RouterGroup {
	group.POST("/register", api.RegisterBizCard)
	group.GET("/:uid", api.FindBizCard)
	group.PUT("/:uid", api.UpdateBizCard)
	group.DELETE("/:uid", api.DeleteBizCard)
	return group
}

func SetupService() {
	repository.RegisterRepositoryBeans()
	application.SetupBizCardService()
}
