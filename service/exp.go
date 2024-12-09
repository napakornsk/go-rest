package service

import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)

func (s *PortfolioSrv) GetWorkExperience(userId uint) ([]model.WorkExperience, error) {
	data, err := s.srv.GetWorkExperienceById(&userId)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}

	return data, nil
}

func (s *PortfolioSrv) CreateWorkExperience(model *model.WorkExperience) (*uint, error) {
	id, err := s.srv.Create(model)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}

	return id, nil
}

func (s *PortfolioSrv) UpdateWorkExperience(userId uint, input *model.WorkExperience) (*uint, error) {
	tx := s.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return nil, tx.Error
	}

	if err := tx.Model(&model.WorkExperience{}).Where("user_id = ? and work_id = ?", userId, input.WorkId).Update("status", "2").Error; err != nil {
		log.Printf("Failed to update status for existing work experience: %v", err)
		tx.Rollback()
		return nil, err
	}

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
