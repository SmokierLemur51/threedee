package routes

import (

	"github.com/go-chi/chi/v5"
	"github.com/SmokierLemur51/greenleafcleaning/v2/handlers"
)


func ConfigureRoutes(router *chi.Mux) {
	router.Get("/", handlers.IndexHandler)
}