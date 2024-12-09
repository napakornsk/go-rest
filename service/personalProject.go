package service

import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)

func (s *PortfolioSrv) CreatePersonalProject(model *model.PersonalProject) (*uint, error) {
	id, err := s.srv.Create(model)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return id, nil
}

func (s *PortfolioSrv) GetPersonalProjectById(userId *uint) ([]model.PersonalProject, error) {
	data, err := s.srv.GetPersonalProjectById(userId)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return data, nil
}
