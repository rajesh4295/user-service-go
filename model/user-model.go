package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	CreatedAt int64     `gorm:"autoUpdateTime:milli" json:"createdAt"`
	UpdatedAt int64     `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}

type User struct {
	Name     string    `json:"name"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"-"`
	OrgID    uuid.UUID `json:"orgId"`
	Org      Org       `json:"-"`
	Base
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := uuid.NewV4()
	u.Base.ID = id
	// u.Password = app.HashPwd([]byte(u.Password))
	return
}

type Signup struct {
	User    User
	OrgName string `json:"orgName"`
}
