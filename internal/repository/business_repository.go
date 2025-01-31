package repository

import (
	"github.com/ChunHou23/catalogue-service/internal/base"
	"github.com/ChunHou23/catalogue-service/internal/models"
)

type BusinessRepository struct {
	*base.BaseRepository[models.Business]
}

func NewBusinessRepository() *BusinessRepository {
	return &BusinessRepository{
		base.NewBaseRepository[models.Business](),
	}
}
