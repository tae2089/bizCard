package test

import (
	"bizCard/application"
	"bizCard/domain"
	"bizCard/ent"
	mockrepo "bizCard/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestBizCardServiceImpl_RegisterBizCard(t *testing.T) {
	bizCardDto := domain.BizCardRegister{
		Email:       "tae2089",
		Name:        "taebin",
		PhoneNumber: "010-xxxx-xxxx",
		Age:         25,
	}
	bizCardRepository := &mockrepo.MockBizCardRepository{}
	bizCardService := &application.BizCardServiceImpl{bizCardRepository}
	bizCardRepository.On("RegisterBizCard", mock.Anything).Return(&ent.BizCard{
		Email:       "tae2089",
		Name:        "taebin",
		PhoneNumber: "010-xxxx-xxxx",
		Age:         25,
	}, nil)
	result := bizCardService.RegisterBizCard(&bizCardDto)
	assert.Equal(t, "tae2089", result.Email)
}
