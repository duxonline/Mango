package router

import (
	"github.com/gorilla/mux"
)

func addAppRoutes(r *mux.Router) {
	r.Handle("/api/v1/categories", handlers.AuthHandler(summary.SubmitRequest)).Methods("PUT")
}

// addBaseRoutes is the base of all our routes
func addBaseRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.IndexHandler).Methods("GET")
}
