package main

import (
	"fmt"
	"madyasantosa/ruangkegiatan/config"
	"madyasantosa/ruangkegiatan/pkg/database"

	"github.com/labstack/echo/v4"
)

func main() {
	config := *config.InitConfig()
	e := echo.New()
	db := database.InitDB(config)
	database.Migrate(db)

	fmt.Println(config)
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "PORT: "+config.AppPort)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppPort)))
}
