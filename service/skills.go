package service

import (
	"log"

	"github.com/napakornsk/go-rest/orm/model"
)

func (s *PortfolioSrv) CreateSkill(model *model.Skill) (*uint, error) {
	id, err := s.srv.Create(model)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return id, nil
}

func (s *PortfolioSrv) GetSkillById(userId *uint) ([]model.Skill, error) {
	data, err := s.srv.GetSkillById(userId)
	if err != nil {
		log.Printf("error while calling repository: %v\n", err)
		return nil, err
	}
	return data, nil
}
