package service

import (
	"context"
	"openmusic-api/model/web"
)

type AlbumService interface {
	Create(ctx context.Context, request web.AlbumCreateRequest) (web.AlbumCreateResponse, error)
	Update(ctx context.Context, request web.AlbumUpdateRequest) error
	Delete(ctx context.Context, categoryId int) error
	FindById(ctx context.Context, categoryId int) (web.AlbumDataResponse, error)
	FindAll(ctx context.Context) ([]web.AlbumDataResponse, error)
}
