package main

import (
	"log"
	"net/http"
	"todo-dani-backend/database"
	"todo-dani-backend/routes"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const PORT string = "8000" // TODO: make .env

func main() {
	db := database.Connection()

	database.Migrate(db)

	db.Close()

	router := chi.NewRouter()

	routes.Routes(router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	log.Println("initializing server in port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, c.Handler(router)))
}
