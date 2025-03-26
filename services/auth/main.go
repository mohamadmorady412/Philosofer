package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"auth-service/database"
	"auth-service/handlers"
)

func main() {
	database.InitDB()
	database.InitRedis()

	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	log.Println("Auth Service running on port 8080")
	http.ListenAndServe(":8080", r)
}