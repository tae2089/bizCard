package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/application"
	"main/domain"
	"net/http"
)

func RegisterBizCard(c *gin.Context) {
	var dto domain.BizCardRegister
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("111", dto)
	result := application.BizCardService2.RegisterBizCard(dto)
	log.Println("d", result)

	c.JSON(http.StatusOK, result)
}
