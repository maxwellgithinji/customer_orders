package routes

import (
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production

	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/customer_orders/middlewares"
	"github.com/maxwellgithinji/customer_orders/utils"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RouteHandlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Handle not found
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		utils.ResponseHelper(w, "404", "Error, Page NOt Found")
	})

	var api = r.PathPrefix("/api").Subrouter()

	// Swagger
	defer r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api.Handle("/favicon.ico", http.NotFoundHandler()).Methods("GET")
	api.Use(middlewares.CommonMiddleware)

	//API V1
	apiV1(api)

	return r
}

func index(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	utils.ResponseHelper(w, "200", "Welcome to customer orders, login to continue")
}
