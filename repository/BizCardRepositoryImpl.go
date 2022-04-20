package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"bizCard/ent/bizcard"
	"bizCard/util"
	"context"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ BizCardRepository = (*BizCardRepositoryImpl)(nil)

type BizCardRepositoryImpl struct {
	Client *ent.BizCardClient
}

func (b *BizCardRepositoryImpl) RegisterBizCard(dto *domain.BizCardRegister, ctx context.Context) (*ent.BizCard, error) {
	_, span := otel.Tracer("Register BizCard").Start(context.Background(), "Register BizCard Repository")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start register bizCard Repository"))
	savedBizCard, err := b.Client.
		Create().
		SetAge(dto.Age).
		SetName(dto.Name).
		SetEmail(dto.Email).
		SetPhoneNumber(dto.PhoneNumber).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end register bizCard Repository"))
	return savedBizCard, nil
}

func (b *BizCardRepositoryImpl) FindBIzCardByUid(uid int, ctx context.Context) (*ent.BizCard, error) {
	_, span := otel.Tracer("Find BizCard").Start(context.Background(), "Find BizCard Repository")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start find bizCard Repository"))
	bizCard, err := b.Client.Query().Where(bizcard.ID(uid)).First(context.Background())
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end find bizCard Repository"))
	return bizCard, err
}

func (b *BizCardRepositoryImpl) UpdateBizCard(uid int, bizCardUpdate *domain.BizCardUpdate, ctx context.Context) (*ent.BizCard, error) {
	_, span := otel.Tracer("Update BizCard").Start(context.Background(), "Update BizCard Repository")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start update bizCard Repository"))
	bizCard, err := b.Client.UpdateOneID(uid).
		SetAge(bizCardUpdate.Age).
		SetEmail(bizCardUpdate.Email).
		SetName(bizCardUpdate.Name).
		SetPhoneNumber(bizCardUpdate.PhoneNumber).
		Save(context.Background())
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end update bizCard Repository"))
	return bizCard, err
}

func (b *BizCardRepositoryImpl) DeleteBizCardByUid(uid int, ctx context.Context) error {
	_, span := otel.Tracer("Delete BizCard").Start(context.Background(), "Delete BizCard Repository")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start delete bizCard Repository"))
	err := b.Client.DeleteOneID(uid).Exec(context.Background())
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end delete bizCard Repository"))
	return err
}
