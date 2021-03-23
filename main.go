package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/maxwellgithinji/customer_orders/databases"
	_ "github.com/maxwellgithinji/customer_orders/docs"
	"github.com/maxwellgithinji/customer_orders/routes"
)

var (
	database databases.Database = databases.NewDatabase()
)

// @title Client Orders
// @version 1.0.0
// @description this is a service that helps customers order items
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email maxwellgithinji@gmail.com

// @license.name MIT
// @license.url https://github.com/maxwellgithinji/customer_orders/blob/develop/LICENSE
//
// @BasePath /api/v1
func main() {
	// Initialize dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize db connection
	_, err = database.InitializeDbConnection()
	if err != nil {
		log.Fatal("Error connecting to db", err)
	}
	// port := os.Getenv("PORT")
	port := os.Getenv("PORT")
	http.Handle("/", routes.RouteHandlers())
	log.Println("Running on port...", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
