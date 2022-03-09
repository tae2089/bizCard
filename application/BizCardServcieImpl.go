package application

import (
	"main/domain"
)

var _ BizCardService = (*BizCardServiceImpl)(nil)

type BizCardServiceImpl struct{}

func (b *BizCardServiceImpl) RegisterBizCard(dto domain.BizCardRegister) domain.BizCardInfo {
	bizCard := domain.CreateBizCard(dto)
	//bizCard db saved
	bizCardInfo := domain.CreateBizCardInfo(bizCard)
	return bizCardInfo
}
