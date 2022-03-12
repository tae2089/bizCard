package repository_test

import (
	"bizCard/ent/bizcard"
	"bizCard/ent/enttest"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBizCardRepositoryImpl_RegisterBizCard(t *testing.T) {
	client := enttest.Open(t, "mysql", "root:secret@tcp(localhost:13306)/bizcardtest?parseTime=true")
	client.Schema.Create(context.Background())
	data, err := client.BizCard.Create().
		SetAge(15).
		SetName("taebin").
		SetEmail("tae2089").
		SetPhoneNumber("010-xxxx-xxxx").
		Save(context.Background())
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 15, data.Age)
}
func TestBizCardRepositoryImpl_FindBIzCardByUid(t *testing.T) {
	client := enttest.Open(t, "mysql", "root:secret@tcp(localhost:13306)/bizcardtest?parseTime=true")
	client.Schema.Create(context.Background())
	repo, err := client.BizCard.Query().Where(bizcard.ID(1)).First(context.Background())
	if err != nil {
		log.Println(err)
	}
	assert.Equal(t, "taebin", repo.Name)
}
