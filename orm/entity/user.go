package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique; size:100" json:"username"`
	Password string `json:"password"`
	Nickname string `gorm:"size:255"`
}

func (User) TableName() string {
	return "user"
}
