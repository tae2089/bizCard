package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"context"
	"sync"
)

var bizCardRepositoryOnce sync.Once

//go:generate mockery --name BizCardRepository --case underscore --inpackage
type BizCardRepository interface {
	RegisterBizCard(dto *domain.BizCardRegister, ctx context.Context) (*ent.BizCard, error)
	FindBIzCardByUid(uid int, ctx context.Context) (*ent.BizCard, error)
	UpdateBizCard(uid int, bizCardUpdate *domain.BizCardUpdate, ctx context.Context) (*ent.BizCard, error)
	DeleteBizCardByUid(uid int, ctx context.Context) error
}

func SetupBizCardRepository() BizCardRepository {
	if BizCardRepositoryBean == nil {
		bizCardRepositoryOnce.Do(func() {
			BizCardRepositoryBean = &BizCardRepositoryImpl{Client: Client.BizCard}
		})
	}
	return BizCardRepositoryBean
}
