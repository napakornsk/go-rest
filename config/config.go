package config

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AccomConfig struct {
	Host      string
	User      string
	Password  string
	DbName    string
	Port      string
	Timezone  string
	SslMode   string
	AppPort   string
	AppMode   string
	JWTSecret string
}

var PrvKey *ecdsa.PrivateKey
var PbcKey *ecdsa.PublicKey

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

	jwtSecret := os.Getenv("JWTSECRET")
	if appPort == "" {
		log.Fatalf("JWTSECRET not set in .env")
	}

	prvKey, err := loadECDSAPrivateKey("/etc/secrets/private.pem")
	if err != nil {
		log.Fatalf("Cannot load private key")
	}
	PrvKey = prvKey

	PbcKey, err = loadECDSAPublicKey("/etc/secrets/public.pem")
	if err != nil {
		log.Fatalf("Cannot load public key")
	}

	return &AccomConfig{
		Host:      host,
		User:      user,
		Password:  password,
		DbName:    dbName,
		Port:      port,
		Timezone:  timezone,
		SslMode:   sslMode,
		AppPort:   appPort,
		AppMode:   appMode,
		JWTSecret: jwtSecret,
	}
}

func loadECDSAPrivateKey(filename string) (*ecdsa.PrivateKey, error) {
	// Read the PEM file containing the private key
	keyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode the PEM file to extract the private key
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	// Parse the private key into ECDSA format
	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

func loadECDSAPublicKey(filename string) (*ecdsa.PublicKey, error) {
	// Read the PEM file containing the private key
	keyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode the PEM file to extract the private key
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	// Parse the private key into ECDSA format
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pbcKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not of type *ecdsa.PublicKey")
	}

	return pbcKey, nil
}
