package service

import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)

func (s *PortfolioSrv) CreateEducation(model *model.Education) (*uint, error) {
	id, err := s.srv.Create(model)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return id, nil
}

func (s *PortfolioSrv) GetEducationById(userId *uint) ([]model.Education, error) {
	id, err := s.srv.GetEducationById(userId)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return id, nil
}
