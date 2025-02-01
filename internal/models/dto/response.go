package dto

import (
	"github.com/ChunHou23/catalogue-service/internal/models/enums"
	"github.com/google/uuid"
)

type ProductResponseDto struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	Available       bool      `json:"available"`
	ProductCategory string    `json:"product_category"`
}

type ProductDetailResponseDto struct{}

type ProductCreateResponseDto struct{}

type BusinessResponseDto struct {
	ID             uuid.UUID                `json:"id"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description"`
	BusinessNature enums.BusinessNatureType `json:"business_nature"`
	Products       []ProductResponseDto
}

type CartResponseDto struct {
	ID         uuid.UUID             `json:"id"`
	Name       string                `json:"name"`
	TotalPrice float64               `json:"total_price"`
	Active     bool                  `json:"active"`
	CartItem   []CartItemResponseDto `json:"cart_item"`
}

type CartItemResponseDto struct {
	ID       uuid.UUID          `json:"id"`
	Name     string             `json:"name"`
	Quantity int                `json:"quantity"`
	Product  ProductResponseDto `json:"product"`
}
