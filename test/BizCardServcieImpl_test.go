package test

import (
	"bizCard/application"
	domain "bizCard/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBizCardServiceImpl_RegisterBizCard(t *testing.T) {
	bizCardDto := domain.BizCardRegister{
		Email:       "tae2089",
		Name:        "taebin",
		PhoneNumber: "010-xxxx-xxxx",
		Age:         25,
	}
	a := application.BizCardServiceImpl{}
	result := a.RegisterBizCard(bizCardDto)
	assert.Equal(t, "tae2089", result)
}
