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

type AlbumServiceImpl struct {
	AlbumRepository repository.AlbumRepository
	db              *gorm.DB
	Validate        *validator.Validate
}

func NewAlbumServiceImpl(albumRepository repository.AlbumRepository, db *gorm.DB, validate *validator.Validate) AlbumService {
	return &AlbumServiceImpl{
		AlbumRepository: albumRepository,
		db:              db,
		Validate:        validate,
	}
}

func (s *AlbumServiceImpl) Create(ctx context.Context, request web.AlbumCreateRequest) (web.AlbumCreateResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		return web.AlbumCreateResponse{}, err
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return web.AlbumCreateResponse{}, tx.Error
	}

	createdAlbum, err := s.AlbumRepository.Create(ctx, tx, domain.Album{
		Name: request.Name,
		Year: request.Year,
	})

	if err != nil {
		tx.Rollback()
		return web.AlbumCreateResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return web.AlbumCreateResponse{}, err
	}

	return web.AlbumCreateResponse{
		AlbumID: createdAlbum.ID,
	}, nil
}

func (s *AlbumServiceImpl) Update(ctx context.Context, request web.AlbumUpdateRequest) error {
	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	tx := s.db.Begin()

	exist, err := s.AlbumRepository.Exist(ctx, tx, request.ID)

	if err := helper.ErrorIfNotExist(fmt.Sprintf("album with id %s not found", request.ID), exist, err); err != nil {
		tx.Rollback()
		return err
	}

	_, err = s.AlbumRepository.Update(ctx, tx, domain.Album{
		ID:   request.ID,
		Name: request.Name,
		Year: request.Year,
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

func (s *AlbumServiceImpl) Delete(ctx context.Context, id string) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	exist, err := s.AlbumRepository.Exist(ctx, tx, id)

	if err := helper.ErrorIfNotExist(fmt.Sprintf("album with id %s not found", id), exist, err); err != nil {
		return err
	}

	err = s.AlbumRepository.Delete(ctx, tx, domain.Album{
		ID: id,
	})

	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *AlbumServiceImpl) FindById(ctx context.Context, id string) (web.AlbumDataResponse, error) {
	album, err := s.AlbumRepository.FindById(ctx, s.db, id)

	if err != nil {
		return web.AlbumDataResponse{}, err
	}

	return web.AlbumDataResponse{
		ID:   album.ID,
		Name: album.Name,
		Year: album.Year,
	}, nil
}

func (s *AlbumServiceImpl) FindAll(ctx context.Context) ([]web.AlbumDataResponse, error) {
	albums, err := s.AlbumRepository.FindAll(ctx, s.db)

	if err != nil {
		return []web.AlbumDataResponse{}, err
	}

	var albumResponses []web.AlbumDataResponse
	for _, album := range albums {
		albumResponses = append(albumResponses, web.AlbumDataResponse{
			ID:   album.ID,
			Name: album.Name,
			Year: album.Year,
		})
	}

	return albumResponses, nil
}
