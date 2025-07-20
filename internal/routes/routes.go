package routes

import (
	"github.com/VincentLAU5142/FEM_project_GO/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HeathCheck)
	return r
}
