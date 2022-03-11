package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"sync"
)

//go:generate mockery --name BizCardService --case underscore --inpackage
type BizCardService interface {
	RegisterBizCard(bizCardDto *domain.BizCardRegister) *domain.BizCardInfo
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
