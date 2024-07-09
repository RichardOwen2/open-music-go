package repository

import (
	"context"
	"openmusic-api/helper"
	"openmusic-api/model/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AlbumRepositoryImpl struct {
}

func NewAlbumRepositoryImpl() AlbumRepository {
	return &AlbumRepositoryImpl{}
}

func (r *AlbumRepositoryImpl) Create(ctx context.Context, db *gorm.DB, album domain.Album) (domain.Album, error) {
	id, err := helper.GenerateId("album")
	if err != nil {
		return album, err
	}

	album.ID = id

	if err := db.WithContext(ctx).Create(&album).Error; err != nil {
		return album, err
	}

	return album, nil
}

func (r *AlbumRepositoryImpl) Update(ctx context.Context, db *gorm.DB, album domain.Album) (domain.Album, error) {
	if err := db.WithContext(ctx).Save(&album).Error; err != nil {
		return album, err
	}

	return album, nil
}

func (r *AlbumRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, album domain.Album) error {
	result := db.WithContext(ctx).Delete(&album)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fiber.NewError(fiber.StatusNotFound, "album not found")
	}

	return nil
}

func (r *AlbumRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, albumId string) (domain.Album, error) {
	var album domain.Album
	if err := db.WithContext(ctx).Where("id = ?", albumId).First(&album).Error; err != nil {
		return album, err
	}

	return album, nil
}

func (r *AlbumRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Album, error) {
	var albums []domain.Album
	if err := db.WithContext(ctx).Find(&albums).Error; err != nil {
		return albums, err
	}

	return albums, nil
}
