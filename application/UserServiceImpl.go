package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"bizCard/util"
	"context"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (userServiceImpl *UserServiceImpl) LoginUser(loginForm domain.UserLoginForm, ctx context.Context) (domain.UserInfo, int) {
	childCtx, span := otel.Tracer("Login User").Start(ctx, "Login User Service")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start login user service"))
	findUser, err := userServiceImpl.UserRepository.FindUser(loginForm.Email, childCtx)
	if err != nil {
		util.Log.Error(err.Error())
		return domain.UserInfo{Present: false}, 0
	}
	comparePassword, err := util.Compare(findUser.Password, loginForm.Password)
	if err != nil || comparePassword == false {
		util.Log.Error(err.Error())
		return domain.UserInfo{Present: false}, 0
	}
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start login user service"))
	return domain.CreateUserInfo(findUser), findUser.ID
}

func (userServiceImpl *UserServiceImpl) RegisterUser(userRegister domain.UserRegister, ctx context.Context) domain.UserInfo {
	childCtx, span := otel.Tracer("Register User").Start(ctx, "Register User Service")
	defer span.End()
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "start login user service"))
	_, err := userServiceImpl.UserRepository.FindUser(userRegister.Email, childCtx)
	if err == nil {
		return domain.UserInfo{
			Present: false,
		}
	}

	encryptPassword, err := util.GenerateBcrypt(userRegister.Password)
	if err != nil {
		util.Log.Error(err.Error())
		return domain.UserInfo{
			Present: false,
		}
	}
	userRegister.Password = encryptPassword
	savedUser, err := userServiceImpl.UserRepository.RegisterUser(userRegister, childCtx)
	if err != nil {
		util.Log.Error(err.Error())
		return domain.UserInfo{
			Present: false,
		}
	}
	userInfo := domain.CreateUserInfo(savedUser)
	util.Log.Info("", zap.String("traceId", span.SpanContext().TraceID().String()), zap.String("spanId", span.SpanContext().SpanID().String()), zap.String("msg", "end login user service"))
	return userInfo
}
