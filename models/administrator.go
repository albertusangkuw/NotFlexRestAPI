package models

import time "time"

type Administrator struct {
	User
	LastLogin time.Time `form:"LastLogin" json:"LastLogin"`
}

type AdministratorResponse struct {
	Response
	Data []Administrator `form:"Data" json:"Data"`
}
