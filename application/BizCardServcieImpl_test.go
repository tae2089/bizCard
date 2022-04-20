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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel"
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

	tp := trace.InitTestTrace()
	otel.SetTracerProvider(tp)
	util.SetupLogging()

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
	ctx, span := otel.Tracer("Register BizCard").Start(context.Background(), "Register BizCard Application")
	defer span.End()
	ets.BizCardRepository.On("RegisterBizCard", mock.Anything, mock.Anything).Return(ets.BizCard, nil)
	result := ets.BizCardService.RegisterBizCard(&ets.BizCardDto, ctx)
	ets.Equal("tae2089", result.Email)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_FindBIzCardByUid() {
	ctx, span := otel.Tracer("Find BizCard").Start(context.Background(), "Find BizCard Application")
	defer span.End()
	ets.BizCardRepository.On("FindBIzCardByUid", mock.Anything, mock.Anything).Return(ets.BizCard, nil)
	result := ets.BizCardService.FindBizCard(1, ctx)
	ets.Equal("tae2089", result.Email)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_UpdateBizCard() {
	ctx, span := otel.Tracer("Update BizCard").Start(context.Background(), "Update BizCard Application")
	defer span.End()
	ets.BizCardRepository.On("FindBIzCardByUid", mock.Anything, mock.Anything).Return(ets.BizCard, nil)
	ets.BizCard.Age = 26
	ets.BizCardRepository.On("UpdateBizCard", mock.Anything, mock.Anything, mock.Anything).Return(ets.BizCard, nil)
	result := ets.BizCardService.UpdateBizCard(1, &domain.BizCardUpdate{
		Age: 26,
	}, ctx)
	ets.Equal(26, result.Age)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_UpdateBizCard_NotfoundBizCard() {
	ctx, span := otel.Tracer("Update BizCard").Start(context.Background(), "Update BizCard Application")
	defer span.End()
	bizCardRepository := &mockrepo.MockBizCardRepository{}
	bizCardService := &application.BizCardServiceImpl{BizCardRepository: bizCardRepository}
	bizCardRepository.On("FindBIzCardByUid", mock.Anything, mock.Anything).Return(nil, errors.New("not found bizcard"))
	result := bizCardService.UpdateBizCard(1, &domain.BizCardUpdate{
		Age: 26,
	}, ctx)
	ets.Nil(result)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_UpdateBizCard_SavedError() {
	ctx, span := otel.Tracer("Update BizCard").Start(context.Background(), "Update BizCard Application")
	defer span.End()
	ets.BizCardRepository.On("FindBIzCardByUid", mock.Anything, mock.Anything).Return(ets.BizCard, nil)
	ets.BizCardRepository.On("UpdateBizCard", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update error"))
	result := ets.BizCardService.UpdateBizCard(1, &domain.BizCardUpdate{
		Age: 26,
	}, ctx)
	ets.Nil(result)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_DeleteBizCard() {
	ctx, span := otel.Tracer("Delete BizCard").Start(context.Background(), "Delete BizCard Application")
	defer span.End()
	ets.BizCardRepository.On("DeleteBizCardByUid", mock.AnythingOfType("int"), mock.Anything).Return(nil)
	result := ets.BizCardService.DeleteBizCard(1, ctx)
	ets.Equal("success", result)
}

func (ets *BizCardServiceTestSuite) TestBizCardServiceImpl_DeleteBizCard_error() {
	ctx, span := otel.Tracer("Delete BizCard").Start(context.Background(), "Delete BizCard Application")
	defer span.End()
	ets.BizCardRepository.On("DeleteBizCardByUid", mock.AnythingOfType("int"), mock.Anything).Return(errors.New("fail"))
	result := ets.BizCardService.DeleteBizCard(1, ctx)
	ets.Equal("fail", result)
}

func TestBizCardServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BizCardServiceTestSuite))
}

func TestBizCardUpdate(t *testing.T) {
	data := &ent.BizCard{
		Email:       "tae2089",
		Name:        "taebin",
		PhoneNumber: "010-xxxx-xxxx",
		Age:         25,
	}
	bizCard := domain.CreateBizCardUpdate(data)
	bizCard.Update(&domain.BizCardUpdate{
		Name: "tester",
		Age:  100,
	})
	assert.Equal(t, 100, bizCard.Age)
}
