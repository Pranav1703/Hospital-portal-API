package server

import (
	"hospital-portal/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux){
	r.Post("/signup",handlers.Signup)
	r.Post("/login",handlers.Login)
}