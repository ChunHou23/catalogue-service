package services

import (
	"encoding/json"
	"net/http"

	"github.com/ChunHou23/catalogue-service/internal/config"
	"github.com/ChunHou23/catalogue-service/internal/mapper"
	"github.com/ChunHou23/catalogue-service/internal/models/dto"
	"github.com/ChunHou23/catalogue-service/internal/repository"
)

type CartItemService struct {
	appConfig          *config.AppConfig
	cartItemRepository *repository.CartItemRepository
}

func NewCartItemService(appConfig *config.AppConfig) *CartItemService {
	return &CartItemService{
		appConfig:          appConfig,
		cartItemRepository: repository.NewCartItemRepository(),
	}
}

func (cis *CartItemService) AddCartItem(w http.ResponseWriter, r *http.Request) {
	var addCartItemRequestDto dto.AddCartItemRequestDto
	err := json.NewDecoder(r.Body).Decode(&addCartItemRequestDto)
	if err != nil {
		cis.appConfig.ErrorLog.Printf("Error decoding request: %v", err)
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	if err := cis.cartItemRepository.Create(mapper.CreateCartItemEntity(addCartItemRequestDto)); err != nil {
		cis.appConfig.ErrorLog.Printf("Error creating cart item: %v", err)
		http.Error(w, "Error creating cart item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
