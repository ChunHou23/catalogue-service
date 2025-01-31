package base

import "github.com/go-chi/chi/v5"

type Controller interface {
	SetupRoutes(router *chi.Mux)
}
