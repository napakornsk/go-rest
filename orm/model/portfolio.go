package model

import (
	"time"

	"gorm.io/gorm"
)

type GetByUserId struct {
	UserId uint `json:"user_id"`
}

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

type WorkExperience struct {
	gorm.Model       `json:"omitempty"`
	CompanyName      string            `gorm:"size:255" json:"company_name"`
	Role             string            `gorm:"size:255" json:"role"`
	IsDone           bool              `json:"is_done"`
	StartDate        *time.Time        `json:"start_date"`
	EndDate          *time.Time        `json:"end_date"`
	Status           string            `gorm:"size:1" json:"status"`
	UserId           uint              `json:"user_id"`
	WorkId           uint              `json:"work_id"`
	WorkDescriptions []WorkDescription `gorm:"foreignKey:WorkExperienceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"work_descriptions"`
}

type WorkDescription struct {
	gorm.Model       `json:"omitempty"`
	WorkExperienceID uint   `json:"work_experience_id"`
	Description      string `gorm:"size:255" json:"description"`
}

func (WorkExperience) TableName() string {
	return "work_experience"
}

func (WorkDescription) TableName() string {
	return "work_description"
}

type Skill struct {
	gorm.Model
	Name              string              `gorm:"size:255" json:"name"`
	UserId            uint                `json:"user_id"`
	SkillId           uint                `json:"skill_id"`
	SkillDescriptions []SkillDescriptions `gorm:"foreignKey:SkillId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"skill_descriptions"`
}

type SkillDescriptions struct {
	gorm.Model
	SkillId     uint   `json:"skill_id"`
	Description string `gorm:"size:255" json:"description"`
}

func (Skill) TableName() string {
	return "skill"
}

func (SkillDescriptions) TableName() string {
	return "skill_descriptions"
}
