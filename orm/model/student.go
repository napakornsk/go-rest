package model

import "gorm.io/gorm"

type CreateStudent struct {
	Student []*Student `json:"student"`
}

type Student struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint   `json:"age"`
	Grade     uint   `json:"grade"`
}

func (Student) TableName() string {
	return "student"
}

type Teacher struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Subject   string `json:"subject"`
	Age       uint   `json:"age"`
}

func (Teacher) TableName() string {
	return "teacher"
}

type Classroom struct {
	gorm.Model
	Name       string    `json:"name"`
	Capacity   uint      `json:"capacity"`
	TeacherID  uint      `json:"teacher_id"`
	Teacher    Teacher   `gorm:"foreignKey:TeacherID"`
	GradeLevel uint      `json:"grade_level"`
	Students   []Student `gorm:"many2many:classroom_students" json:"students"`
}

func (Classroom) TableName() string {
	return "classroom"
}
