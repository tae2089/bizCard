package domain

type UserLoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
