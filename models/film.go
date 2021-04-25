package models

import time "time"

type Film struct {
	IDFilm      int64     `form:"IDFilm" json:"IDFilm"`
	Judul       string    `form:"Judul" json:"Judul"`
	Sinopsis    string    `form:"Sinopsis" json:"Sinopsis"`
	PemainUtama string    `form:"PemainUtama" json:"PemainUtama"`
	Genre       string    `form:"Genre" json:"Genre"`
	Sutradara   string    `form:"Sutradara" json:"Sutradara"`
	TahunRilis  time.Time `form:"TahunRilis" json:"TahunRilis"`
}

type FilmResponse struct {
	Response
	Data []Film `form:"Data" json:"Data"`
}
