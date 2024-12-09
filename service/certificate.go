package service

import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)


func (s *PortfolioSrv) CreateCertificate(model *model.Certificate) (*uint, error) {
	id, err := s.srv.Create(model)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return id, nil
}

func (s *PortfolioSrv) GetCertificateById(userId *uint) ([]model.Certificate, error) {
	data, err := s.srv.GetCertificateById(userId)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return data, nil
}
