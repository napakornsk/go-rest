package service

import (
	"github.com/napakornsk/go-rest/database"
	"github.com/napakornsk/go-rest/repository"
)

type PortfolioSrv struct {
	srv  *repository.PortfolioRepository
	repo *database.Database
}

func InitPortfolioSrv(srv *repository.PortfolioRepository,
	repo *database.Database) *PortfolioSrv {
	return &PortfolioSrv{
		srv:  srv,
		repo: repo,
	}
}
