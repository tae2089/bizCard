package domain

type BizCardInfo struct {
	name        string
	phoneNumber string
	email       string
	age         int
}

func CreateBizCardInfo(dto BizCard) BizCardInfo {
	return BizCardInfo{
		name:        dto.Name(),
		phoneNumber: dto.PhoneNumber(),
		email:       dto.Email(),
		age:         dto.Age(),
	}
}
