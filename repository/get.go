package repository

import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)

func (r *PortfolioRepository) GetSkillById(userId *uint) ([]model.Skill, error) {
	tx := r.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v\n", tx.Error)
		return nil, tx.Error
	}
	var data []model.Skill
	if err := tx.Preload("SkillDescriptions").Where("user_id = ?", *userId).Find(&data).Error; err != nil {
		log.Printf("Failed to create model: %v\n", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v\n", err)
		return nil, err
	}

	return data, nil
}

func (r *PortfolioRepository) GetWorkExperienceById(userId *uint) ([]model.WorkExperience, error) {
	tx := r.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v\n", tx.Error)
		return nil, tx.Error
	}
	var data []model.WorkExperience
	if err := tx.Preload("WorkDescriptions").Where("user_id = ?", *userId).Find(&data).Error; err != nil {
		log.Printf("Failed to create model: %v\n", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v\n", err)
		return nil, err
	}

	return data, nil
}

func (r *PortfolioRepository) GetIntroById(userId *uint) (*model.Intro, error) {
	tx := r.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v\n", tx.Error)
		return nil, tx.Error
	}
	var data model.Intro
	if err := tx.Preload("Contact").Where("user_id = ?", *userId).Find(&data).Error; err != nil {
		log.Printf("Failed to create model: %v\n", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v\n", err)
		return nil, err
	}

	return &data, nil
}
