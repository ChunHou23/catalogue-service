package models

import (
	"github.com/ChunHou23/catalogue-service/internal/base"
	"github.com/google/uuid"
)

type CartItem struct {
	base.Entity
	ProductId uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	CartId    uuid.UUID `gorm:"type:uuid;not null" json:"cart_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Deleted   bool      `gorm:"default:false" json:"deleted"`
	Product   Product   `gorm:"foreignKey:ID;references:ProductId" json:"product"`
}

func (cartItem *CartItem) TableName() string {
	return "cart_item"
}
