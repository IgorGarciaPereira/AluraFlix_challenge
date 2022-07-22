package main

import (
	"aluraFlixAPI/database"
	"aluraFlixAPI/router"
)

func main() {
	database.Connect()
	database.AutoMigrations(database.DB)
	router.HandleRequests()
}
