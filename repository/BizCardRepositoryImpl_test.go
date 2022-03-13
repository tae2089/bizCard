package repository_test

import (
	"bizCard/ent"
	"bizCard/ent/bizcard"
	"bizCard/ent/enttest"
	"context"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type BizCardRepositoryTestSuite struct {
	suite.Suite
	Client *ent.Client
}

func (ets *BizCardRepositoryTestSuite) SetupTest() {
	ets.Client = enttest.Open(ets.T(), "mysql", "root:secret@tcp(localhost:13306)/bizcardtest?parseTime=true")
}
func (ets *BizCardRepositoryTestSuite) TestBizCardRepositoryImpl_1_RegisterBizCard() {
	data, err := ets.Client.BizCard.Create().
		SetAge(15).
		SetName("taebin").
		SetEmail("tae2089").
		SetPhoneNumber("010-xxxx-xxxx").
		Save(context.Background())
	if err != nil {
		panic(err)
	}
	ets.Equal(15, data.Age)
}
func (ets *BizCardRepositoryTestSuite) TestBizCardRepositoryImpl_2_FindBIzCardByUid() {
	repo, err := ets.Client.BizCard.Query().Where(bizcard.ID(1)).First(context.Background())
	if err != nil {
		log.Panic(err)
	}
	ets.Equal("taebin", repo.Name)
}

func (ets *BizCardRepositoryTestSuite) TestBizCardRepositoryImpl_3_UpdateBIzCardByUid() {
	repo, err := ets.Client.BizCard.UpdateOneID(1).
		SetAge(16).
		SetEmail("test2089").
		SetName("taebint").
		Save(context.Background())
	if err != nil {
		log.Panic(err)
	}
	ets.Equal(16, repo.Age)
	ets.Equal("test2089", repo.Email)
	ets.Equal("taebint", repo.Name)
}

func (ets *BizCardRepositoryTestSuite) TestBizCardRepositoryImpl_4_DeleteBIzCardByUid() {
	err := ets.Client.BizCard.DeleteOneID(1).Exec(context.Background())
	ets.Nil(err)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BizCardRepositoryTestSuite))
}
