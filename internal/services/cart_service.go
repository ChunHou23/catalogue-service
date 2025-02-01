package services

import (
	"encoding/json"
	"net/http"

	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/mapper"
	"github.com/ChunHou23/catalogue-service/internal/models/dto"
	"github.com/ChunHou23/catalogue-service/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CartService struct {
	appConfig      *config.AppConfig
	cartRepository *repository.CartRepository
}

func NewCartService(appConfig *config.AppConfig) *CartService {
	return &CartService{
		appConfig:      appConfig,
		cartRepository: repository.NewCartRepository(),
	}
}

func (cs *CartService) CreateCart(w http.ResponseWriter, r *http.Request) {
	var createCartRequestDto dto.CreateCartRequestDto
	err := json.NewDecoder(r.Body).Decode(&createCartRequestDto)
	if err != nil {
		cs.appConfig.ErrorLog.Printf("Error decoding request: %v", err)
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	if err := cs.cartRepository.Create(mapper.CreateNewCartEntity(createCartRequestDto)); err != nil {
		cs.appConfig.ErrorLog.Printf("Error creating cart: %v", err)
		http.Error(w, "Error creating cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (cs *CartService) GetCartDetailsById(w http.ResponseWriter, r *http.Request) {
	cartId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		cs.appConfig.ErrorLog.Printf("Error parsing id: %v", err)
		http.Error(w, "Error parsing id", http.StatusBadRequest)
		return
	}

	cart, err := cs.cartRepository.FindById(cartId, "CartItem", "CartItem.Product")
	cs.appConfig.InfoLog.Printf("Cart: %v", cart.CartItem)
	if err != nil {
		cs.appConfig.ErrorLog.Printf("Error getting cart: %v", err)
		http.Error(w, "Error getting cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(mapper.MapCartEntityToDetailsDTO(cart)); err != nil {
		cs.appConfig.ErrorLog.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
