package database

import (
	"github.com/rajesh4295/user-service-go/env"
	"github.com/rajesh4295/user-service-go/model"
)

type Provider interface {
	Connect(e env.Provider) error
	CreateUser(u *model.User) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	Signup(u *model.Signup) (*model.User, error)
	CreateOrg(o *model.Org) (*model.Org, error)
}
