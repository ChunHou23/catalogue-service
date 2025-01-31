package setup

import (
	"net/http"

	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/controller"
	"github.com/go-chi/chi/v5"
)

func SetupControllers(appConfig *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	productController := controller.NewProductController(appConfig)

	productController.SetupRoutes(router)

	return router
}
