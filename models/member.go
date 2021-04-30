package models

import "time"

type Member struct {
	User
	DateOfBirth time.Time `form:"DateOfBirth" json:"DateOfBirth"`
	Gender string    `form:"Gender" json:"Gender"`
	OriginCountry   string    `form:"OriginCountry" json:"OriginCountry"`
	Status       bool      `form:"Password" json:"Password"`
}

type MemberResponse struct {
	Response
	Data []Member `form:"Data" json:"Data"`
}
