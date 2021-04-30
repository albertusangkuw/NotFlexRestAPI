package models

type Film struct {
	IDFilm      int64  `form:"IDFilm" json:"IDFilm"`
	Title       string `form:"Title" json:"Title"`
	Sinopsis    string `form:"Sinopsis" json:"Sinopsis"`
	MainCast string `form:"MainCast" json:"MainCast"`
	Genre       string `form:"Genre" json:"Genre"`
	Director   string `form:"Director" json:"Director"`
	ReleaseYear  int32  `form:"ReleaseYear" json:"ReleaseYear"`
}

type FilmResponse struct {
	Response
	Data []Film `form:"Data" json:"Data"`
}
