package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AccomConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	Timezone string
	SslMode  string
	AppPort  string
	AppMode  string
}

func InitConfig() *AccomConfig {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Printf("WARNING cannot load .env file: %v", err)
	// }
	// uncomment if deploy
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Printf("WARNING cannot load .env file: %v", err)
	}

	appMode := os.Getenv("APP_MODE")
	if appMode == "" {
		log.Fatalf("APP_MODE not set in .env")
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatalf("DB_HOST not set in .env")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatalf("DB_USER not set in .env")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatalf("DB_PASSWORD not set in .env")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatalf("DB_NAME not set in .env")
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatalf("DB_PORT not set in .env")
	}
	timezone := os.Getenv("TIMEZONE")
	if timezone == "" {
		log.Fatalf("TIMEZONE not set in .env")
	}
	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		log.Fatalf("SSL_MODE not set in .env")
	}

	appPort := os.Getenv("PORT")
	if appPort == "" {
		log.Fatalf("PORT not set in .env")
	}

	return &AccomConfig{
		Host:     host,
		User:     user,
		Password: password,
		DbName:   dbName,
		Port:     port,
		Timezone: timezone,
		SslMode:  sslMode,
		AppPort:  appPort,
		AppMode:  appMode,
	}
}
