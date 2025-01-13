package main

import (
	"backend/routers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Mount("/emails", routers.EmailsRouter())

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
