package entity

import (
	"time"

	"gorm.io/gorm"
)

type Intro struct {
	gorm.Model
	FristName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Description string  `gorm:"size:255" json:"description"`
	Status      string  `gorm:"size:1" json:"status"`
	UserID      uint    `json:"user_id"`
	ContactID   uint    `json:"contact_id"`
	Contact     Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"contact"`
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
	gorm.Model
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
	gorm.Model
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
	Status            string              `gorm:"size:1" json:"status"`
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

type Education struct {
	gorm.Model
	Name        string     `gorm:"size:255" json:"name"`
	Major       string     `gorm:"size:255" json:"major"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	EducationId uint       `json:"education_id"`
	UserId      uint       `json:"user_id"`
	Status      string     `gorm:"size:1" json:"status"`
}

func (Education) TableName() string {
	return "education"
}

type Language struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       string `json:"level"`
	UserId      uint   `json:"user_id"`
	LanguageId  uint   `json:"language_id"`
	Status      string `gorm:"size:1" json:"status"`
}

func (Language) TableName() string {
	return "language"
}

type PersonalProject struct {
	gorm.Model
	Name                  string                  `gorm:"size:255" json:"name"`
	Description           string                  `gorm:"size:255" json:"description"`
	QrCode                string                  `gorm:"size:255" json:"qr_code"`
	GithubURL             string                  `gorm:"size:255" json:"github_url"`
	Status                string                  `gorm:"size:1" json:"status"`
	ProjectId             uint                    `json:"project_id"`
	PersonalProjectDetail []PersonalProjectDetail `gorm:"foreignKey:ProjectId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"personal_project_detail"`
}

func (PersonalProject) TableName() string {
	return "personal_project"
}

type PersonalProjectDetail struct {
	gorm.Model
	Type        string `gorm:"size:1" json:"type"`
	Description string `json:"description"`
	ProjectId   uint   `json:"project_id"`
}

func (PersonalProjectDetail) TableName() string {
	return "personal_project_detail"
}

type Certificate struct {
	gorm.Model
	Name          string     `json:"name"`
	IssuedDate    *time.Time `json:"issued_date"`
	Publisher     string     `json:"publisher"`
	URL           string     `json:"url"`
	CertificateId uint       `json:"certificate_id"`
	Status        string     `gorm:"size:1" json:"status"`
}

func (Certificate) TableName() string {
	return "certificate"
}
