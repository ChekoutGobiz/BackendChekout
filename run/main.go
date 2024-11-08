// main.go
package main

import (
	"log"
	"net/http"

	"github.com/ChekoutGobiz/ChekoutGobiz-Backend/config"
	"github.com/ChekoutGobiz/ChekoutGobiz-Backend/routes"
)

func main() {
	// Initialize MongoDB connection
	config.ConnectDB()

	// Setup routes
	router := routes.SetupRoutes()

	// Start the server
	log.Println("Server berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
