package controller

import (
	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/services"
	"github.com/go-chi/chi/v5"
)

type CartItemController struct {
	appConfig       *config.AppConfig
	cartItemService *services.CartItemService
}

func NewCartItemController(appConfig *config.AppConfig) *CartItemController {
	return &CartItemController{
		appConfig:       appConfig,
		cartItemService: services.NewCartItemService(appConfig),
	}
}

func (cartItemController *CartItemController) SetupRoutes(router *chi.Mux) {
	router.Post("/cart-item", cartItemController.cartItemService.AddCartItem)
}
