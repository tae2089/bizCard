package application

import "main/domain"

type BizCardService interface {
	RegisterBizCard(bizCardDto domain.BizCardDto) domain.BizCardInfo
}
