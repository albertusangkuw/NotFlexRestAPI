package models

import "time"

type Langganan struct {
	IdLangganan              string    `form:"IdLangganan" json:"IdLangganan"`
	TanggalTransaksi         time.Time `form:"TanggalTransaksi" json:"TanggalTransaksi"`
	StatusLangganan          bool      `form:"StatusLangganan" json:"StatusLangganan"`
	JenisLangganan           string    `form:"JenisLangganan" json:"JenisLangganan"`
	AwalTanggalBerlangganan  time.Time `form:"AwalTanggalBerlangganan" json:"AwalTanggalBerlangganan"`
	AkhirTanggalBerlangganan time.Time `form:"AkhirTanggalBerlangganan" json:"AkhirTanggalBerlangganan"`
	Token                    string    `form:"Token" json:"Token"`
}

type LanggananResponse struct {
	Response
	Data []Langganan `form:"Data" json:"Data"`
}
