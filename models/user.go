package models

type User struct {
	IDUser      string `form:"IDUser" json:"IDUser"`
	Nama        string `form:"Nama" json:"Nama"`
	Email       string `form:"Email" json:"Email"`
	NamaLengkap string `form:"NamaLengkap" json:"NamaLengkap"`
	Password    string `form:"Password" json:"Password"`
	Tipe        int    `form:"Tipe" json:"Tipe"`
}

type UserResponse struct {
	Response
	Data []User `form:"Data" json:"Data"`
}
