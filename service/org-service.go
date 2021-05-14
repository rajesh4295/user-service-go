package service

import (
	"github.com/rajesh4295/user-service-go/model"
)

type OrgService interface {
	GetOrgById(id string) (*model.Org, error)
}

type orgService struct {
}

// var (
// 	db database.Provider = database.NewPG()
// )

func NewOrgService() OrgService {
	return &orgService{}
}

func (s *orgService) GetOrgById(id string) (*model.Org, error) {
	return Db.GetOrgById(id)
}
