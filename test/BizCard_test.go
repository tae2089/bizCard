package test

import (
	"github.com/stretchr/testify/assert"
	domain2 "main/domain"
	"testing"
)

func TestCreateBizCard(t *testing.T) {
	bizCardDto := domain2.BizCardDto{Email: "tae2089", Name: "taebin", PhoneNumber: "010-xxxx-xxxx", Age: 12}
	bizCard := domain2.CreateBizCard(bizCardDto)
	assert.Equal(t, "taebin", bizCard.Name(), "not same name")
	assert.Equal(t, "tae2089", bizCard.Email())
	assert.Equal(t, 12, bizCard.Age())
}
