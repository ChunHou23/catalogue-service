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

type BusinessService struct {
	appConfig          *config.AppConfig
	businessRepository *repository.BusinessRepository
}

func NewBusinessService(appConfig *config.AppConfig) *BusinessService {
	return &BusinessService{
		appConfig:          appConfig,
		businessRepository: repository.NewBusinessRepository(),
	}
}

func (businessService *BusinessService) GetAllBusinesses(w http.ResponseWriter, r *http.Request) {
	businesses, err := businessService.businessRepository.FindAll("Products")
	if err != nil {
		businessService.appConfig.ErrorLog.Printf("Error retrieving businesses: %v", err)
		http.Error(w, "Error retrieving business", http.StatusInternalServerError)
		return
	}

	if len(businesses) == 0 {
		businessService.appConfig.InfoLog.Println("No businesses found")
		return
	}

	var businessDtos []dto.BusinessResponseDto
	for _, business := range businesses {
		businessService.appConfig.InfoLog.Printf("Business: %v", business)
		businessDtos = append(businessDtos, mapper.MapBusinessEntityToDTO(business))
	}

	businessService.returnResponse(w, businessDtos)
}

func (businessService *BusinessService) GetBusinessById(w http.ResponseWriter, r *http.Request) {
	businessId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		businessService.appConfig.ErrorLog.Printf("Error parsing id: %v", err)
		http.Error(w, "Error parsing id", http.StatusBadRequest)
		return
	}

	business, err := businessService.businessRepository.FindById(businessId)
	if err != nil {
		businessService.appConfig.ErrorLog.Printf("Error retrieving business: %v", err)
		return
	}

	businessService.appConfig.InfoLog.Printf("Business: %v", business)
}

func (businessService *BusinessService) returnResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding products to JSON: %v", err)
		http.Error(w, "Error encoding products", http.StatusInternalServerError)
	}
}
