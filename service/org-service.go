package service

import (
	"github.com/rajesh4295/user-service-go/model"
)

type OrgService interface {
	GetOrgById(id string) (*model.Org, error)
}

/*
 *	Org service layer to help interaction between org controller and databse.
**/
type orgService struct {
}

func NewOrgService() OrgService {
	return &orgService{}
}

func (s *orgService) GetOrgById(id string) (*model.Org, error) {
	return Db.GetOrgById(id)
}
