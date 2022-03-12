package api

import (
	"bizCard/application"
	"bizCard/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func RegisterBizCard(c *gin.Context) {
	var dto domain.BizCardRegister
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("111", dto)
	result := application.BizCardServiceBean.RegisterBizCard(&dto)
	log.Println("d", result)

	c.JSON(http.StatusOK, result)
}

func FindBizCard(c *gin.Context) {
	uid := c.Param("uid")
	intUid, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result := application.BizCardServiceBean.FindBizCard(intUid)
	c.JSON(http.StatusOK, result)
}

func UpdateBizCard(c *gin.Context) {
	uid := c.Param("uid")
	intUid, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var update domain.BizCardUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := application.BizCardServiceBean.UpdateBizCard(intUid, &update)
	c.JSON(http.StatusOK, result)
}
