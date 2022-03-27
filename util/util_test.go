package util_test

import (
	"bizCard/domain"
	"bizCard/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwt(t *testing.T) {
	info := domain.UserInfo{Name: "test", Email: "test@example.com"}
	token, err := util.CreateJwt(info, 1)
	println(token)
	if err != nil {
		panic(err)
	}
	id := util.ParseJwt(token)
	println(id)
	assert.Equal(t, 1, id)
}
