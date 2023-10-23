package database

import (
	"fmt"
	"log"
	"madyasantosa/ruangkegiatan/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c config.Config) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUsername,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName)
	fmt.Println(connStr)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect database, ", err.Error())
	}

	return db
}
