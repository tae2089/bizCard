package application

import "main/domain"

type BizCardService interface {
	RegisterBizCard(bizCardDto domain.BizCardRegister) domain.BizCardInfo
}
