package main

import (
	"aluraFlixAPI/database"
	"aluraFlixAPI/router"
)

func main() {
	database.Connect()
	database.AutoMigrations(database.DB)
	database.LoadSeeds(database.DB)
	router.HandleRequests()
}
