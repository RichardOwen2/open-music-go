package service

import (
	"context"
	"openmusic-api/model/web"
)

type SongService interface {
	Create(ctx context.Context, request web.SongCreateRequest) (web.SongCreateResponse, error)
	Update(ctx context.Context, request web.SongUpdateRequest) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.SongDataResponse, error)
	FindAll(ctx context.Context) ([]web.SongDataResponse, error)
}
