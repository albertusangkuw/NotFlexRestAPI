package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
	"github.com/gorilla/mux"
)

// Femi
// Tambah history berfungsi untuk menambah history setelah member selesai menonton
// Method : POST
// Parameter : Path param
// Nilai Parameter Wajib : iduser, tanggalakses, idfilm
func TambahHistory(w http.ResponseWriter, r *http.Request) {
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

	iduser := r.Form.Get("iduser")
	tanggalakses := r.Form.Get("tanggalakses")
	vars := mux.Vars(r)
	idfilm := vars["idfilm"]

	query := "INSERT INTO history(iduser, tanggalakses, idfilm) VALUES(?,?,?)"
	_, errQuery := DBConnection.Exec(query, iduser, tanggalakses, idfilm)

	if errQuery != nil {
		println("Tidak berharil di insert")
	} else {
		println("Berharil di insert")
	}
}

// Elangel
// LihatHistoryFilm berfungsi untuk melihat seluruh data history pada dataase
// Method : GET
// Parameter : Query Params
// Nilai Parameter Wajib : iduser
func LihatHistoryFilm(w http.ResponseWriter, r *http.Request) {
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

	idUser := r.URL.Query()["iduser"]
	query := "SELECT h.idhistory, f.idfilm, f.judul, f.sinopsis, f.genre, h.tanggalakses FROM history h JOIN film f ON h.idfilm = f.idfilm WHERE h.iduser = " + idUser[0]
	resultSet, errQuery := DBConnection.Query(query)

	var history models.History
	var film models.Film
	var films []models.Film
	for resultSet.Next() {
		if err := resultSet.Scan(&history.IdHistory, &film.IDFilm, &film.Judul, &film.Sinopsis, &film.Genre, &history.TanggalAkses); err != nil {
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
