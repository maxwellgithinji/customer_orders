package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/maxwellgithinji/customer_orders/routes"
)

func main() {
	// Initialize dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// port := os.Getenv("PORT")
	http.Handle("/", routes.RouteHandlers())
	log.Printf("listening on http://%s/", "127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
