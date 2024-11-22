package handler

import (
	"github.com/napakornsk/go-rest/service"
)

type PortfolioHandler struct {
	srv *service.PortfolioSrv
}

func InitPortfolioHandler(srv *service.PortfolioSrv) *PortfolioHandler {
	return &PortfolioHandler{
		srv: srv,
	}
}
