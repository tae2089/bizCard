package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"log"
	"sync"
)

var userRepositoryOnce sync.Once

//go:generate mockery --name UserRepository --case underscore --inpackage
type UserRepository interface {
	RegisterUser(userRegister domain.UserRegister) (*ent.User, error)
	FindUser(email string) (*ent.User, error)
}

func SetupUserRepository() UserRepository {
	if UserRepositoryBean == nil {
		userRepositoryOnce.Do(func() {
			UserRepositoryBean = &UserRepositoryImpl{Client: Client.User}
		})
	}
	log.Println(Client.User == nil)
	return UserRepositoryBean
}
