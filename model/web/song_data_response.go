package web

type SongDataResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
	Performer string `json:"performer"`
	Duration  int    `json:"duration"`
}

type SongDataResponseWithAlbum struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
	Performer string `json:"performer"`
	Duration  int    `json:"duration"`
	Album     AlbumDataResponse `json:"album"`
}
