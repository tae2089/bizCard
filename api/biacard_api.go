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
	result := domain.Success()
	if err := c.ShouldBindJSON(&dto); err != nil {
		result = domain.Fail()
		result.Data = err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	log.Println("111", dto)
	data := application.BizCardServiceBean.RegisterBizCard(&dto)
	result.Data = data
	log.Println("d", result)

	c.JSON(http.StatusOK, result)
}

func FindBizCard(c *gin.Context) {
	uid := c.Param("uid")
	intUid, err := strconv.Atoi(uid)
	result := domain.Success()
	if err != nil {
		result = domain.Fail()
		result.Data = err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	data := application.BizCardServiceBean.FindBizCard(intUid)
	result.Data = data
	c.JSON(http.StatusOK, result)
}

func UpdateBizCard(c *gin.Context) {
	uid := c.Param("uid")
	intUid, err := strconv.Atoi(uid)
	result := domain.Success()
	if err != nil {
		result = domain.Fail()
		result.Data = err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	var update domain.BizCardUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := application.BizCardServiceBean.UpdateBizCard(intUid, &update)
	result.Data = data
	c.JSON(http.StatusOK, result)
}

func DeleteBizCard(c *gin.Context) {
	uid := c.Param("uid")
	intUid, err := strconv.Atoi(uid)
	result := domain.Success()
	if err != nil {
		result = domain.Fail()
		result.Data = err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	data := application.BizCardServiceBean.DeleteBizCard(intUid)
	result.Data = data
	c.JSON(http.StatusOK, result)
}
