package application_test

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/ent"
	mockrepo "bizCard/mock/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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
	ets.UserRepository.On("FindUser", mock.AnythingOfType("string")).Return(ent.User{Email: ""}, nil)
	ets.UserRepository.On("RegisterUser", mock.Anything).Return(ets.User, nil)
	result := ets.UserService.RegisterUser(ets.UserRegister)
	ets.Equal("tester", result.Name)
}

func (ets *UserServiceTestSuite) TestUserServiceImpl_RegisterUser_Email_Exist() {
	ets.UserRepository.On("FindUser", mock.AnythingOfType("string")).Return(ent.User{}, nil)
	ets.UserRepository.On("RegisterUser", mock.Anything).Return(ets.User, nil)
	result := ets.UserService.RegisterUser(ets.UserRegister)
	ets.Equal(false, result.Present)
}

func (ets *UserServiceTestSuite) TestUserServiceImpl_FindUser() {
	ets.UserRepository.On("FindUser", mock.Anything).Return(ent.User{
		ID:           1,
		Name:         "tester",
		Email:        "test@example.com",
		Password:     "$2a$10$9DbcpPw2hfYmqzYtQIp3t.nA.YBFPIB7jGbD87AbKrCGg/BoP7B.i",
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}, nil)
	result, id := ets.UserService.LoginUser(domain.UserLoginForm{Email: "test@example.com", Password: "hello01"})
	ets.Equal(id, 1)
	ets.Equal(result.Name, "tester")
}

func (ets *UserServiceTestSuite) TestUserServiceImpl_NotFindUser() {
	ets.UserRepository.On("FindUser", mock.Anything).Return(ent.User{
		ID:           1,
		Name:         "tester",
		Email:        "test@example.com",
		Password:     "$2a$10$9DbcpPw2hfYmqzYtQIp3t.nA.YBFPIB7jGbD87AbKrCGg/BoP7B.i",
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}, nil)
	result, id := ets.UserService.LoginUser(domain.UserLoginForm{Email: "test@example.com", Password: "hello02"})
	ets.Equal(id, 0)
	ets.Equal(result.Present, false)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
