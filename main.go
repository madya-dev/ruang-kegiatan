package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		// log.Fatalf("Error loading .env file: %v", err)
		log.Println("Error loading .env file: %v", err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		port := os.Getenv("APP_PORT")
		if port == "" {
			return c.String(500, "APP_PORT environment variable not set")
		}
		return c.String(200, "PORT: "+port)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
