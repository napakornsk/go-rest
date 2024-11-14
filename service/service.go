package service

import (
	"log"

	"github.com/napakornsk/go-rest/database"
	"github.com/napakornsk/go-rest/orm/model"
)

type StudentSrv struct {
	repo *database.Database
}

func InitStudentSrv(repo *database.Database) *StudentSrv {
	return &StudentSrv{
		repo: repo,
	}
}

func (s *StudentSrv) GetAllStudents() ([]*model.Student, error) {
	tx := s.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return nil, tx.Error
	}

	var students []*model.Student

	if err := tx.Find(&students).Error; err != nil {
		log.Printf("Failed to fetch students: %v", err)
		tx.Rollback() // Rollback the transaction on error
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return students, nil
}

func (s *StudentSrv) CreateStudents(studentModel []*model.Student) ([]*uint, error) {
	tx := s.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return nil, tx.Error
	}

	var successID []*uint
	if err := tx.Omit("id").CreateInBatches(studentModel, len(studentModel)).Error; err != nil {
		log.Printf("Failed to create student: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	for _, student := range studentModel {
		successID = append(successID, &student.ID)
	}
	return successID, nil
}
