package repository

import (
	"github.com/ChunHou23/catalogue-service/internal/base"
	"github.com/ChunHou23/catalogue-service/internal/models"
)

type CartRepository struct {
	*base.BaseRepository[models.Cart]
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		base.NewBaseRepository[models.Cart](),
	}
}
