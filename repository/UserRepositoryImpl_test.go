package repository

import (
	"bizCard/ent"
	"bizCard/ent/enttest"
	"bizCard/util"
	"context"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	Client *ent.UserClient
}

func (ets *UserRepositoryTestSuite) SetupTest() {
	ets.Client = enttest.Open(ets.T(), "mysql", "root:secret@tcp(localhost:13306)/bizcardtest?parseTime=true").User
}
func (ets *UserRepositoryTestSuite) TestUserRepositoryImpl_1_RegisterUser() {
	password, err := util.GenerateBcrypt("hello01")
	if err != nil {
		panic(err)
	}
	log.Println(">>>")
	savedUser, err := ets.Client.Create().
		SetName("taeaet").
		SetEmail("test").
		SetPassword(password).
		SetCreatedDate(time.Now()).
		SetModifiedDate(time.Now()).
		Save(context.Background())
	ets.Equal(savedUser.Name, "taeaet")
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
