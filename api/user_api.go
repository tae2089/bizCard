package api

import (
	"bizCard/application"
	"bizCard/domain"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var userRegister domain.UserRegister
	result := domain.Success()
	if err := c.BindJSON(&userRegister); err != nil {
		result = domain.Fail()
		result.Data = err
		c.JSON(500, result)
	}
	data := application.UserServiceBean.RegisterUser(userRegister)
	result.Data = data
	c.JSON(200, result)
}

func LoginUser(c *gin.Context) {
	var userLoginForm domain.UserLoginForm
	result := domain.Success()
	if err := c.ShouldBind(&userLoginForm); err != nil {
		result = domain.Fail()
		result.Data = err
		c.JSON(500, result)
	}
	data := application.UserServiceBean.LoginUser(userLoginForm)
	result.Data = data
	c.JSON(200, result)
}
