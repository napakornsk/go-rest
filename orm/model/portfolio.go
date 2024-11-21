package model

import "gorm.io/gorm"

type Intro struct {
	gorm.Model
	FristName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Description string  `gorm:"size:255" json:"description"`
	UserID      uint    `json:"user_id"`
	ContactID   uint    `json:"contact_id"`
	Contact     Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"contact,omitempty"`
}
type Contact struct {
	gorm.Model
	Email        string `gorm:"size:100" json:"email"`
	MobileNo     string `gorm:"size:255" json:"mobile_no"`
	LinkedinLink string `gorm:"size:255" json:"linkedinLink"`
}

func (Contact) TableName() string {
	return "contact"
}

func (Intro) TableName() string {
	return "intro"
}
