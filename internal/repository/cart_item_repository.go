package repository

import (
	"github.com/ChunHou23/catalogue-service/internal/base"
	"github.com/ChunHou23/catalogue-service/internal/models"
)

type CartItemRepository struct {
	*base.BaseRepository[models.CartItem]
}

func NewCartItemRepository() *CartItemRepository {
	return &CartItemRepository{
		base.NewBaseRepository[models.CartItem](),
	}
}
