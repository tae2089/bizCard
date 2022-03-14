package domain

type UserRegister struct {
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
}
