package api

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/util"
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
)

func RegisterBizCard(c *gin.Context) {
	ctx, span := otel.Tracer("Register BizCard").Start(context.Background(), "Register BizCard Application")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start register bizCard api"))
	var dto domain.BizCardRegister
	result := domain.Success()
	if err := c.ShouldBindJSON(&dto); err != nil {
		result = domain.Fail()
		result.Data = err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	log.Println("111", dto)
	data := application.BizCardServiceBean.RegisterBizCard(&dto, ctx)
	result.Data = data
	log.Println("d", result)
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end register bizCard api"))
	c.JSON(http.StatusOK, result)
}

func FindBizCard(c *gin.Context) {
	ctx, span := otel.Tracer("Find BizCard").Start(context.Background(), "Find BizCard Application")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start find bizCard api"))
	uid := c.Param("uid")
	intUid, err := strconv.Atoi(uid)
	result := domain.Success()
	if err != nil {
		result = domain.Fail()
		result.Data = err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	data := application.BizCardServiceBean.FindBizCard(intUid, ctx)
	result.Data = data
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end find bizCard api"))
	c.JSON(http.StatusOK, result)
}

func UpdateBizCard(c *gin.Context) {
	ctx, span := otel.Tracer("Update BizCard").Start(context.Background(), "Update BizCard Application")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start update bizCard api"))
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
	data := application.BizCardServiceBean.UpdateBizCard(intUid, &update, ctx)
	result.Data = data
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end update bizCard api"))
	c.JSON(http.StatusOK, result)
}

func DeleteBizCard(c *gin.Context) {
	ctx, span := otel.Tracer("Delete BizCard").Start(context.Background(), "Delete BizCard Application")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start delete bizCard api"))
	uid := c.Param("uid")
	intUid, err := strconv.Atoi(uid)
	result := domain.Success()
	if err != nil {
		result = domain.Fail()
		result.Data = err.Error()
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	data := application.BizCardServiceBean.DeleteBizCard(intUid, ctx)
	result.Data = data
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end delete bizCard api"))
	c.JSON(http.StatusOK, result)
}
