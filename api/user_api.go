package api

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/util"
	"errors"
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
	if data.Present == false {
		c.JSON(500, errors.New("user register fail"))
	}
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
	data, id := application.UserServiceBean.LoginUser(userLoginForm)

	if data.Present == false {
		c.JSON(500, errors.New("Login Fail"))
	}

	token, err := util.CreateJwt(data, id)
	if err != nil {
		result = domain.Fail()
		result.Data = err
		c.JSON(500, result)
	}
	c.SetCookie("accessToken", token, 60*60*24, "/", "localhost", true, true)
	result.Data = data
	c.JSON(200, result)
}
