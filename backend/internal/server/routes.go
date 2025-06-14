package server

import (
	"hospital-portal/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux){
	r.Post("/signup",handlers.Signup)
	r.Post("/login",handlers.Login)
	r.Route("/patients", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)
    	r.Post("/add", handlers.RegisterPatient)
    	r.Get("/", handlers.GetAllPatients)
    	r.Put("/{id}", handlers.UpdatePatient)
    	r.Delete("/{id}", handlers.DeletePatient)
	})
	r.Post("/logout", handlers.Logout)

}