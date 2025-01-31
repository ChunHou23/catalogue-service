package dto

type CreateProductRequestDto struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	ProductCategory string  `json:"product_category"`
	Available       bool    `json:"available"`
}

type UpdateProductRequestDto struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	ProductCategory string  `json:"product_category"`
	Available       bool    `json:"available"`
}
