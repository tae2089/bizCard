package domain

import (
	"bizCard/ent"
	"time"
)

type UserInfo struct {
	Name         string    `form:"name" json:"name" `
	Email        string    `form:"email" json:"email" `
	CreatedDate  time.Time `form:"createdDate" json:"createdDate" `
	ModifiedDate time.Time `form:"modifiedDate" json:"modifiedDate" `
	Present      bool      `form:"present" json:"present" `
}

func CreateUserInfo(entity *ent.User) UserInfo {
	return UserInfo{
		Name:         entity.Name,
		Email:        entity.Email,
		CreatedDate:  entity.CreatedDate,
		ModifiedDate: entity.ModifiedDate,
		Present:      true,
	}
}
