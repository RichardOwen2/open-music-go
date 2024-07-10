package web

type AlbumDataResponse struct {
	ID   string    `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

type AlbumDataResponseWithSongs struct {
	ID    int                `json:"id"`
	Name  string             `json:"name"`
	Year  int                `json:"year"`
	Songs []SongDataResponse `json:"songs"`
}
