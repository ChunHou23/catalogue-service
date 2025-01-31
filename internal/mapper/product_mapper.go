package mapper

import (
	"github.com/ChunHou23/catalogue-service/internal/models"
	"github.com/ChunHou23/catalogue-service/internal/models/dto"
)

func MapProductToDTO(productEntity models.Product) dto.ProductResponseDto {
	return dto.ProductResponseDto{
		ID:              productEntity.ID,
		Name:            productEntity.Name,
		Description:     productEntity.Description,
		Price:           productEntity.Price,
		Available:       productEntity.Available,
		ProductCategory: productEntity.ProductCategory,
	}
}

func MapDtoToEntity(createProductRequestDto dto.CreateProductRequestDto) models.Product {
	return models.Product{
		Name:            createProductRequestDto.Name,
		Description:     createProductRequestDto.Description,
		Price:           createProductRequestDto.Price,
		Available:       createProductRequestDto.Available,
		ProductCategory: createProductRequestDto.ProductCategory,
	}
}

func MapDtoToEntityForUpdate(product *models.Product, updateProductRequestDto dto.UpdateProductRequestDto) models.Product {
	product.Available = updateProductRequestDto.Available
	product.Description = updateProductRequestDto.Description
	product.Name = updateProductRequestDto.Name
	product.Price = updateProductRequestDto.Price
	product.ProductCategory = updateProductRequestDto.ProductCategory
	return *product
}
