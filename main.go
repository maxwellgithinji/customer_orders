package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/maxwellgithinji/customer_orders/auth"
	_ "github.com/maxwellgithinji/customer_orders/docs"
	"github.com/maxwellgithinji/customer_orders/routes"
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

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl http://localhost:8080/api/v1/login
func main() {
	// Initialize dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Init Session
	err = auth.InitSession()
	if err != nil {
		log.Fatal("Error initializing session")
	}

	// port := os.Getenv("PORT")
	http.Handle("/", routes.RouteHandlers())
	log.Printf("listening on http://%s/", "127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}