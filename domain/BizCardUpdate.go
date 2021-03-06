package domain

import (
	"bizCard/ent"
	"reflect"
)

type BizCardUpdate struct {
	Name        string `form:"name" json:"name" binding:"required"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" binding:"required"`
	Email       string `form:"email" json:"email"  binding:"required"`
	Age         int    `form:"age"  json:"age" binding:"required"`
}

func CreateBizCardUpdate(dto *ent.BizCard) *BizCardUpdate {
	return &BizCardUpdate{
		Name:        dto.Name,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
		Age:         dto.Age,
	}
}

func (b *BizCardUpdate) Update(dto *BizCardUpdate) {
	e := reflect.ValueOf(dto).Elem()
	for i := 0; i < e.NumField(); i++ {
		dtoName := e.Type().Field(i).Name
		if reflect.ValueOf(dto).Elem().FieldByName(dtoName).Interface() != nil && reflect.ValueOf(dto).Elem().FieldByName(dtoName).Interface() != "" && reflect.ValueOf(dto).Elem().FieldByName(dtoName).Interface() != 0 {
			reflect.ValueOf(b).Elem().FieldByName(dtoName).Set(reflect.ValueOf(dto).Elem().FieldByName(dtoName))
		}
	}
}
