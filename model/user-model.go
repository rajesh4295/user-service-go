package model

import (
	"fmt"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	CreatedAt int64     `gorm:"autoUpdateTime:milli" json:"createdAt"`
	UpdatedAt int64     `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}

type User struct {
	Name     string    `gorm:"unique" json:"name"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"password,omitempty"`
	OrgID    uuid.UUID `json:"orgId"`
	Org      Org       `json:"-"`
	Base
}

type Signup struct {
	User    User
	OrgName string `json:"orgName"`
}

type Login struct {
	Name     string
	Email    string
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := uuid.NewV4()
	u.Base.ID = id
	u.Password = HashPwd([]byte(u.Password))
	return
}

func HashPwd(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("failed to hash password: ", err)
	}
	return string(hash)
}

func ComparePwd(hash string, pwd []byte) bool {
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	return err == nil
}
