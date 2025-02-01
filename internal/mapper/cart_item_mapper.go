package mapper

import (
	"github.com/ChunHou23/catalogue-service/internal/models"
	"github.com/ChunHou23/catalogue-service/internal/models/dto"
)

func CreateCartItemEntity(addCartItemRequestDto dto.AddCartItemRequestDto) models.CartItem {
	return models.CartItem{
		Quantity:  addCartItemRequestDto.Quantity,
		ProductId: addCartItemRequestDto.ProductId,
		CartId:    addCartItemRequestDto.CartId,
		Deleted:   false,
	}
}

func MapCartItemEntityToDTO(cartItem models.CartItem) dto.CartItemResponseDto {
	return dto.CartItemResponseDto{
		ID:       cartItem.ID,
		Quantity: cartItem.Quantity,
		Product:  MapProductEntityToDTO(cartItem.Product),
	}
}

func MapListOfCartItemEntityToDTO(cartItems []models.CartItem) []dto.CartItemResponseDto {
	var cartItemDTOs []dto.CartItemResponseDto
	for _, cartItem := range cartItems {
		cartItemDTOs = append(cartItemDTOs, MapCartItemEntityToDTO(cartItem))
	}
	return cartItemDTOs
}
