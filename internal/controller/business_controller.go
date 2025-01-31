package controller

import (
	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/services"
	"github.com/go-chi/chi/v5"
)

type BusinessController struct {
	appConfig       *config.AppConfig
	businessService *services.BusinessService
}

func NewBusinessController(appConfig *config.AppConfig) *BusinessController {
	return &BusinessController{
		appConfig:       appConfig,
		businessService: services.NewBusinessService(appConfig),
	}
}

func (businessController *BusinessController) SetupRoutes(router *chi.Mux) {
	router.Get("/business", businessController.businessService.GetAllBusinesses)
	router.Get("/business/{id}", businessController.businessService.GetBusinessById)
	// router.Get("/business/type/{type}", businessController.businessService.GetBusinessByType)
}
