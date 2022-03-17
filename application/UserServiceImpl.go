package application

import (
	"bizCard/domain"
	"bizCard/repository"
	"bizCard/util"
	"log"
)

var _ UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (userServiceImpl *UserServiceImpl) LoginUser(loginForm domain.UserLoginForm) (domain.UserInfo, int) {
	findUser, err := userServiceImpl.UserRepository.FindUser(loginForm.Email)
	if err != nil {
		log.Println(err)
		return domain.UserInfo{Present: false}, 0
	}
	comparePassword, err := util.Compare(findUser.Password, loginForm.Password)
	if err != nil || comparePassword == false {
		log.Println(err)
		return domain.UserInfo{Present: false}, 0
	}
	return domain.CreateUserInfo(&findUser), findUser.ID
}

func (userServiceImpl *UserServiceImpl) RegisterUser(userRegister domain.UserRegister) domain.UserInfo {
	_, err := userServiceImpl.UserRepository.FindUser(userRegister.Email)
	if err == nil {
		return domain.UserInfo{
			Present: false,
		}
	}

	encryptPassword, err := util.GenerateBcrypt(userRegister.Password)
	if err != nil {
		log.Println(err)
		return domain.UserInfo{
			Present: false,
		}
	}
	userRegister.Password = encryptPassword
	savedUser, err := userServiceImpl.UserRepository.RegisterUser(userRegister)
	if err != nil {
		log.Println(err)
		return domain.UserInfo{
			Present: false,
		}
	}
	userInfo := domain.CreateUserInfo(savedUser)
	return userInfo
}
