package router

import (
	"bizCard/api"
	"bizCard/application"
	"bizCard/middleware"
	"bizCard/repository"
	"bizCard/util"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	util.SetupLogging()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	bizcard := router.Group("/bizcard")
	bizcard.Use(middleware.CheckJwtToken)
	user := router.Group("/user")
	user.Use(middleware.DummyMiddleWare)
	tokenUser := router.Group("/user")
	tokenUser.Use(middleware.CheckJwtToken)
	SetupBizCardApi(bizcard)
	SetupUserApi(user)
	SetupTokenUserApi(tokenUser)
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
	group.POST("/login", api.LoginUser)
}

func SetupTokenUserApi(group *gin.RouterGroup) {
	group.GET("/test", api.HealthCheck)
}

func SetupService() {
	repository.RegisterRepositoryBeans()
	application.RegisterServiceBeans()
}
