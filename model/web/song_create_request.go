package web

type SongCreateRequest struct {
	Title  string `json:"title" validate:"required"`
	Year   int    `json:"year" validate:"required"`
	Genre  string `json:"genre" validate:"required"`
	Performer string `json:"performer" validate:"required"`
	Duration int `json:"duration"`
	AlbumID string `json:"album_id"`
}
