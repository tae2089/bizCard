package application_test

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/ent"
	mockrepo "bizCard/mock/repository"
	"bizCard/trace"
	"bizCard/util"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel"
	"testing"
	"time"
)

type UserServiceTestSuite struct {
	suite.Suite
	UserRegister   domain.UserRegister
	User           *ent.User
	UserService    application.UserService
	UserRepository mockrepo.MockUserRepository
}

func (ets *UserServiceTestSuite) SetupTest() {
	tp := trace.InitTestTrace()
	otel.SetTracerProvider(tp)
	util.SetupLogging()
	ets.UserRegister = domain.UserRegister{
		Name:     "tester",
		Password: "hello01",
		Email:    "tester@gmain.com",
	}
	ets.User = &ent.User{
		Name:         "tester",
		Password:     "ttqwem",
		Email:        "tester@gmain.com",
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}
	ets.UserRepository = mockrepo.MockUserRepository{}
	ets.UserService = &application.UserServiceImpl{UserRepository: &ets.UserRepository}
}
func (ets *UserServiceTestSuite) TestUserServiceImpl_RegisterUser() {
	ctx, span := otel.Tracer("Register User").Start(context.Background(), "Register User Application")
	defer span.End()
	ets.UserRepository.On("FindUser", mock.AnythingOfType("string"), mock.Anything).Return(&ent.User{Email: ""}, errors.New("user not found"))
	ets.UserRepository.On("RegisterUser", mock.Anything, mock.Anything).Return(ets.User, nil)
	result := ets.UserService.RegisterUser(ets.UserRegister, ctx)
	ets.Equal("tester", result.Name)
}

func (ets *UserServiceTestSuite) TestUserServiceImpl_RegisterUser_Email_Exist() {
	ctx, span := otel.Tracer("Register User").Start(context.Background(), "Register User Application")
	defer span.End()
	ets.UserRepository.On("FindUser", mock.AnythingOfType("string"), mock.Anything).Return(&ent.User{}, nil)
	ets.UserRepository.On("RegisterUser", mock.Anything).Return(&ets.User, nil)
	result := ets.UserService.RegisterUser(ets.UserRegister, ctx)
	ets.Equal(false, result.Present)
}

func (ets *UserServiceTestSuite) TestUserServiceImpl_FindUser() {
	ctx, span := otel.Tracer("Login User").Start(context.Background(), "Login User Application")
	defer span.End()
	ets.UserRepository.On("FindUser", mock.Anything, mock.Anything).Return(&ent.User{
		ID:           1,
		Name:         "tester",
		Email:        "test@example.com",
		Password:     "$2a$10$9DbcpPw2hfYmqzYtQIp3t.nA.YBFPIB7jGbD87AbKrCGg/BoP7B.i",
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}, nil)
	result, id := ets.UserService.LoginUser(domain.UserLoginForm{Email: "test@example.com", Password: "hello01"}, ctx)
	ets.Equal(id, 1)
	ets.Equal(result.Name, "tester")
}

func (ets *UserServiceTestSuite) TestUserServiceImpl_NotFindUser() {
	ctx, span := otel.Tracer("Login User").Start(context.Background(), "Login User Application")
	defer span.End()
	ets.UserRepository.On("FindUser", mock.Anything, mock.Anything).Return(&ent.User{
		ID:           1,
		Name:         "tester",
		Email:        "test@example.com",
		Password:     "$2a$10$9DbcpPw2hfYmqzYtQIp3t.nA.YBFPIB7jGbD87AbKrCGg/BoP7B.i",
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}, nil)
	result, id := ets.UserService.LoginUser(domain.UserLoginForm{Email: "test@example.com", Password: "hello02"}, ctx)
	ets.Equal(id, 0)
	ets.Equal(result.Present, false)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
