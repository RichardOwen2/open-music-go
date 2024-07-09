package repository

import (
	"context"
	"openmusic-api/model/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SongRepositoryImpl struct {
}

func NewSongRepositoryImpl() SongRepository {
	return &SongRepositoryImpl{}
}

func (r *SongRepositoryImpl) checkAlbumExists(ctx context.Context, db *gorm.DB, albumId string) (bool, error) {
	var exists bool

	if err := db.WithContext(ctx).Model(&domain.Album{}).Select("count(*) > 0").Where("id = ?", albumId).Find(&exists).Error; err != nil {
		return false, err
	}

	return exists, nil
}

func (r *SongRepositoryImpl) Create(ctx context.Context, db *gorm.DB, song domain.Song) (domain.Song, error) {
	exist, err := r.checkAlbumExists(ctx, db, *song.AlbumID)

	if err != nil {
		return song, err
	}

	if !exist {
		return song, fiber.NewError(fiber.StatusNotFound, "album not found")
	}

	if err := db.WithContext(ctx).Create(&song).Error; err != nil {
		return song, err
	}

	return song, nil
}

func (r *SongRepositoryImpl) Update(ctx context.Context, db *gorm.DB, song domain.Song) (domain.Song, error) {
	exist, err := r.checkAlbumExists(ctx, db, *song.AlbumID)

	if err != nil {
		return song, err
	}

	if !exist {
		return song, fiber.NewError(fiber.StatusNotFound, "album not found")
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

func (r *SongRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, songId int) (domain.Song, error) {
	var song domain.Song
	if err := db.WithContext(ctx).Where("id = ?", songId).First(&song).Error; err != nil {
		return song, err
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
