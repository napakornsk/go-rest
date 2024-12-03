package repository

import (
	"log"

	"github.com/napakornsk/go-rest/util"
)

func (r *PortfolioRepository) Create(model interface{}) (*uint, error) {
	tx := r.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v\n", tx.Error)
		return nil, tx.Error
	}

	if err := tx.Omit("id").Create(model).Error; err != nil {
		log.Printf("Failed to create model: %v\n", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v\n", err)
		return nil, err
	}

	successId := util.GetID(model)

	return &successId, nil
}
