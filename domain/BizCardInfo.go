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
func (b *BizCardInfo) Name() string {
	return b.name
}

func (b *BizCardInfo) PhoneNumber() string {
	return b.phoneNumber
}

func (b *BizCardInfo) Email() string {
	return b.email
}

func (b *BizCardInfo) Age() int {
	return b.age
}
