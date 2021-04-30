package models

type User struct {
	IDUser   string `form:"IDUser" json:"IDUser"`
	Name     string `form:"Name" json:"Name"`
	Email    string `form:"Email" json:"Email"`
	FullName string `form:FullName" json:FullName"`
	Password string `form:"Password" json:"Password"`
	Type     int    `form:"Type" json:"Type"`
}

type UserResponse struct {
	Response
	Data []User `form:"Data" json:"Data"`
}
