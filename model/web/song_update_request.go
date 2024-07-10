package web

type SongUpdateRequest struct {
	ID     string `json:"id" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Year   int    `json:"year" validate:"required"`
	Genre  string `json:"genre" validate:"required"`
	Performer string `json:"performer" validate:"required"`
	Duration int `json:"duration"`
	AlbumID string `json:"album_id"`
}
