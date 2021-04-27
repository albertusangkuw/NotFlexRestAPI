package models

type Koleksi struct {
	IDKoleksi   int    `form:"IDKoleksi" json:"IDKoleks"`
	NamaKoleksi string `form:"NamaKoleksi" json:"NamaKoleksi"`
}

type KoleksiResponse struct {
	Response
	Data []Koleksi `form:"Data" json:"Data"`
}

type DetailKoleksi struct {
	IDDetailKoleksi int     `form:"IDDetailKoleksi" json:"IDDetailKoleksi"`
	IDKoleksi       Koleksi `form:"IDKoleksi" json:"IDKoleksi"`
	IDFilm          Film    `form:"IDFilm" json:"IDFilm"`
}

type DetailKoleksiResponse struct {
	Response
	Data []DetailKoleksi `form:"data" json:"data"`
}
