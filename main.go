package main

import (
	"log"
	"net/http"

	"github.com/frankwang0/MangoCommerce/app/catalogue"
	"github.com/frankwang0/MangoCommerce/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	r := mux.NewRouter()
	r.Handle("/categories", routes.RequestHandler(catalogue.CreateCategory)).Methods("POST")
	r.Handle("/categories/{id}", routes.RequestHandler(catalogue.GetCategory)).Methods("GET")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func LoadEnv() {
	fileErr := godotenv.Load("config/.env")
	if fileErr != nil {
		log.Fatal("Error loading .env file")
	}
}
