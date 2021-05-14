package database

import (
	"github.com/rajesh4295/user-service-go/env"
	"github.com/rajesh4295/user-service-go/model"
)

type Provider interface {
	// app
	Connect(e env.Provider) error
	// user
	CreateUser(u *model.User) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	// org
	CreateOrg(o *model.Org) (*model.Org, error)
	GetOrgById(id string) (*model.Org, error)
	// usecase
	Signup(u *model.Signup) (*model.User, error)
}
