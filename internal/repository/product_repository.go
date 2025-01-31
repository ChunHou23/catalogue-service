package repository

import (
	"github.com/ChunHou23/catalogue-service/internal/base"
	"github.com/ChunHou23/catalogue-service/internal/models"
	"github.com/google/uuid"
)

type ProductRepository struct {
	*base.BaseRepository[models.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		base.NewBaseRepository[models.Product](),
	}
}

func (productRepository *ProductRepository) GetProductByBusinessId(id uuid.UUID) ([]models.Product, error) {
	var products []models.Product
	if err := productRepository.DB.Where("business_id = ?", id).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
