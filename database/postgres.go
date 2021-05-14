package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"github.com/rajesh4295/user-service-go/env"
	"github.com/rajesh4295/user-service-go/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *Postgres

type Postgres struct {
	Db *gorm.DB
}

func NewPG() *Postgres {
	if instance != nil {
		fmt.Println("db instance exists")
		return instance
	}
	fmt.Println("new db instance")
	instance = &Postgres{}
	return instance
}

func (pg *Postgres) autoMigrate() {
	pg.Db.AutoMigrate(&model.User{})
	pg.Db.AutoMigrate(&model.Org{})
	fmt.Println("Automigrate complete")
}

func (pg *Postgres) Connect(e env.Provider) error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Kolkata",
		e.Get("DB.HOST"), e.Get("DB.USERNAME"), e.Get("DB.PASSWORD"), e.Get("DB.DATABASE"), e.Get("DB.PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
		return err
	}
	pg.Db = db
	_ = db
	fmt.Println("Connected to db")
	instance.autoMigrate()
	return nil
}

func (pg *Postgres) CreateUser(user *model.User) (*model.User, error) {
	if e := pg.Db.Create(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func (pg *Postgres) GetUserById(id string) (*model.User, error) {
	var user *model.User
	if e := pg.Db.First(&user, "id = ?", id).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func (pg *Postgres) CreateOrg(org *model.Org) (*model.Org, error) {
	if e := pg.Db.Create(&org).Error; e != nil {
		return nil, e
	}

	return org, nil
}

func (pg *Postgres) Signup(u *model.Signup) (*model.User, error) {
	var e error
	tx := pg.Db.Begin()
	// create new user in existing org
	if u.User.OrgID != uuid.FromStringOrNil("") {
		if _, e = pg.CreateUser(&u.User); e != nil {
			tx.Rollback()
			return &u.User, e
		}
	} else {
		// create new user in new org
		if u.OrgName == "" {
			return &u.User, errors.New("org name is required")
		}
		org := &model.Org{Name: u.OrgName}
		if _, e = pg.CreateOrg(org); e != nil {
			tx.Rollback()
			return &u.User, e
		}

		u.User.OrgID = org.Base.ID
		if _, e = pg.CreateUser(&u.User); e != nil {
			tx.Rollback()
			return &u.User, e
		}
	}

	tx.Commit()
	return &u.User, nil
}
