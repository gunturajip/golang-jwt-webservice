package main

import (
	"golang-jwt-auth/database"
	"golang-jwt-auth/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
