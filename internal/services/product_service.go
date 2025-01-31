package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/mapper"
	"github.com/ChunHou23/catalogue-service/internal/models/dto"
	"github.com/ChunHou23/catalogue-service/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ProductService struct {
	appConfig         *config.AppConfig
	productRepository *repository.ProductRepository
}

func NewProductService(appConfig *config.AppConfig) *ProductService {
	return &ProductService{
		appConfig:         appConfig,
		productRepository: repository.NewProductRepository(),
	}
}

func (productService *ProductService) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := productService.productRepository.FindAll()
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error retrieving products: %v", err)
		http.Error(w, "Error retrieving products", http.StatusInternalServerError)
		return
	}

	if len(products) == 0 {
		productService.appConfig.InfoLog.Println("No products found")
		productService.returnResponse(w, dto.ProductResponseDto{})
		return
	}

	var productDTOs []dto.ProductResponseDto
	for _, product := range products {
		productDTOs = append(productDTOs, mapper.MapProductEntityToDTO(product))
	}

	productService.returnResponse(w, productDTOs)
}

func (productService *ProductService) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error parsing id: %v", err)
		http.Error(w, "Error parsing id", http.StatusBadRequest)
		return
	}

	product, err := productService.productRepository.FindById(productId)
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error retrieving products: %v", err)
		http.Error(w, "Error retrieving products", http.StatusInternalServerError)
		return
	}

	productService.returnResponse(w, mapper.MapProductEntityToDTO(product))
}

func (productService *ProductService) GetProductByBusinessId(w http.ResponseWriter, r *http.Request) {
	businessId, err := uuid.Parse(chi.URLParam(r, "businessId"))
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error parsing businessId: %v", err)
		http.Error(w, "Error parsing businessId", http.StatusBadRequest)
		return
	}

	products, err := productService.productRepository.GetProductByBusinessId(businessId)
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error retrieving products: %v", err)
		http.Error(w, "Error retrieving products", http.StatusInternalServerError)
		return
	}

	var productDTOs []dto.ProductResponseDto
	for _, product := range products {
		productDTOs = append(productDTOs, mapper.MapProductEntityToDTO(product))
	}

	productService.returnResponse(w, productDTOs)
}

func (productService *ProductService) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductRequestDto dto.CreateProductRequestDto
	err := json.NewDecoder(r.Body).Decode(&createProductRequestDto)
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error decoding request: %v", err)
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	productService.appConfig.InfoLog.Printf("Creating product: %v", createProductRequestDto)
	err = productService.productRepository.Create(mapper.CreateNewProductEntity(createProductRequestDto))

	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error creating product: %v", err)
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	productService.returnResponse(w, "Product created successfully")
}

func (productService *ProductService) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error parsing id: %v", err)
		http.Error(w, "Error parsing id", http.StatusBadRequest)
		return
	}

	product, err := productService.productRepository.FindById(productId)
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error retrieving product: %v", err)
		http.Error(w, "Error retrieving product", http.StatusInternalServerError)
		return
	}

	if product.ID == uuid.Nil {
		productService.appConfig.InfoLog.Panic("Product not found")
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	var updateProductRequestDto dto.UpdateProductRequestDto
	err = json.NewDecoder(r.Body).Decode(&updateProductRequestDto)
	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error decoding request: %v", err)
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	productService.appConfig.InfoLog.Printf("Updating product: %v", updateProductRequestDto)
	err = productService.productRepository.Update(mapper.UpdateProductEntity(&product, updateProductRequestDto))

	if err != nil {
		productService.appConfig.ErrorLog.Printf("Error updating product: %v", err)
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	productService.returnResponse(w, "Product updated successfully")
}

func (productService *ProductService) returnResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding products to JSON: %v", err)
		http.Error(w, "Error encoding products", http.StatusInternalServerError)
	}
}
