package models

import (
	"github.com/ChunHou23/catalogue-service/internal/base"
	"github.com/ChunHou23/catalogue-service/internal/models/enums"
)

type Business struct {
	base.Entity
	Name           string                   `gorm:"size:255;not null" json:"name"`
	Description    string                   `gorm:"type:text" json:"description"`
	BusinessNature enums.BusinessNatureType `gorm:"size:255" json:"business_nature"`
	Products       []Product                `gorm:"foreignKey:BusinessId;references:ID" json:"products"`
}

func (business *Business) TableName() string {
	return "business"
}
