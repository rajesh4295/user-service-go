package service

import (
	"github.com/rajesh4295/user-service-go/database"
	"github.com/rajesh4295/user-service-go/model"
)

type UserService interface {
	Signup(s *model.Signup) (*model.User, error)
	GetUserById(id string) (*model.User, error)
}

type service struct {
}

var (
	db database.Provider = database.NewPG()
)

func NewUserService() UserService {
	return &service{}
}

func (s *service) Signup(u *model.Signup) (*model.User, error) {
	return db.Signup(u)
}

func (s *service) GetUserById(id string) (*model.User, error) {
	return db.GetUserById(id)
}
