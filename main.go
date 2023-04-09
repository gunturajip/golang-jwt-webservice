package main

import (
	"golang-jwt-auth/database"
	"golang-jwt-auth/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run("127.0.0.1:8080")
}
