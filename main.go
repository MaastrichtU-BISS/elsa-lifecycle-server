package main

import (
	"server/database"
	"server/routes"
)

func main() {

	// Initialize the database connection
	database.ConnectDB()

	// Initialize the router
	r := routes.SetupRouter()

	// Start server
	r.Run(":8080")
}
