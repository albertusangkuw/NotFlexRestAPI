package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/NotFlexRestAPI/models"
	"github.com/gorilla/mux"
)

// Femi
// Tambah history berfungsi untuk menambah history setelah member selesai menonton
// Method : POST
// Parameter : -
// Nilai Parameter Wajib : iduser, tanggalakses, idfilm
func AddHistory(w http.ResponseWriter, r *http.Request) {
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

	_, iduser, _, _ := validateUserTokenFromCookies(r)
	tanggalakses := time.Now()
	vars := mux.Vars(r)
	idfilm := vars["idfilm"]

	query := "INSERT INTO history(iduser, accessdate, idfilm) VALUES(?,?,?)"
	_, errQuery := DBConnection.Exec(query, iduser, tanggalakses, idfilm)

	if errQuery != nil {
		println("Tidak berhasil di insert")
	} else {
		println("Berhasil di insert")
	}
}

// Elangel
// Getungsi untuk melihat seluruh data history pada dataase
// Method : GET
// Parameter : Query Params
// Nilai Parameter Wajib : iduser
func GetHistoryFilm(w http.ResponseWriter, r *http.Request) {
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
	vars := mux.Vars(r)
	idUser := vars["iduser"]
	if len(idUser) == 0 {
		var response models.Response
		ResponseManager(&response, 400, " Error when parsing form")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	query := "SELECT h.idhistory, f.idfilm, f.title, f.synopsis, f.genre, h.accessdate FROM history h JOIN film f ON h.idfilm = f.idfilm WHERE h.iduser =? "
	resultSet, errQuery := DBConnection.Query(query, idUser)

	var history models.History
	var film models.Film
	var films []models.Film
	if errQuery != nil {
		var response models.Response
		ResponseManager(&response, 500, errQuery.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	for resultSet.Next() {
		if err := resultSet.Scan(&history.IdHistory, &film.IDFilm, &film.Title, &film.Sinopsis, &film.Genre, &history.DateAccess); err != nil {
			log.Println(err.Error())
		} else {
			films = append(films, film)
		}
	}

	var response models.FilmResponse
	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		ResponseManager(&response.Response, 200, "Success GET History")
		response.Data = films
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
