package controllers

import (
	"database/sql"
	"log"
)

var DBConnection *sql.DB = nil

var ipdatabase = "18.140.59.14"
var portnumber = "3306"
var username = "apitesting"
var password = "12345678"
var databasename = "db_api_not_flex"

func connect() bool {
	if DBConnection != nil {
		err := DBConnection.Ping()
		if err == nil {
			return true
		} else {
			log.Fatal(err)
			return false
		}
	}
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+ipdatabase+":"+portnumber+")/"+databasename+"?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	DBConnection = db
	return true
}
