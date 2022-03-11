package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"sync"
)

type BizCardService interface {
	RegisterBizCard(bizCardDto domain.BizCardRegister) *domain.BizCardInfo
}

var onceBizCardService sync.Once

func SetupBizCardService() BizCardService {
	if BizCardServiceBean == nil {
		onceBizCardService.Do(func() {
			BizCardServiceBean = &BizCardServiceImpl{
				bizCardRepository: repository.BizCardRepositoryBean,
			}
		})
	}
	return BizCardServiceBean
}
