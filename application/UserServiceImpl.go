package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"bizCard/util"
	"github.com/rs/zerolog/log"
)

var _ UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (userServiceImpl *UserServiceImpl) RegisterUser(userRegister domain.UserRegister) domain.UserInfo {
	encryptPassword, err := util.GenerateBcrypt(userRegister.Password)
	if err != nil {
		return domain.UserInfo{
			Present: false,
		}
	}
	userRegister.Password = encryptPassword
	savedUser, err := userServiceImpl.UserRepository.RegisterUser(userRegister)
	if err != nil {
		log.Err(err)
		return domain.UserInfo{
			Present: false,
		}
	}
	userInfo := domain.CreateUserInfo(savedUser)
	return userInfo
}
