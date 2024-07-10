package repository

import (
	"context"
	"openmusic-api/model/domain"

	"gorm.io/gorm"
)

type SongRepository interface {
	Exist(ctx context.Context, db *gorm.DB, albumId string) (bool, error)
	Create(ctx context.Context, db *gorm.DB, song domain.Song) (domain.Song, error)
	Update(ctx context.Context, db *gorm.DB, song domain.Song) (domain.Song, error)
	Delete(ctx context.Context, db *gorm.DB, song domain.Song) error
	FindById(ctx context.Context, db *gorm.DB, id string) (domain.Song, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.Song, error)
}
