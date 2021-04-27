package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
)

// Elangel
func TambahKoleksi(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	if !connect() {
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	err := r.ParseForm()
	if err != nil {
		ResponseManager(&response, 400, " Error when parsing form")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	namaKoleksi := r.Form.Get("namakoleksi")

	query := "INSERT INTO koleksi(namakoleksi) VALUES(?)"
	_, errQuery := DBConnection.Exec(query, namaKoleksi)

	if errQuery != nil {
		ResponseManager(&response, 500, "Error Insert New Colection")
	} else {
		ResponseManager(&response, 200, "Success Insert New Colection")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Elangel
func TambahFilmKeKoleksi(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	if !connect() {
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	err := r.ParseForm()
	if err != nil {
		ResponseManager(&response, 400, " Error when parsing form")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	idKoleksi := r.Form.Get("idkoleksi")
	idFilm := r.Form.Get("idfilm")

	query := "INSERT INTO detailkoleksi(idkoleksi, idfilm) VALUES(?,?)"
	_, errQuery := DBConnection.Exec(query, idKoleksi, idFilm)

	if errQuery != nil {
		ResponseManager(&response, 500, "Error Insert New Film to Collection")
	} else {
		ResponseManager(&response, 200, "Success Insert New Film to Collection")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Elangel
func LihatDetailKoleksi(w http.ResponseWriter, r *http.Request) {
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

	idDetailKoleksi := r.URL.Query()["iddetailkoleksi"]
	query := "SELECT dk.iddetailkoleksi, k.idkoleksi, k.namakoleksi, f.idfilm, f.judul, f.sinopsis, f.pemainutama, f.genre, f.sutradara, f.tahunrilis FROM detailkoleksi dk INNER JOIN koleksi k ON dk.idkoleksi = k.idkoleksi INNER JOIN film f ON dk.idfilm = f.idfilm WHERE dk.iddetailkoleksi = " + idDetailKoleksi[0]
	resultSet, errQuery := DBConnection.Query(query)

	var detailCollection models.DetailKoleksi
	var detailCollections []models.DetailKoleksi
	for resultSet.Next() {
		if err := resultSet.Scan(
			&detailCollection.IDDetailKoleksi,
			&detailCollection.IDKoleksi.IDKoleksi,
			&detailCollection.IDKoleksi.NamaKoleksi,
			&detailCollection.IDFilm.IDFilm,
			&detailCollection.IDFilm.Judul,
			&detailCollection.IDFilm.Sinopsis,
			&detailCollection.IDFilm.PemainUtama,
			&detailCollection.IDFilm.Genre,
			&detailCollection.IDFilm.Sutradara,
			&detailCollection.IDFilm.TahunRilis); err != nil {
			log.Println(err.Error())
		} else {
			detailCollections = append(detailCollections, detailCollection)
		}
	}

	var response models.DetailKoleksiResponse
	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		ResponseManager(&response.Response, 200, "Success GET Detail Collection")
		response.Data = detailCollections
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
