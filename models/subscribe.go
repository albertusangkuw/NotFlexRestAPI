package models

import "time"

type Subscribe struct {
	IDSubscribe        string    `form:"IDSubscribe" json:"IDSubscribe"`
	IDUser             string    `form:"IDUser" json:"IDUser"`
	StatusSubscribe    bool      `form:"StatusSubscribe" json:"StatusSubscribe"`
	DateTransaction    time.Time `form:"DateTransaction" json:"DateTransaction"`
	TypeSubscribe      string    `form:"TypeSubscribe" json:"TypeSubscribe"`
	BeginDateSubscribe time.Time `form:"BeginDateSubscribe" json:"BeginDateSubscribe"`
	EndDateSubscribe   time.Time `form:"EndDateSubscribe" json:"EndDateSubscribe"`
	Token              string    `form:"Token" json:"Token"`
}

type SubscribeResponse struct {
	Response
	Data []Subscribe `form:"Data" json:"Data"`
}
