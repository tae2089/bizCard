package domain

type BizCard struct {
	name        string
	phoneNumber string
	email       string
	age         int
}

func CreateBizCard(dto BizCardDto) BizCard {
	return BizCard{
		name:        dto.Name,
		phoneNumber: dto.PhoneNumber,
		email:       dto.Email,
		age:         dto.Age,
	}
}

func (b *BizCard) Name() string {
	return b.name
}

func (b *BizCard) PhoneNumber() string {
	return b.phoneNumber
}

func (b *BizCard) Email() string {
	return b.email
}

func (b *BizCard) Age() int {
	return b.age
}
