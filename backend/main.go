package main

import (
	"backend/routers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	// Configure CORS middleware
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	r.Use(corsOptions.Handler)

	r.Mount("/emails", routers.EmailsRouter())

	log.Println("Server running on port 8080")
	http.ListenAndServe("0.0.0.0:8080", r)
}
