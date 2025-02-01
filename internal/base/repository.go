package base

import (
	"github.com/ChunHou23/catalogue-service/internal/base/driver"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{
		DB: driver.InitDBInstance(),
	}
}

func (r *BaseRepository[T]) FindAll(eagerLoading ...string) ([]T, error) {
	var result []T
	query := r.DB
	for _, preload := range eagerLoading {
		query = query.Preload(preload)
	}

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepository[T]) FindById(id uuid.UUID, eagerLoading ...string) (T, error) {
	var result T
	query := r.DB
	for _, preload := range eagerLoading {
		query = query.Preload(preload)
	}

	if err := query.First(&result, id).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (r *BaseRepository[T]) Create(data T) error {
	if err := r.DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Update(data T) error {
	if err := r.DB.Save(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Delete(data T) error {
	if err := r.DB.Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
