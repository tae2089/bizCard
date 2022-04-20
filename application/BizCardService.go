package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"context"
	"sync"
)

//go:generate mockery --name BizCardService --case underscore --inpackage
type BizCardService interface {
	RegisterBizCard(bizCardDto *domain.BizCardRegister, ctx context.Context) *domain.BizCardInfo
	FindBizCard(uid int, ctx context.Context) *domain.BizCardInfo
	UpdateBizCard(uid int, bizCardUpdate *domain.BizCardUpdate, ctx context.Context) *domain.BizCardInfo
	DeleteBizCard(uid int, ctx context.Context) string
}

var onceBizCardService sync.Once

func SetupBizCardService() BizCardService {
	if BizCardServiceBean == nil {
		onceBizCardService.Do(func() {
			BizCardServiceBean = &BizCardServiceImpl{
				BizCardRepository: repository.BizCardRepositoryBean,
			}
		})
	}
	return BizCardServiceBean
}
