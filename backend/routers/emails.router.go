package routers

import (
	"backend/controllers"

	"github.com/go-chi/chi/v5"
)

func EmailsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/all", controllers.GetEmails)
	r.Get("/filters", controllers.SearchEmailsWithFilters)
	r.Get("/index", controllers.SearchEmailByID)

	return r
}
