package main

import (
	"log"
	"syncspend/config"
	"syncspend/routes"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Setup router
	router := routes.SetupRouter()

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
