package service

import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)

func (s *PortfolioSrv) GetIntroById(userId *uint) (*model.Intro, error) {
	data, err := s.srv.GetIntroById(userId)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return data, nil
}

func (s *PortfolioSrv) GetAllIntro() ([]*model.Intro, error) {
	tx := s.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v", tx.Error)
		return nil, tx.Error
	}

	var intro []*model.Intro

	if err := tx.Find(&intro).Error; err != nil {
		log.Printf("Failed to fetch intro: %v", err)
		tx.Rollback() // Rollback the transaction on error
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return intro, nil
}

func (s *PortfolioSrv) CreateIntro(model *model.Intro) (*uint, error) {
	id, err := s.srv.Create(model)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return id, nil
}
