package web

type AlbumCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Year int    `json:"year" validate:"required"`
}
