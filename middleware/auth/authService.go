package auth


import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)

func (s *AuthService) FindUser(username string) (*model.User, error) {
	tx := s.repo.Postgres.Begin()
	if tx.Error != nil {
		log.Printf("Failed to start transaction: %v\n", tx.Error)
		return nil, tx.Error
	}
	user := new(model.User)
	if err := tx.Where("username = ?", username).First(user).Error; err != nil {
		log.Printf("Failed to find user: %v\n", err)
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v\n", err)
		return nil, err
	}
	return user, nil
}
