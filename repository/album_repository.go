package repository

import (
	"context"
	"openmusic-api/model/domain"

	"gorm.io/gorm"
)

type AlbumRepository interface {
	Exist(ctx context.Context, db *gorm.DB, id string) (bool, error)
	Create(ctx context.Context, db *gorm.DB, album domain.Album) (domain.Album, error)
	Update(ctx context.Context, db *gorm.DB, album domain.Album) (domain.Album, error)
	Delete(ctx context.Context, db *gorm.DB, album domain.Album) error
	FindById(ctx context.Context, db *gorm.DB, id string) (domain.Album, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.Album, error)
}
