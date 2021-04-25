package models

type Koleksi struct {
	IDKoleksi   int    `form:"IDKoleksi" json:"IDKoleks"`
	NamaKoleksi string `form:"NamaKoleksi" json:"NamaKoleksi"`
}

type KoleksiResponse struct {
	Response
	Data []Koleksi `form:"Data" json:"Data"`
}
