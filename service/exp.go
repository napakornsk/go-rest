package service

import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)

func (s *PortfolioSrv) GetWorkExperience(userId uint) (*model.WorkExperience, error) {
	tx := s.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return nil, tx.Error
	}

	exp := new(model.WorkExperience)
	if err := tx.Where("user_id = ?", userId).Preload("WorkDescriptions").Find(exp).Error; err != nil {
		log.Printf("Failed to fetch intro: %v", err)
		tx.Rollback() // Rollback the transaction on error
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return exp, nil
}

func (s *PortfolioSrv) CreateWorkExperience(model *model.WorkExperience) (*uint, error) {
	tx := s.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return nil, tx.Error
	}

	if err := tx.Omit("id").Create(model).Error; err != nil {
		log.Printf("Failed to create intro: %v", err)
		tx.Rollback() // Rollback the transaction on error
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return &model.ID, nil
}

func (s *PortfolioSrv) UpdateWorkExperience(userId uint, input *model.WorkExperience) (*uint, error) {
	tx := s.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return nil, tx.Error
	}

	// Update existing work experiences' status to "2" for history tracking
	if err := tx.Model(&model.WorkExperience{}).Where("user_id = ? and work_id = ?", userId, input.WorkId).Update("status", "2").Error; err != nil {
		log.Printf("Failed to update status for existing work experience: %v", err)
		tx.Rollback()
		return nil, err
	}

	// Create the new WorkExperience entry
	if err := tx.Omit("id").Create(input).Error; err != nil {
		log.Printf("Failed to create new work experience: %v", err)
		tx.Rollback()
		return nil, err
	}

	for i := range input.WorkDescriptions {
		input.WorkDescriptions[i].WorkExperienceID = input.ID
	}

	if len(input.WorkDescriptions) > 0 {
		if err := tx.Omit("id").Create(&input.WorkDescriptions).Error; err != nil {
			log.Printf("Failed to create work descriptions: %v", err)
			tx.Rollback()
			return nil, err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return &input.ID, nil
}
