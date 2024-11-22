package service

import (
	"github.com/napakornsk/go-rest/database"
)

type PortfolioSrv struct {
	repo *database.Database
}

func InitPortfolioSrv(repo *database.Database) *PortfolioSrv {
	return &PortfolioSrv{
		repo: repo,
	}
}
