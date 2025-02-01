package mapper

import (
	"github.com/ChunHou23/catalogue-service/internal/models"
	"github.com/ChunHou23/catalogue-service/internal/models/dto"
)

func MapProductEntityToDTO(productEntity models.Product) dto.ProductResponseDto {
	return dto.ProductResponseDto{
		ID:              productEntity.ID,
		Name:            productEntity.Name,
		Description:     productEntity.Description,
		Price:           productEntity.Price,
		Available:       productEntity.Available,
		ProductCategory: productEntity.ProductCategory,
	}
}

func MapListOfProductEntitiesToDTOs(productEntities []models.Product) []dto.ProductResponseDto {
	var productDTOs []dto.ProductResponseDto
	for _, product := range productEntities {
		productDTOs = append(productDTOs, MapProductEntityToDTO(product))
	}
	return productDTOs
}

func CreateNewProductEntity(createProductRequestDto dto.CreateProductRequestDto) models.Product {
	return models.Product{
		Name:            createProductRequestDto.Name,
		Description:     createProductRequestDto.Description,
		Price:           createProductRequestDto.Price,
		Available:       createProductRequestDto.Available,
		ProductCategory: createProductRequestDto.ProductCategory,
		BusinessId:      createProductRequestDto.BusinessId,
	}
}

func UpdateProductEntity(product *models.Product, updateProductRequestDto dto.UpdateProductRequestDto) models.Product {
	product.Available = updateProductRequestDto.Available
	product.Description = updateProductRequestDto.Description
	product.Name = updateProductRequestDto.Name
	product.Price = updateProductRequestDto.Price
	product.ProductCategory = updateProductRequestDto.ProductCategory
	return *product
}
