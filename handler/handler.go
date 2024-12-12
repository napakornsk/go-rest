package handler

import (
	customevalidator "github.com/napakornsk/go-rest/customValidator"
	"github.com/napakornsk/go-rest/service"
)

type PortfolioHandler struct {
	srv *service.PortfolioSrv
	v   *customevalidator.Validator
}

func InitPortfolioHandler(srv *service.PortfolioSrv, v *customevalidator.Validator) *PortfolioHandler {
	return &PortfolioHandler{
		srv: srv,
		v:   v,
	}
}
