package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/napakornsk/go-rest/config"
	customevalidator "github.com/napakornsk/go-rest/customValidator"
	"github.com/napakornsk/go-rest/database"
	"github.com/napakornsk/go-rest/handler"
	"github.com/napakornsk/go-rest/repository"
	"github.com/napakornsk/go-rest/router"
	"github.com/napakornsk/go-rest/service"
)

func StartRESTServer() {
	c := config.InitConfig()
	g := gin.Default()

	db, err := database.InitPostgres(c.AppMode, c.Host, c.User, c.Password, c.DbName, c.Port, c.Timezone, c.SslMode)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// Initialize REST services
	repo := repository.InitRepository(db)
	// authSrv := middleware.InitAuthService(c.JWTSecret)
	srv := service.InitPortfolioSrv(repo, db)
	v := validator.New()
	myValidator := customevalidator.InitValidator(v)
	v.RegisterValidation("password", myValidator.Password)

	h := handler.InitPortfolioHandler(srv, myValidator)
	r := router.InitPortfolioRouter(h)

	// userId := uint(1)
	// token, err := authSrv.GenerateToken(userId)
	// if err != nil {
	// fmt.Printf("Failed to generate token: %v\n", err)
	// return
	// }

	// fmt.Printf("Generated Token: %s\n", token)
	//
	protected := g.Group("/")
	// protected.Use(middleware.AuthMiddleware(authSrv))
	r.SetupProtectedRouter(protected)

	r.SetupRouter(g)

	// Start the server using the specified port from the configuration
	serverAddress := fmt.Sprintf(":%s", c.AppPort)
	log.Printf("Starting REST server on %s...", serverAddress)
	if err := g.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}
