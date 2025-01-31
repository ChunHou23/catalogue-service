package controller

import (
	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/services"
	"github.com/go-chi/chi/v5"
)

type ProductController struct {
	appConfig      *config.AppConfig
	productService *services.ProductService
}

func NewProductController(appConfig *config.AppConfig) *ProductController {
	return &ProductController{
		appConfig:      appConfig,
		productService: services.NewProductService(appConfig),
	}
}

func (productController *ProductController) SetupRoutes(router *chi.Mux) {
	router.Get("/products", productController.productService.GetAllProducts)
	router.Get("/products/{id}", productController.productService.GetProductById)
	router.Get("/products/biz/{businessId}", productController.productService.GetProductByBusinessId)

	router.Post("/products", productController.productService.CreateProduct)
	router.Patch("/products/{id}", productController.productService.UpdateProduct)
}
