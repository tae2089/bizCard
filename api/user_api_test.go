package api_test

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/ent"
	mockapp "bizCard/mock/application"
	mockrepo "bizCard/mock/repository"
	"bizCard/router"
	"bizCard/util"
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

type UserApiTestSuite struct {
	suite.Suite
	UserRegister   domain.UserRegister
	User           *ent.User
	UserService    mockapp.MockUserService
	UserInfo       domain.UserInfo
	UserRepository mockrepo.MockUserRepository
	E              *httpexpect.Expect
	Data           map[string]interface{}
	AccessToken    string
}

func (ets *UserApiTestSuite) SetupTest() {
	handler := router.SetupRouter()

	ets.UserService = mockapp.MockUserService{}
	application.UserServiceBean = &ets.UserService

	ets.E = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(ets.T()),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(ets.T(), true),
		},
	})
	ets.Data = map[string]interface{}{
		"name":     "test",
		"email":    "test@example.com",
		"password": "test",
	}
	ets.UserRegister = domain.UserRegister{
		Name:     "test",
		Email:    "test@example.com",
		Password: "test",
	}

	ets.UserInfo = domain.UserInfo{
		Name:         "test",
		Email:        "test@example.com",
		Present:      true,
		ModifiedDate: time.Now(),
		CreatedDate:  time.Now(),
	}
	ets.AccessToken, _ = util.CreateJwt(domain.UserInfo{
		Name:  "tae2089",
		Email: "test@example.com",
	}, 1)
}

func (ets *UserApiTestSuite) TestRegisterUser() {
	ets.UserService.On("RegisterUser", mock.Anything).Return(ets.UserInfo)
	ets.E.POST("/user/register").
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).
		Expect().
		Status(200).
		JSON().
		Path("$.data.name").Equal("test")
}

func (ets *UserApiTestSuite) TestRegisterUser_error() {
	ets.UserService.On("RegisterUser", mock.Anything).Return(domain.UserInfo{Present: false})
	ets.E.POST("/user/register").
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).
		Expect().
		Status(500)
}

func (ets *UserApiTestSuite) TestLoginUser() {
	ets.UserService.On("LoginUser", mock.Anything).Return(ets.UserInfo, 1)
	ets.E.POST("/user/login").
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).
		Expect().
		Status(200).
		JSON().
		Path("$.data.name").Equal("test")
}

func (ets *UserApiTestSuite) TestLoginUser_error() {
	ets.UserService.On("LoginUser", mock.Anything).Return(domain.UserInfo{Present: false}, 0)
	ets.E.POST("/user/login").
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).
		Expect().
		Status(500)
}

func (ets *UserApiTestSuite) TestHealthCheck() {
	ets.E.GET("/user/test").
		WithCookie("accessToken", ets.AccessToken).
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(200)
}

func TestUserApiTestSuite(t *testing.T) {
	suite.Run(t, new(UserApiTestSuite))
}
