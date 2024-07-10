package service

import (
	"context"
	"fmt"
	"openmusic-api/helper"
	"openmusic-api/model/domain"
	"openmusic-api/model/web"
	"openmusic-api/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SongServiceImpl struct {
	SongRepository repository.SongRepository
	db             *gorm.DB
	Validate       *validator.Validate
}

func NewSongServiceImpl(songRepository repository.SongRepository, db *gorm.DB, validate *validator.Validate) SongService {
	return &SongServiceImpl{
		SongRepository: songRepository,
		db:             db,
		Validate:       validate,
	}
}

func (s *SongServiceImpl) Create(ctx context.Context, request web.SongCreateRequest) (web.SongCreateResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		return web.SongCreateResponse{}, err
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return web.SongCreateResponse{}, tx.Error
	}

	createdSong, err := s.SongRepository.Create(ctx, tx, domain.Song{
		Title:   request.Title,
		Year:    request.Year,
		Genre:   request.Genre,
		Performer: request.Performer,
		Duration: request.Duration,
		AlbumID:  &request.AlbumID,
	})

	if err != nil {
		tx.Rollback()
		return web.SongCreateResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return web.SongCreateResponse{}, err
	}

	return web.SongCreateResponse{
		SongID: createdSong.ID,
	}, nil
}

func (s *SongServiceImpl) Update(ctx context.Context, request web.SongUpdateRequest) error {
	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	exist, err := s.SongRepository.Exist(ctx, tx, request.ID)

	if err := helper.ErrorIfNotExist(fmt.Sprintf("song with id %s not found", request.ID), exist, err); err != nil {
		tx.Rollback()
		return err;
	}

	_, err = s.SongRepository.Update(ctx, tx, domain.Song{
		ID:        request.ID,
		Title:     request.Title,
		Year:      request.Year,
		Genre:     request.Genre,
		Performer: request.Performer,
		Duration:  request.Duration,
		AlbumID:   &request.AlbumID,
	})

	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *SongServiceImpl) Delete(ctx context.Context, id string) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	exist, err := s.SongRepository.Exist(ctx, tx, id)

	if err := helper.ErrorIfNotExist(fmt.Sprintf("song with id %s not found", id), exist, err); err != nil {
		tx.Rollback()
		return err;
	}

	err = s.SongRepository.Delete(ctx, tx, domain.Song{
		ID: id,
	})

	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *SongServiceImpl) FindById(ctx context.Context, id string) (web.SongDataResponse, error) {
	song, err := s.SongRepository.FindById(ctx, s.db, id)

	if err != nil {
		return web.SongDataResponse{}, err
	}

	return web.SongDataResponse{
		ID:        song.ID,
		Title:     song.Title,
		Year:      song.Year,
		Genre:     song.Genre,
		Performer: song.Performer,
		Duration:  song.Duration,
	}, nil
}

func (s *SongServiceImpl) FindAll(ctx context.Context) ([]web.SongDataResponse, error) {
	songs, err := s.SongRepository.FindAll(ctx, s.db)

	if err != nil {
		return []web.SongDataResponse{}, err
	}

	var songResponses []web.SongDataResponse
	for _, song := range songs {
		songResponses = append(songResponses, web.SongDataResponse{
			ID:        song.ID,
			Title:     song.Title,
			Year:      song.Year,
			Genre:     song.Genre,
			Performer: song.Performer,
			Duration:  song.Duration,
		})
	}

	return songResponses, nil
}
