package dto

import "github.com/google/uuid"

type CreateProductRequestDto struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	ProductCategory string    `json:"product_category"`
	Available       bool      `json:"available"`
	BusinessId      uuid.UUID `json:"business_id"`
}

type UpdateProductRequestDto struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	ProductCategory string  `json:"product_category"`
	Available       bool    `json:"available"`
}

type CreateCartRequestDto struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type AddCartItemRequestDto struct {
	Quantity  int       `json:"quantity"`
	ProductId uuid.UUID `json:"product_id"`
	CartId    uuid.UUID `json:"cart_id"`
}
