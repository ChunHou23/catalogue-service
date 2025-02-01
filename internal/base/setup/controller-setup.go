package setup

import (
	"net/http"

	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/controller"
	"github.com/go-chi/chi/v5"
)

// todo: switch router to single switch of app config
func SetupControllers(appConfig *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	productController := controller.NewProductController(appConfig)
	businessController := controller.NewBusinessController(appConfig)
	cartController := controller.NewCartController(appConfig)
	cartItemController := controller.NewCartItemController(appConfig)

	productController.SetupRoutes(router)
	businessController.SetupRoutes(router)
	cartController.SetupRoutes(router)
	cartItemController.SetupRoutes(router)

	return router
}
