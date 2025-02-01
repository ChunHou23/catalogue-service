package controller

import (
	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/services"
	"github.com/go-chi/chi/v5"
)

type CartController struct {
	appConfig   *config.AppConfig
	cartService *services.CartService
}

func NewCartController(appConfig *config.AppConfig) *CartController {
	return &CartController{
		appConfig:   appConfig,
		cartService: services.NewCartService(appConfig),
	}
}

func (cartController *CartController) SetupRoutes(router *chi.Mux) {
	router.Get("/cart/{id}", cartController.cartService.GetCartDetailsById)
	router.Post("/cart", cartController.cartService.CreateCart)
}
