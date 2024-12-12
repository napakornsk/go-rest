package auth


import (
	"github.com/napakornsk/go-rest/database"
)

type AuthService struct {
	repo *database.Database
}

func InitAuthService(repo *database.Database) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
