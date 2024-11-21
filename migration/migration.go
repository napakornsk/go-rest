package main

import (
	"log"

	"github.com/napakornsk/go-rest/config"
	"github.com/napakornsk/go-rest/database"
	"github.com/napakornsk/go-rest/orm/entity"
	"gorm.io/gorm"
)

func main() {
	c := config.InitConfig()
	db, err := database.InitPostgres(
		c.AppMode,
		c.Host, c.User,
		c.Password, c.DbName,
		c.Port, c.Timezone,
		c.SslMode,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	migration(db.Postgres)
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Intro{},
		&entity.Contact{},
	); err != nil {
		return err
	}
	return nil
}
