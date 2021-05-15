package service

import (
	"github.com/rajesh4295/user-service-go/database"
	"github.com/rajesh4295/user-service-go/model"
)

type UserService interface {
	Signup(s *model.Signup) (*model.User, error)
	Login(l *model.Login) (*model.User, error)
	GetUserById(id string) (*model.User, error)
}

type userService struct {
}

var (
	Db database.Provider = database.NewPG()
)

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Signup(u *model.Signup) (*model.User, error) {
	return Db.Signup(u)
}

func (s *userService) Login(u *model.Login) (*model.User, error) {
	return Db.Login(u)
}

func (s *userService) GetUserById(id string) (*model.User, error) {
	return Db.GetUserById(id)
}
