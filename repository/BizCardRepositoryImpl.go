package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"bizCard/ent/bizcard"
	"context"
)

var _ BizCardRepository = (*BizCardRepositoryImpl)(nil)

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

func (b *BizCardRepositoryImpl) UpdateBizCard(uid int, bizCardUpdate *domain.BizCardUpdate) (*ent.BizCard, error) {

	bizCard, err := b.Client.UpdateOneID(uid).
		SetAge(bizCardUpdate.Age).
		SetEmail(bizCardUpdate.Email).
		SetName(bizCardUpdate.Name).
		SetPhoneNumber(bizCardUpdate.PhoneNumber).
		Save(context.Background())
	return bizCard, err
}

func (b *BizCardRepositoryImpl) DeleteBizCardByUid(uid int) error {
	err := b.Client.DeleteOneID(uid).Exec(context.Background())
	return err
}
