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
	ets.UserRepository.On("RegisterUser", mock.Anything).Return(ets.User, nil)
	result := ets.UserService.RegisterUser(ets.UserRegister)
	ets.Equal("tester", result.Name)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
