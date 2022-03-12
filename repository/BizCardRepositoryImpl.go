package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"bizCard/ent/bizcard"
	"context"
)

type BizCardRepositoryImpl struct {
	Client *ent.BizCardClient
}

func (b *BizCardRepositoryImpl) RegisterBizCard(dto *domain.BizCardRegister) (*ent.BizCard, error) {
	savedBizCard, err := b.Client.
		Create().
		SetAge(dto.Age).
		SetName(dto.Name).
		SetEmail(dto.Email).
		SetPhoneNumber(dto.PhoneNumber).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return savedBizCard, nil
}

func (b *BizCardRepositoryImpl) FindBIzCardByUid(uid int) (*ent.BizCard, error) {
	bizCard, err := b.Client.Query().Where(bizcard.ID(uid)).First(context.Background())
	return bizCard, err
}

func (b *BizCardRepositoryImpl) UpdateBizCard(findBizCard *ent.BizCard, dto *domain.BizCardUpdate) (*ent.BizCard, error) {
	bizCardUpdate := domain.CreateBizCardUpdate(findBizCard)
	bizCardUpdate = bizCardUpdate.Update(dto)
	bizCard, err := b.Client.UpdateOneID(findBizCard.ID).
		SetAge(bizCardUpdate.Age).
		SetEmail(bizCardUpdate.Email).
		SetName(bizCardUpdate.Name).
		SetPhoneNumber(bizCardUpdate.PhoneNumber).
		Save(context.Background())
	return bizCard, err
}
