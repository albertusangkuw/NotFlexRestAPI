package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
)

// Elangel
// TambahKoleksi berfungsi untuk melakukan insert data koleksi baru ke dalam database
// Method : POST
// Parameter : Body  x-www-form-urlencoded
// Nilai Parameter Wajib : namakoleksi
func AddCollection(w http.ResponseWriter, r *http.Request) {
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

	if len(namaKoleksi) > 0 {
		query := "INSERT INTO collection(collectionname) VALUES(?)"
		_, errQuery := DBConnection.Exec(query, namaKoleksi)

		if errQuery != nil {
			ResponseManager(&response, 500, "Error Insert New Colection")
		} else {
			ResponseManager(&response, 200, "Success Insert New Colection")
		}
	} else {
		ResponseManager(&response, 400, "Form does not meet the requirement")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Elangel
// TambahFilmKeKoleksi berfungsi untuk melakukan insert data film baru dan koleksi ke dalam detail koleksi pada database
// Method : POST
// Parameter : Body  x-www-form-urlencoded
// Nilai Parameter Wajib : idkoleksi, idlm
func AddFilmToCollection(w http.ResponseWriter, r *http.Request) {
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

	if len(idKoleksi) > 0 && len(idFilm) > 0 {
		query := "INSERT INTO collectiondetail(idcollection, idfilm) VALUES(?,?)"
		_, errQuery := DBConnection.Exec(query, idKoleksi, idFilm)

		if errQuery != nil {
			ResponseManager(&response, 500, "Error Insert New Film to Collection")
		} else {
			ResponseManager(&response, 200, "Success Insert New Film to Collection")
		}
	} else {
		ResponseManager(&response, 400, "Form does not meet the requirement")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Elangel
// LihatDetailKoleksi berfungsi untuk melihat seluruh data detail koleksi pada dataase
// Method : GET
// Parameter : Query Params
// Nilai Parameter Wajib : idDetailKoleksi
func GetDetailCollection(w http.ResponseWriter, r *http.Request) {
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
	if len(idDetailKoleksi) == 0 {
		var response models.Response
		ResponseManager(&response, 400, "")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	query := "SELECT dk.iddetailcollection, k.idcollection, k.collectionname, f.idfilm, f.title, f.synopsis, f.mainactor, f.genre, f.director, f.releaseyear FROM collectiondetail dk INNER JOIN collection k ON dk.idcollection = k.idcollection INNER JOIN film f ON dk.idfilm = f.idfilm WHERE dk.iddetailcollection = " + idDetailKoleksi[0]
	//quer?y := "SELECT dk.iddetailcollection, k.idcollection, k.collectionname, f.idfilm, f.title, f.synopsis, f.mainactor, f.genre, f.director, f.releaseyear FROM detailcollection dk INNER JOIN koleksi k ON dk.idcollection = k.idcollection INNER JOIN film f ON dk.idfilm = f.idfilm WHERE dk.iddetailcollection = " + idDetailKoleksi[0]
	resultSet, errQuery := DBConnection.Query(query)
	if errQuery != nil {
		var response models.Response
		ResponseManager(&response, 500, errQuery.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var detailCollection models.DetailCollection
	var detailCollections []models.DetailCollection
	for resultSet.Next() {
		if err := resultSet.Scan(
			&detailCollection.IDDetailCollection,
			&detailCollection.IDCollection.IDCollection,
			&detailCollection.IDCollection.NameCollection,
			&detailCollection.IDFilm.IDFilm,
			&detailCollection.IDFilm.Title,
			&detailCollection.IDFilm.Sinopsis,
			&detailCollection.IDFilm.MainCast,
			&detailCollection.IDFilm.Genre,
			&detailCollection.IDFilm.Director,
			&detailCollection.IDFilm.ReleaseYear); err != nil {
			log.Println(err.Error())
		} else {
			detailCollections = append(detailCollections, detailCollection)
		}
	}

	var response models.DetailCollectionResponse
	ResponseManager(&response.Response, 200, "Success GET Detail Collection")
	response.Data = detailCollections

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
