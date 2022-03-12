package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"log"
)

var _ BizCardService = (*BizCardServiceImpl)(nil)

type BizCardServiceImpl struct {
	BizCardRepository repository.BizCardRepository
}

func (b *BizCardServiceImpl) RegisterBizCard(dto *domain.BizCardRegister) *domain.BizCardInfo {
	bizCard, err := b.BizCardRepository.RegisterBizCard(dto)
	if err != nil {
		log.Println(err)
		return nil
	}
	bizCardInfo := domain.CreateBizCardInfo(bizCard)
	return &bizCardInfo
}

func (b *BizCardServiceImpl) FindBizCard(uid int) *domain.BizCardInfo {
	bizCard, err := b.BizCardRepository.FindBIzCardByUid(uid)
	if err != nil {
		log.Println(err)
		return nil
	}
	bizCardInfo := domain.CreateBizCardInfo(bizCard)
	return &bizCardInfo
}

func (b *BizCardServiceImpl) UpdateBizCard(uid int, bizCardUpdate *domain.BizCardUpdate) *domain.BizCardInfo {
	findBizCard, err := b.BizCardRepository.FindBIzCardByUid(uid)
	if err != nil {
		log.Println("not found bizcard")
		return nil
	}
	updateBizCard, err := b.BizCardRepository.UpdateBizCard(findBizCard, bizCardUpdate)
	if err != nil {
		log.Println(err)
		return nil
	}
	bizCardInfo := domain.CreateBizCardInfo(updateBizCard)
	return &bizCardInfo
}
