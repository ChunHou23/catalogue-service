package models

import (
	"github.com/ChunHou23/catalogue-service/internal/base"
	"github.com/google/uuid"
)

type Product struct {
	base.Entity
	Name            string    `gorm:"size:255;not null" json:"name"`
	Description     string    `gorm:"type:text" json:"description"`
	Price           float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Available       bool      `gorm:"default:true" json:"available"`
	BusinessId      uuid.UUID `gorm:"type:uuid" json:"business_id"`
	ProductCategory string    `gorm:"size:255" json:"product_category"`
}

func (product *Product) TableName() string {
	return "products"
}
