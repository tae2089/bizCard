package repository

import (
	"bizCard/domain"
	"bizCard/ent"
	"bizCard/ent/user"
	"context"
	"time"
)

var _ UserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
	Client *ent.UserClient
}

func (u *UserRepositoryImpl) RegisterUser(userRegister domain.UserRegister) (*ent.User, error) {
	savedUser, err := u.Client.Create().
		SetName(userRegister.Name).
		SetEmail(userRegister.Email).
		SetPassword(userRegister.Password).
		SetCreatedDate(time.Now()).
		SetModifiedDate(time.Now()).
		Save(context.Background())
	return savedUser, err
}

func (u *UserRepositoryImpl) FindUser(email string) (ent.User, error) {
	findUser, err := u.Client.Query().Where(user.Email(email)).Only(context.Background())
	return *findUser, err
}
