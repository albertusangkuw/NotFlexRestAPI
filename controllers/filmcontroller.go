package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NotFlexRestAPI/models"
)

// femi
func CariFilm(w http.ResponseWriter, r *http.Request) {
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	err := r.ParseForm()
	if err != nil {
		var response models.Response
		ResponseManager(&response, 400, " Error when parsing form")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var response models.FilmResponse

	query := "SELECT * FROM film"
	filterBY, valuesList := GenerateSQLWhere(r, []string{"Judul", "Sutradara", "TahunRilis", "Genre", "PemainUtama", "Sinopsis"}, "OR", "GET")
	// hasil filterBY (string) = id=? OR title=?
	query += "WHERE " + filterBY
	resultSet, errQuery := DBConnection.Query(query, valuesList)

	if errQuery == nil {
		ResponseManager(&response.Response, 404, "Data Not Found")
	} else {
		println(resultSet)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

//femi
func Menonton(w http.ResponseWriter, r *http.Request) {
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	err := r.ParseForm()
	if err != nil {
		var response models.Response
		ResponseManager(&response, 400, " Error when parsing form")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	query := "SELECT ....  FROM .... "
	filterBY, valuesList := GenerateSQLWhere(r, []string{"Judul", "Sutradara", "TahunRilis", "Genre", "PemainUtama", "Sinopsis"}, "OR", "GET")
	// hasil filterBY (string) = id=? OR title=?
	query += "WHERE " + filterBY
	_, err = DBConnection.Query(query, valuesList)
	println(query)
	if err == nil {
		var response models.Response
		ResponseManager(&response, 404, "Data Not Found")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		TambahHistory(w, r)
	}
}
