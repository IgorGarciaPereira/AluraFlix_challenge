package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"aluraFlixAPI/database/models"
	"aluraFlixAPI/utils"
)

var DB *gorm.DB
var err error

func Connect() *gorm.DB {
	DSN := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	if DB == nil {
		DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

		if err != nil {
			log.Fatalln(err)
		}
	}
	return DB
}

func AutoMigrations(DB *gorm.DB) {
	DB.AutoMigrate(&models.Video{})
	DB.AutoMigrate(&models.Category{})
}

func LoadSeeds(DB *gorm.DB){
	seed := utils.GetContentFromFile("database/seed.sql")
	DB.Exec(seed)
}
