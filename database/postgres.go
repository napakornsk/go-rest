package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Postgres *gorm.DB
}

func InitPostgres(host, user, password, dbname, port, timezone, sslMode string) (*Database, error) {
	d := &Database{}
	err := d.connectPostgres(host, user, password, dbname, port, timezone, sslMode)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Database) connectPostgres(host, user, password, dbname, port, timezone, sslMode string) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s timezone=%s sslmode=%s", host, user, password, dbname, port, timezone, sslMode)
	// Custom logger configuration
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer with timestamp
		logger.Config{
			SlowThreshold:             time.Second, // Log SQL queries that exceed this duration
			LogLevel:                  logger.Info, // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logging
			Colorful:                  true,        // Enable colored output
		},
	)

	// Open the database with the custom logger
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Printf("Error connecting to PostgreSQL: %v", err)
		return err
	}
	d.Postgres = db
	log.Println("PostgreSQL connection established.")
	return nil
}
