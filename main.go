package main

import (
	"log"
	"net/http"

	"github.com/frankwang0/MangoCommerce/app/catalogue"
	_ "github.com/frankwang0/MangoCommerce/config"
	"github.com/frankwang0/MangoCommerce/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/products", routes.RequestHandler(catalogue.CreateCategory)).Methods("POST")
	r.Handle("/products/{id}", routes.RequestHandler(catalogue.GetCategory)).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
