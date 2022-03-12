package api_test

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/ent"
	mockapp "bizCard/mock/application"
	mockrepo "bizCard/mock/repository"
	"bizCard/router"
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

}

func (ets *BizCardApiTestSuite) TestRegisterBizCard() {
	ets.BizCardService.On("RegisterBizCard", mock.Anything).Return(ets.BizCardInfo)
	ets.E.POST("/register").
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).Expect().
		JSON().
		Object().
		ContainsKey("name").
		ValueEqual("name", "taebin")
}

func (ets *BizCardApiTestSuite) TestFindBizCard() {

	ets.BizCardService.On("FindBizCard", mock.Anything).Return(ets.BizCardInfo)
	ets.E.GET("/1").
		WithHeader("Content-Type", "application/json").
		Expect().
		JSON().
		Object().
		ContainsKey("name").
		ValueEqual("name", "taebin")
}

func (ets *BizCardApiTestSuite) TestUpdateBizCard() {
	ets.BizCardInfo.Age = 26
	ets.BizCardService.On("UpdateBizCard", mock.AnythingOfType("int"), mock.Anything).Return(ets.BizCardInfo)
	ets.E.PUT("/1").
		WithHeader("Content-Type", "application/json").
		WithJSON(ets.Data).Expect().
		JSON().
		Object().
		ContainsKey("age").
		ValueEqual("age", 26)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BizCardApiTestSuite))
}
