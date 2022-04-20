package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"bizCard/util"
	"context"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"log"
)

var _ BizCardService = (*BizCardServiceImpl)(nil)

type BizCardServiceImpl struct {
	BizCardRepository repository.BizCardRepository
	Log               *zap.Logger
}

func (b *BizCardServiceImpl) RegisterBizCard(dto *domain.BizCardRegister, ctx context.Context) *domain.BizCardInfo {
	childCtx, span := otel.Tracer("Register BizCard").Start(context.Background(), "Register BizCard Service")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start register bizCard Service"))
	bizCard, err := b.BizCardRepository.RegisterBizCard(dto, childCtx)

	if err != nil {
		log.Println(err)
		return nil
	}
	bizCardInfo := domain.CreateBizCardInfo(bizCard)
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start register bizCard Service"))
	return &bizCardInfo
}

func (b *BizCardServiceImpl) FindBizCard(uid int, ctx context.Context) *domain.BizCardInfo {
	childCtx, span := otel.Tracer("Find BizCard").Start(context.Background(), "Find BizCard Service")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start find bizCard Service"))
	bizCard, err := b.BizCardRepository.FindBIzCardByUid(uid, childCtx)
	if err != nil {
		log.Println(err)
		return nil
	}
	bizCardInfo := domain.CreateBizCardInfo(bizCard)
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end find bizCard Service"))
	return &bizCardInfo
}

func (b *BizCardServiceImpl) UpdateBizCard(uid int, dto *domain.BizCardUpdate, ctx context.Context) *domain.BizCardInfo {
	childCtx, span := otel.Tracer("Update BizCard").Start(context.Background(), "Update BizCard Service")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start update bizCard Service"))
	findBizCard, err := b.BizCardRepository.FindBIzCardByUid(uid, childCtx)
	if err != nil {
		log.Println("not found bizcard")
		return nil
	}
	bizCardUpdate := domain.CreateBizCardUpdate(findBizCard)
	bizCardUpdate.Update(dto)
	updateBizCard, err := b.BizCardRepository.UpdateBizCard(findBizCard.ID, bizCardUpdate, childCtx)
	if err != nil {
		log.Println(err)
		return nil
	}
	bizCardInfo := domain.CreateBizCardInfo(updateBizCard)
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end update bizCard Service"))
	return &bizCardInfo
}

func (b *BizCardServiceImpl) DeleteBizCard(uid int, ctx context.Context) string {
	childCtx, span := otel.Tracer("Delete BizCard").Start(context.Background(), "Delete BizCard Service")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start delete bizCard Service"))
	err := b.BizCardRepository.DeleteBizCardByUid(uid, childCtx)
	if err != nil {
		return "fail"
	}
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end delete bizCard Service"))
	return "success"
}
