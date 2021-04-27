package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
)

// Femi
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
	var response models.Response
	iduser := r.Form.Get("iduser")
	tanggalakses := r.Form.Get("tanggalakses")
	idfilm := r.Form.Get("idfilm")

	query := "INSERT INTO history(idhistory, tanggalakses, idfilm) VALUES(?,?,?)"
	_, errQuery := DBConnection.Exec(query, iduser, tanggalakses, idfilm)

	if errQuery != nil {
		ResponseManager(&response, 500, "Error Insert New History")
	} else {
		ResponseManager(&response, 200, "Success Insert New History")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Elangel
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
