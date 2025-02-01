package mapper

import (
	"github.com/ChunHou23/catalogue-service/internal/models"
	"github.com/ChunHou23/catalogue-service/internal/models/dto"
)

func CreateNewCartEntity(createCartRequestDto dto.CreateCartRequestDto) models.Cart {
	return models.Cart{
		Name:       createCartRequestDto.Name,
		Active:     createCartRequestDto.Active,
		TotalPrice: 0.0,
	}
}

func MapCartEntityToDetailsDTO(cart models.Cart) dto.CartResponseDto {
	return dto.CartResponseDto{
		ID:         cart.ID,
		Name:       cart.Name,
		TotalPrice: cart.TotalPrice,
		Active:     cart.Active,
		CartItem:   MapListOfCartItemEntityToDTO(cart.CartItem),
	}
}
