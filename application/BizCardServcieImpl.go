package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"log"
)

var _ BizCardService = (*BizCardServiceImpl)(nil)

type BizCardServiceImpl struct {
	bizCardRepository repository.BizCardRepository
}

func (b *BizCardServiceImpl) RegisterBizCard(dto domain.BizCardRegister) *domain.BizCardInfo {
	bizCard, err := b.bizCardRepository.RegisterBizCard(dto)
	if err != nil {
		log.Println(err)
		return nil
	}
	bizCardInfo := domain.CreateBizCardInfo(bizCard)
	return &bizCardInfo
}
