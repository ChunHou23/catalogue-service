package mapper

import (
	"github.com/ChunHou23/catalogue-service/internal/models"
	"github.com/ChunHou23/catalogue-service/internal/models/dto"
)

func MapBusinessEntityToDTO(business models.Business) dto.BusinessResponseDto {
	return dto.BusinessResponseDto{
		ID:             business.ID,
		Name:           business.Name,
		Description:    business.Description,
		BusinessNature: business.BusinessNature,
		Products:       MapListOfProductEntitiesToDTOs(business.Products),
	}

}

func MapBusinessToBusinessDetailsDTO(business models.Business) dto.BusinessResponseDto {
	return dto.BusinessResponseDto{
		ID:             business.ID,
		Name:           business.Name,
		Description:    business.Description,
		BusinessNature: business.BusinessNature,
	}
}
