package application_test

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/ent"
	mockrepo "bizCard/mock/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BizCardServiceTestSuite struct {
	suite.Suite
	BizCardDto        domain.BizCardRegister
	BizCard           *ent.BizCard
	BizCardService    application.BizCardService
	BizCardRepository mockrepo.MockBizCardRepository
}

func (ets *BizCardServiceTestSuite) SetupTest() {
	ets.BizCardDto = domain.BizCardRegister{
		Email:       "tae2089",
		Name:        "taebin",
		PhoneNumber: "010-xxxx-xxxx",
		Age:         25,
	}
	ets.BizCard = &ent.BizCard{
		Email:       "tae2089",
		Name:        "taebin",
		PhoneNumber: "010-xxxx-xxxx",
		Age:         25,
	}
	ets.BizCardRepository = mockrepo.MockBizCardRepository{}
	ets.BizCardService = &application.BizCardServiceImpl{BizCardRepository: &ets.BizCardRepository}
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_RegisterBizCard() {
	ets.BizCardRepository.On("RegisterBizCard", mock.Anything).Return(ets.BizCard, nil)
	result := ets.BizCardService.RegisterBizCard(&ets.BizCardDto)
	ets.Equal("tae2089", result.Email)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_FindBIzCardByUid() {
	ets.BizCardRepository.On("FindBIzCardByUid", mock.Anything).Return(ets.BizCard, nil)
	result := ets.BizCardService.FindBizCard(1)
	ets.Equal("tae2089", result.Email)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_UpdateBizCard() {
	ets.BizCardRepository.On("FindBIzCardByUid", mock.Anything).Return(ets.BizCard, nil)
	ets.BizCard.Age = 26
	ets.BizCardRepository.On("UpdateBizCard", mock.Anything, mock.Anything).Return(ets.BizCard, nil)
	result := ets.BizCardService.UpdateBizCard(1, &domain.BizCardUpdate{
		Age: 26,
	})
	ets.Equal(26, result.Age)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_UpdateBizCard_NotfoundBizCard() {
	bizCardRepository := &mockrepo.MockBizCardRepository{}
	bizCardService := &application.BizCardServiceImpl{BizCardRepository: bizCardRepository}
	bizCardRepository.On("FindBIzCardByUid", mock.Anything).Return(nil, errors.New("not found bizcard"))
	result := bizCardService.UpdateBizCard(1, &domain.BizCardUpdate{
		Age: 26,
	})
	ets.Nil(result)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_UpdateBizCard_SavedError() {
	ets.BizCardRepository.On("FindBIzCardByUid", mock.Anything).Return(ets.BizCard, nil)
	ets.BizCardRepository.On("UpdateBizCard", mock.Anything, mock.Anything).Return(nil, errors.New("update error"))
	result := ets.BizCardService.UpdateBizCard(1, &domain.BizCardUpdate{
		Age: 26,
	})
	ets.Nil(result)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BizCardServiceTestSuite))
}

func TestBizCardUpdate(t *testing.T) {
	data := &ent.BizCard{
		Email:       "tae2089",
		Name:        "taebin",
		PhoneNumber: "010-xxxx-xxxx",
		Age:         25,
	}
	b := domain.CreateBizCardUpdate(data)
	b = b.Update(&domain.BizCardUpdate{
		Name: "tester",
		Age:  100,
	})
	assert.Equal(t, 100, b.Age)
}
