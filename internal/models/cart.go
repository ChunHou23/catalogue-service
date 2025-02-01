package models

import "github.com/ChunHou23/catalogue-service/internal/base"

type Cart struct {
	base.Entity
	Name       string     `gorm:"size:255;not null" json:"name"`
	TotalPrice float64    `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Active     bool       `gorm:"default:true" json:"active"`
	CartItem   []CartItem `gorm:"foreignKey:CartId;references:ID" json:"cart_item"`
}

func (cart *Cart) TableName() string {
	return "cart"
}
