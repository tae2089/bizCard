package repository

import (
	"bizCard/ent"
	"bizCard/ent/enttest"
	"bizCard/ent/user"
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
		SetEmail("test@example.com").
		SetPassword(password).
		SetCreatedDate(time.Now()).
		SetModifiedDate(time.Now()).
		Save(context.Background())
	ets.Equal(savedUser.Name, "taeaet")
}

func (ets *UserRepositoryTestSuite) TestUserRepositoryImpl_2_FindUser() {
	findUser, err := ets.Client.Query().Where(user.Email("test@example.com")).Only(context.Background())
	if err != nil {
		panic(err)
	}
	ets.Equal(findUser.Email, "test@example.com")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
