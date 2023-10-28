package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort         string
	JWTSecret       string
	DBHost          string
	DBPort          string
	DBName          string
	DBUsername      string
	DBPassword      string
	S3SAccessKey    string
	S3SecretKey     string
	S3Region        string
	FirebaseAuthKey string
}

var config *Config

func InitConfig() *Config {
	godotenv.Load()

	config = &Config{
		AppPort:         os.Getenv("APP_PORT"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBName:          os.Getenv("DB_NAME"),
		DBUsername:      os.Getenv("DB_USERNAME"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		S3SAccessKey:    os.Getenv("S3_ACCESS_KEY_ID"),
		S3SecretKey:     os.Getenv("S3_ACCESS_KEY_SECRET"),
		S3Region:        os.Getenv("S3_REGION"),
		FirebaseAuthKey: os.Getenv("FIREBASE_AUTH_KEY"),
	}

	if config.AppPort == "" {
		log.Panic("[Error] App Port cant be empty")
	}

	if config.JWTSecret == "" {
		log.Panic("[Error] JWT Secret cant be empty")
	}
	if config.DBHost == "" {
		log.Panic("[Error] DB Host cant be empty")
	}
	if config.DBPort == "" {
		log.Panic("[Error] DB Port cant be empty")
	}
	if config.DBName == "" {
		log.Panic("[Error] DB Name cant be empty")
	}
	if config.DBUsername == "" {
		log.Panic("[Error] DB Username cant be empty")
	}
	if config.DBPassword == "" {
		log.Panic("[Error] DB Password cant be empty")
	}

	return config
}
