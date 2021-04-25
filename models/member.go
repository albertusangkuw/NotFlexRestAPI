package models

import "time"

type Member struct {
	User
	TanggalLahir time.Time `form:"TanggalLahir" json:"TanggalLahir"`
	JenisKelamin string    `form:"JenisKelamin" json:"JenisKelamin"`
	AsalNegara   string    `form:"AsalNegara" json:"AsalNegara"`
	Status       bool      `form:"Password" json:"Password"`
}

type MemberResponse struct {
	Response
	Data []Member `form:"Data" json:"Data"`
}
