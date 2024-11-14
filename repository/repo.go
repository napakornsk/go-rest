package repository

import "gorm.io/gorm"

type StudentRepository interface {
	Begin() *gorm.DB
	Find(out interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Last(out interface{}) *gorm.DB
}