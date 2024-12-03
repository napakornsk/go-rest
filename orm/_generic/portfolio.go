package generic

import (
	"github.com/napakornsk/go-rest/orm/entity"
	"github.com/napakornsk/go-rest/orm/model"
)

type Portfolio interface {
	entity.Intro | entity.Contact | entity.Skill | entity.SkillDescriptions | entity.WorkDescription | entity.WorkExperience |
		model.Intro | model.Contact | model.Skill | model.SkillDescriptions | model.WorkDescription | model.WorkExperience
}
