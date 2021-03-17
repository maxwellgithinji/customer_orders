package routes

import (
	"net/http"
	_ "net/http/pprof" // For dev only, dont push to production

	"github.com/gorilla/mux"
	"github.com/maxwellgithinji/customer_orders/middlewares"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RouteHandlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	var api = r.PathPrefix("/api").Subrouter()

	// Swagger
	defer r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	api.Handle("/favicon.ico", http.NotFoundHandler()).Methods("GET")
	api.Use(middlewares.CommonMiddleware)

	//API V1
	apiV1(api)

	return r
}
