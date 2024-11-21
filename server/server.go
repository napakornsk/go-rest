package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/napakornsk/go-rest/config"
	"github.com/napakornsk/go-rest/database"
	"github.com/napakornsk/go-rest/handler"
	"github.com/napakornsk/go-rest/router"
	"github.com/napakornsk/go-rest/service"
)

func StartRESTServer() {
	c := config.InitConfig()
	g := gin.Default()

	db, err := database.InitPostgres(c.Host, c.User, c.Password, c.DbName, c.Port, c.Timezone, c.SslMode)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// Initialize REST services
	srv := service.InitPortfolioSrv(db)
	h := handler.InitPortfolioHandler(srv)
	r := router.InitPortfolioRouter(h)
	r.SetupRouter(g)

	// Start the server using the specified port from the configuration
	serverAddress := fmt.Sprintf("localhost:%s", c.AppPort)
	log.Printf("Starting REST server on %s...", serverAddress)
	if err := g.Run(serverAddress); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}
