package model

type Org struct {
	Name string `json:"name"`
	Base Base   `gorm:"embedded"`
}
