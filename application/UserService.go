package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"context"
	"sync"
)

var userServiceOnce sync.Once

//go:generate mockery --name UserService --case underscore --inpackage
type UserService interface {
	RegisterUser(userRegister domain.UserRegister, ctx context.Context) domain.UserInfo
	LoginUser(loginForm domain.UserLoginForm, ctx context.Context) (domain.UserInfo, int)
}

func SetupUserService() UserService {
	if UserServiceBean == nil {
		userServiceOnce.Do(func() {
			UserServiceBean = &UserServiceImpl{
				UserRepository: repository.UserRepositoryBean,
			}
		})
	}
	return UserServiceBean
}
