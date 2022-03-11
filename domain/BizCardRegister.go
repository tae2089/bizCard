package domain

type BizCardRegister struct {
	Name        string `form:"name" json:"name" binding:"required"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" binding:"required"`
	Email       string `form:"email" json:"email"  binding:"required"`
	Age         int    `form:"age"  json:"age" binding:"required"`
}
