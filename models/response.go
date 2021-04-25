package models

type Response struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}
