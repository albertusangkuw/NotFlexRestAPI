package models

import "time"

type Langganan struct {
	IDLangganan              string    `form:"IDLangganan" json:"IDLangganan"`
	IDUser                   string    `form:"IDUser" json:"IDUser"`
	StatusLangganan          bool      `form:"StatusLangganan" json:"StatusLangganan"`
	TanggalTransaksi         time.Time `form:"TanggalTransaksi" json:"TanggalTransaksi"`
	JenisLangganan           string    `form:"JenisLangganan" json:"JenisLangganan"`
	AwalTanggalBerlangganan  time.Time `form:"AwalTanggalBerlangganan" json:"AwalTanggalBerlangganan"`
	AkhirTanggalBerlangganan time.Time `form:"AkhirTanggalBerlangganan" json:"AkhirTanggalBerlangganan"`
	Token                    string    `form:"Token" json:"Token"`
}

type LanggananResponse struct {
	Response
	Data []Langganan `form:"Data" json:"Data"`
}
