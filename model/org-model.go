package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Org struct {
	Name string `json:"name"`
	Base
}

func (o *Org) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := uuid.NewV4()
	o.Base.ID = id
	return
}
