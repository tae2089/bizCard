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
	"log"
	"net/http"
	"testing"
)

type BizCardApiTestSuite struct {
	suite.Suite
	BizCardDto        domain.BizCardRegister
	BizCard           *ent.BizCard
	BizCardService    mockapp.MockBizCardService
	BizCardInfo       *domain.BizCardInfo
	BizCardRepository mockrepo.MockBizCardRepository
	E                 *httpexpect.Expect
	Data              map[string]interface{}
	AccessToken       string
}

func (ets *BizCardApiTestSuite) SetupTest() {
	handler := router.SetupRouter()
	log.Println("111", ets.T())
	ets.BizCardService = mockapp.MockBizCardService{}
	application.BizCardServiceBean = &ets.BizCardService

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
	ets.BizCardInfo = &domain.BizCardInfo{
		Email:       "tae2089",
		Name:        "taebin",
		PhoneNumber: "010-xxxx-xxxx",
		Age:         25,
	}
	ets.Data = map[string]interface{}{
		"name":        "taebin",
		"email":       "tae2089",
		"phoneNumber": "010-xxxx-xxxx",
		"age":         25,
	}
	ets.AccessToken, _ = util.CreateJwt(domain.UserInfo{
		Name:  "tae2089",
		Email: "test@example.com",
	}, 1)
}

func (ets *BizCardApiTestSuite) TestRegisterBizCard() {
	ets.BizCardService.On("RegisterBizCard", mock.Anything).Return(ets.BizCardInfo)
	ets.E.POST("/bizcard/register").
		WithCookie("accessToken", ets.AccessToken).
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).
		Expect().
		Status(200).
		JSON().
		Path("$.data.name").Equal("taebin")
}

func (ets *BizCardApiTestSuite) TestFindBizCard() {

	ets.BizCardService.On("FindBizCard", mock.Anything).Return(ets.BizCardInfo)
	ets.E.GET("/bizcard/1").
		WithCookie("accessToken", ets.AccessToken).
		WithHeader("Content-Type", "application/json").
		Expect().
		Status(200).
		JSON().
		Path("$.data.name").Equal("taebin")

}

func (ets *BizCardApiTestSuite) TestUpdateBizCard() {
	ets.BizCardInfo.Age = 26
	ets.BizCardService.On("UpdateBizCard", mock.AnythingOfType("int"), mock.Anything).Return(ets.BizCardInfo)
	ets.E.PUT("/bizcard/1").
		WithCookie("accessToken", ets.AccessToken).
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).Expect().
		Status(200).
		JSON().
		Path("$.data.name").Equal("taebin")
}

func (ets *BizCardApiTestSuite) TestDeleteBizCard() {
	ets.BizCardService.On("DeleteBizCard", mock.AnythingOfType("int")).Return("success")
	ets.E.DELETE("/bizcard/1").
		WithCookie("accessToken", ets.AccessToken).
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).Expect().
		Status(200).
		JSON().
		Path("$.data").Equal("success")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BizCardApiTestSuite))
}
