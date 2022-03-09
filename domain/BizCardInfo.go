package domain

type BizCardInfo struct {
	Name        string `form:"name" json:"name" `
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber"`
	Email       string `form:"email" json:"email"  `
	Age         int    `form:"age"  json:"age" `
}

func CreateBizCardInfo(dto BizCard) BizCardInfo {
	return BizCardInfo{
		Name:        dto.Name(),
		PhoneNumber: dto.PhoneNumber(),
		Email:       dto.Email(),
		Age:         dto.Age(),
	}
}
