package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique; size:100" json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,password"`
	Nickname string `gorm:"size:255" json:"nickname" validate:"omitempty,alphanum"`
}

func (User) TableName() string {
	return "user"
}
