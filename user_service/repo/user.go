package repo

import (
	"fmt"
	"github.com/micro_game/user_service/model"
)


type Repository interface {
	Create(user *model.User) error
	GetAll() ([]*model.User, error)
}

type UserRepository struct {
}

func (repo *UserRepository) Create(user *model.User) error {
	fmt.Println("Create")
	return nil
}

func (repo *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	return users, nil
}
