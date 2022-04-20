package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"bizCard/ent/user"
	"bizCard/util"
	"context"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"time"
)

var _ UserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
	Client *ent.UserClient
}

func (u *UserRepositoryImpl) RegisterUser(userRegister domain.UserRegister, ctx context.Context) (*ent.User, error) {
	_, span := otel.Tracer("Register User").Start(ctx, "Register User Repository")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start RegisterUser Repository"))
	savedUser, err := u.Client.Create().
		SetName(userRegister.Name).
		SetEmail(userRegister.Email).
		SetPassword(userRegister.Password).
		SetCreatedDate(time.Now()).
		SetModifiedDate(time.Now()).
		Save(context.Background())
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end RegisterUser Repository"))
	return savedUser, err
}

func (u *UserRepositoryImpl) FindUser(email string, ctx context.Context) (*ent.User, error) {
	_, span := otel.Tracer("Login User").Start(ctx, "Find User Repository")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start FindUser Repository"))
	findUser, err := u.Client.Query().Where(user.Email(email)).Only(context.Background())
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end FindUser Repository"))
	return findUser, err
}
