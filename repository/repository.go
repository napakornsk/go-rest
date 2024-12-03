package repository

import (
	"github.com/napakornsk/go-rest/database"
)

type PortfolioRepository struct {
	repo *database.Database
}

func InitRepository(repo *database.Database) *PortfolioRepository {
	return &PortfolioRepository{
		repo: repo,
	}
}
