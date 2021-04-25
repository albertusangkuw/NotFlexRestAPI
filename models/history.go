package models

import "time"

type History struct {
	IdHistory    int       `form:"IdHistory" json:"IdHistory"`
	TanggalAkses time.Time `form:"TanggalAkses" json:"TanggalAkses"`
}

type HistoryResponse struct {
	Response
	Data []History `form:"data" json:"data"`
}
