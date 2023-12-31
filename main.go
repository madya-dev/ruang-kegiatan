package main

import (
	"madyasantosa/ruangkegiatan/config"
	"madyasantosa/ruangkegiatan/pkg/database"
	"madyasantosa/ruangkegiatan/pkg/mongodb"
	"madyasantosa/ruangkegiatan/pkg/s3"
	"madyasantosa/ruangkegiatan/server"

	"github.com/go-playground/validator"
)

func main() {
	config := *config.InitConfig()

	mongo := mongodb.InitMongodb(config)

	validate := validator.New()

	db := database.InitDB(config)
	database.Migrate(db)

	s3.NewUploader(config)

	server.InitServer(config, db, validate, mongo)
}
