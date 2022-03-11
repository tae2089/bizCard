package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"sync"
)

var bizCardRepositoryOnce sync.Once

type BizCardRepository interface {
	RegisterBizCard(dto domain.BizCardRegister) (*ent.BizCard, error)
}

func SetupBizCardRepository() BizCardRepository {
	if BizCardRepositoryBean == nil {
		bizCardRepositoryOnce.Do(func() {
			BizCardRepositoryBean = &BizCardRepositoryImpl{Client: Client.BizCard}
		})
	}
	return BizCardRepositoryBean
}
