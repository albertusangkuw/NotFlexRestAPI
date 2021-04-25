package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
)

//christian
func CariMemberBerdasarkanEmail(w http.ResponseWriter, r *http.Request) {
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	err := r.ParseForm()
	if err != nil {
		return
	}

	email := r.URL.Query()["email"]
	query := "SELECT u.iduser, u.email, u.namalengkap, u.password, u.tipe, m.jeniskelamin, m.asalnegara, m.status, m.tanggallahir FROM user u JOIN member m where u.email = " + email[0] + " "
	resultSet, errQuery := DBConnection.Query(query)

	var member models.Member
	var members []models.Member
	for resultSet.Next() {
		if err := resultSet.Scan(&member.IDUser, &member.Email, &member.NamaLengkap, &member.Password, &member.JenisKelamin, &member.AsalNegara, &member.Status, &member.TanggalLahir); err != nil {
			log.Print(err.Error())
		} else {
			members = append(members, member)
		}
	}

	var response models.MemberResponse
	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		ResponseManager(&response.Response, 200, "Success GET User")
		response.Data = members
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

// Elangel
func TambahKoleksiFilm(w http.ResponseWriter, r *http.Request) { //insert film dan koleksi
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

	judul := r.Form.Get("judul")
	tahunRilis := r.Form.Get("tahunrilis")
	genre := r.Form.Get("genre")
	sutradara := r.Form.Get("sutradara")
	pemainUtama := r.Form.Get("pemainutama")
	sinopsis := r.Form.Get("sinopsis")
	namaKoleksi := r.Form.Get("namakoleksi")

	query := "INSERT INTO film(judul, tahunrilis, genre, sutradara, pemainutama, sinopsis) VALUES(?,?,?,?,?,?)"
	_, errQuery := DBConnection.Exec(query, judul, tahunRilis, genre, sutradara, pemainUtama, sinopsis)

	query = "INSERT INTO koleksi(namakoleksi) VALUES(?)"
	_, errQuery = DBConnection.Exec(query, namaKoleksi)
	var response models.Response
	if errQuery != nil {
		ResponseManager(&response, 500, "Error Insert New Film")
	} else {
		ResponseManager(&response, 200, "Success Insert New Film")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// //christian
/*
func UbahDataFilm(w http.ResponseWriter, r *http.Request) {
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	err := r.ParseForm()
	if err != nil {
		return
	}

	judul := r.URL.Query()["judul"]
	_, errQuery := db.Exec("UPDATE film SET judul=?, sinopsis=?, pemainutama=?, genre=?, sutradara=?, tahunrilis=? FROM film WHERE id =?",
		judul,
		sinopsis,
		pemainutama,
		genre,
		sutradara,
		tahunrilis,
		id,
	)

	var response models.FilmResponse
	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		ResponseManager(&response.Response, 200, "Success GET User")
		response.Data = films
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

//christian
func CariDataFilmBerdasarkanJudul(w http.ResponseWriter, r *http.Request) {
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	err := r.ParseForm()
	if err != nil {
		return
	}

	judul := r.URL.Query()["judul"]
	query := "SELECT judul, sinopsis, pemainutama, genre, sutradara, tahunrilis FROM film where judul = " + judul[0] + " "
	resultSet, errQuery := DBConnection.Query(query)
	var film models.Film
	var films []models.Film
	for resultSet.Next() {
		if err := resultSet.Scan(&film.IDFilm, &film.Judul, &film.Sinopsis, &film.PemainUtama, &film.Genre, &film.Sutradara, &film.TahunRilis); err != nil {
			log.Print(err.Error())
		} else {
			films = append(films, film)
		}
	}

	var response models.FilmResponse
	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		ResponseManager(&response.Response, 200, "Success GET User")
		response.Data = films
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
*/
//christian
func CariDataFilmBerdasarkanId(w http.ResponseWriter, r *http.Request) {
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	err := r.ParseForm()
	if err != nil {
		return
	}

	idFilm := r.URL.Query()["idFilm"]
	query := "SELECT judul, sinopsis, pemainutama, genre, sutradara, tahunrilis FROM film where judul = " + idFilm[0] + " "
	resultSet, errQuery := DBConnection.Query(query)
	var film models.Film
	var films []models.Film
	for resultSet.Next() {
		if err := resultSet.Scan(&film.IDFilm, &film.Judul, &film.Sinopsis, &film.PemainUtama, &film.Genre, &film.Sutradara, &film.TahunRilis); err != nil {
			log.Print(err.Error())
		} else {
			films = append(films, film)
		}
	}

	var response models.FilmResponse
	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		ResponseManager(&response.Response, 200, "Success GET User")
		response.Data = films
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}
