package application

import (
	"main/domain"
)

var _ BizCardService = (*BizCardServiceImpl)(nil)

type BizCardServiceImpl struct{}

func (b *BizCardServiceImpl) RegisterBizCard(dto domain.BizCardDto) domain.BizCardInfo {
	bizCard := domain.CreateBizCard(dto)
	bizCardInfo := domain.CreateBizCardInfo(bizCard)
	return bizCardInfo
}
