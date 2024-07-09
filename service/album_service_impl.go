package service

import (
	"context"
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
	if tx.Error != nil {
		return tx.Error
	}

	_, err := s.AlbumRepository.Update(ctx, tx, domain.Album{
		ID:   request.ID,
		Name: request.Name,
		Year: request.Year,
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

func (s *AlbumServiceImpl) Delete(ctx context.Context, categoryId int) error {
	panic("implement me")
}

func (s *AlbumServiceImpl) FindById(ctx context.Context, categoryId int) (web.AlbumDataResponse, error) {
	panic("implement me")
}

func (s *AlbumServiceImpl) FindAll(ctx context.Context) ([]web.AlbumDataResponse, error) {
	panic("implement me")
}
