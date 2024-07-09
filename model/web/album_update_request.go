package web

type AlbumUpdateRequest struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Year int    `json:"year" validate:"required"`
}
