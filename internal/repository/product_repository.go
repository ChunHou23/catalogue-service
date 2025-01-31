package repository

import (
	"github.com/ChunHou23/catalogue-service/internal/base"
	"github.com/ChunHou23/catalogue-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{DB: base.InitDBInstance()}
}

func (productRepository *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := productRepository.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (productRepository *ProductRepository) GetProductById(id uuid.UUID) (models.Product, error) {
	var product models.Product
	if err := productRepository.DB.First(&product, id).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return models.Product{}, nil
		}

		return models.Product{}, err
	}
	return product, nil
}

func (productRepository *ProductRepository) GetProductByBusinessId(id uuid.UUID) ([]models.Product, error) {
	var products []models.Product
	if err := productRepository.DB.Where("business_id = ?", id).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (productRepository *ProductRepository) CreateProduct(product models.Product) error {
	if err := productRepository.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (productRepository *ProductRepository) UpdateProduct(product models.Product) error {
	if err := productRepository.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}
