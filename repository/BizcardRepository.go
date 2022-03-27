package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"sync"
)

var bizCardRepositoryOnce sync.Once

//go:generate mockery --name BizCardRepository --case underscore --inpackage
type BizCardRepository interface {
	RegisterBizCard(dto *domain.BizCardRegister) (*ent.BizCard, error)
	FindBIzCardByUid(uid int) (*ent.BizCard, error)
	UpdateBizCard(uid int, bizCardUpdate *domain.BizCardUpdate) (*ent.BizCard, error)
	DeleteBizCardByUid(uid int) error
}

func SetupBizCardRepository() BizCardRepository {
	if BizCardRepositoryBean == nil {
		bizCardRepositoryOnce.Do(func() {
			BizCardRepositoryBean = &BizCardRepositoryImpl{Client: Client.BizCard}
		})
	}
	return BizCardRepositoryBean
}
