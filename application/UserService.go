package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"sync"
)

var userServiceOnce sync.Once

//go:generate mockery --name UserService --case underscore --inpackage
type UserService interface {
	RegisterUser(userRegister domain.UserRegister) domain.UserInfo
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
