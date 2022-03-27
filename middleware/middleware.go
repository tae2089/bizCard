package middleware

import (
	"bizCard/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DummyMiddleWare(c *gin.Context) {
	util.Log.Info("controller start")
	c.Next()
	util.Log.Info("controller end")
}

func CheckJwtToken(c *gin.Context) {
	util.Log.Info("CheckJwtToken Middleware start")
	token, err := c.Request.Cookie("accessToken")
	if err != nil {
		util.Log.Error("Authentication failed")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Authentication failed"})
		c.Abort()
		return
	}
	tokenString := token.Value
	if tokenString == "" {
		util.Log.Error("token is empty")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "token is none"})
		c.Abort()
		return
	}

	id := util.ParseJwt(tokenString)
	if id == 0 {
		util.Log.Error("token parsing error")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "token parsing error"})
		c.Abort()
		return
	} else if id == -1 {
		util.Log.Error("token expired")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "token error"})
		c.Abort()
		return
	} else {
		c.Set("userId", id)
	}
	util.Log.Info("CheckJwtToken Middleware End")
	c.Next()
}
