package models

import "time"

type History struct {
	IdHistory    int       `form:"IdHistory" json:"IdHistory"`
	DateAccess time.Time `form:"DateAccess" json:"DateAccess"`
}

type HistoryResponse struct {
	Response
	Data []History `form:"data" json:"data"`
}
