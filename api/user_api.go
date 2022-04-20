package api

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/util"
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	ctx, span := otel.Tracer("Register User").Start(context.Background(), "Register User Application")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start register user api"))
	var userRegister domain.UserRegister
	result := domain.Success()
	if err := c.BindJSON(&userRegister); err != nil {
		result = domain.Fail()
		result.Data = err
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	data := application.UserServiceBean.RegisterUser(userRegister, ctx)
	if data.Present == false {
		result = domain.Fail()
		result.Data = "user register failed"
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	result.Data = data
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end register user api"))
	c.JSON(200, result)
}

func LoginUser(c *gin.Context) {
	ctx, span := otel.Tracer("Login User").Start(context.Background(), "Login User Application")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start login user api"))
	var userLoginForm domain.UserLoginForm
	result := domain.Success()

	if err := c.ShouldBind(&userLoginForm); err != nil {
		result = domain.Fail()
		result.Data = err
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	data, id := application.UserServiceBean.LoginUser(userLoginForm, ctx)
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end login user service"))

	if data.Present == false {
		result = domain.Fail()
		result.Data = "Login Fail"
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	token, err := util.CreateJwt(data, id)
	if err != nil {
		result = domain.Fail()
		result.Data = err
		c.JSON(http.StatusBadRequest, result)
		return
	}
	c.SetCookie("accessToken", token, 60*60*24, "/", "localhost", true, true)
	result.Data = data
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end login user api"))

	c.JSON(200, result)
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "OK"})
}
