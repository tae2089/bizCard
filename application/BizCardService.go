package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"sync"
)

//go:generate mockery --name BizCardService --case underscore --inpackage
type BizCardService interface {
	RegisterBizCard(bizCardDto *domain.BizCardRegister) *domain.BizCardInfo
	FindBizCard(uid int) *domain.BizCardInfo
	UpdateBizCard(uid int, bizCardUpdate *domain.BizCardUpdate) *domain.BizCardInfo
	DeleteBizCard(uid int) string
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
