package api

import (
	"bizCard/application"
	"bizCard/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterBizCard(c *gin.Context) {
	var dto domain.BizCardRegister
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("111", dto)
	result := application.BizCardServiceBean.RegisterBizCard(dto)
	log.Println("d", result)

	c.JSON(http.StatusOK, result)
}
