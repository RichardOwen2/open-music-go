package repository

import (
	"context"
	"fmt"
	"openmusic-api/helper"
	"openmusic-api/model/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SongRepositoryImpl struct {
}

func NewSongRepositoryImpl() SongRepository {
	return &SongRepositoryImpl{}
}

func (r *SongRepositoryImpl) checkAlbumExists(ctx context.Context, db *gorm.DB, id string) (bool, error) {
	var total int64
	if err := db.WithContext(ctx).Model(&domain.Album{}).Where("id = ?", id).Count(&total).Error; err != nil {
		return false, err
	}

	return total > 0, nil
}

func (r *SongRepositoryImpl) Exist(ctx context.Context, db *gorm.DB, id string) (bool, error) {
	var total int64
	if err := db.WithContext(ctx).Model(&domain.Song{}).Where("id = ?", id).Count(&total).Error; err != nil {
		return false, err
	}

	return total > 0, nil
}

func (r *SongRepositoryImpl) Create(ctx context.Context, db *gorm.DB, song domain.Song) (domain.Song, error) {
	id, err := helper.GenerateId("song")
	song.ID = id

	if err != nil {
		return song, err
	}

	if *song.AlbumID != "" {
		exist, err := r.checkAlbumExists(ctx, db, *song.AlbumID)
		if err != nil {
			return song, err
		}

		if !exist {
			return song, fiber.NewError(fiber.StatusNotFound, "album not found")
		}
	} else {
		song.AlbumID = nil
	}

	if err := db.WithContext(ctx).Create(&song).Error; err != nil {
		return song, err
	}

	return song, nil
}

func (r *SongRepositoryImpl) Update(ctx context.Context, db *gorm.DB, song domain.Song) (domain.Song, error) {
	if *song.AlbumID != "" {
		exist, err := r.checkAlbumExists(ctx, db, *song.AlbumID)
		if err != nil {
			return song, err
		}

		if !exist {
			return song, fiber.NewError(fiber.StatusNotFound, "album not found")
		}
	} else {
		song.AlbumID = nil
	}

	if err := db.WithContext(ctx).Save(&song).Error; err != nil {
		return song, err
	}

	return song, nil
}

func (r *SongRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, song domain.Song) error {
	result := db.WithContext(ctx).Delete(&song)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fiber.NewError(fiber.StatusNotFound, "song not found")
	}

	return nil
}

func (r *SongRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, id string) (domain.Song, error) {
	var song domain.Song
	if err := db.WithContext(ctx).Where("id = ?", id).First(&song).Error; err != nil {
		return song, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("song with id %s not found", id))
	}

	return song, nil
}

func (r *SongRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Song, error) {
	var songs []domain.Song
	if err := db.WithContext(ctx).Find(&songs).Error; err != nil {
		return songs, err
	}

	return songs, nil
}
