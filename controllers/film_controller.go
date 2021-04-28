package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
	"github.com/gorilla/mux"
)

// Elangel
// TambahFilm berfungsi untuk melakukan insert data film baru ke dalam database
// Method : POST
// Parameter : Body  x-www-form-urlencoded
// Nilai Parameter Wajib : judul
func TambahFilm(w http.ResponseWriter, r *http.Request) {
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

	judul := r.Form.Get("judul")
	sinopsis := r.Form.Get("sinopsis")
	pemainUtama := r.Form.Get("pemainutama")
	genre := r.Form.Get("genre")
	sutradara := r.Form.Get("sutradara")
	tahunRilis := r.Form.Get("tahunrilis")

	if len(judul) > 0 {
		query := "INSERT INTO film(judul, sinopsis, pemainutama, genre, sutradara, tahunrilis) VALUES(?,?,?,?,?,?)"
		_, errQuery := DBConnection.Exec(query, judul, sinopsis, pemainUtama, genre, sutradara, tahunRilis)

		if errQuery != nil {
			ResponseManager(&response, 500, "Error Insert New Film")
		} else {
			ResponseManager(&response, 200, "Success Insert New Film")
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

// femi
// Cari film berfungsi untuk melihat film berdasarkan judul / sutradara / pemain utama / sinopsis / pemainutama / tahunrilis/ genre,
// salah satu dari ini bisa munculin filmnya
// Method : GET
// Parameter : Query params
// Nilai Parameter Wajib : -
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

	query := "SELECT judul, sinopsis, pemainutama, genre, sutradara, tahunrilis FROM film"
	filterBY, valuesList := GenerateSQLWhere(r, []string{"judul", "sinopsis", "pemainutama", "tahunrilis", "genre", "sutradara"}, "OR", "GET")
	query += " WHERE " + filterBY
	println(query)

	resultSet, errQuery := DBConnection.Query(query, valuesList...)
	var Film models.Film
	var Films []models.Film

	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		for resultSet.Next() {
			if errors := resultSet.Scan(&Film.Judul, &Film.Sinopsis, &Film.PemainUtama, &Film.Genre, &Film.Sutradara, &Film.TahunRilis); errors != nil {

			} else {
				Films = append(Films, Film)
			}
		}
		if len(Films) > 0 {
			response.Data = Films
			ResponseManager(&response.Response, 200, " ")
		} else {
			ResponseManager(&response.Response, 400, errQuery.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Christian
// cari data film berfungsi mencari data film berdasarkan judul
// Method : GET
// Parameter : para
// Nilai Parameter Wajib : judul
//
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
	query := "SELECT judul, sinopsis, pemainutama, genre, sutradara, tahunrilis FROM film where judul = '" + judul[0] + "' "
	resultSet, errQuery := DBConnection.Query(query)
	var film models.Film
	var films []models.Film
	for resultSet.Next() {
		if err := resultSet.Scan(&film.Judul, &film.Sinopsis, &film.PemainUtama, &film.Genre, &film.Sutradara, &film.TahunRilis); err != nil {
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

// Christian
// cari data film berfungsi mencari data film berdasarkan id
// Method : GET
// Parameter : Params
// Nilai Parameter Wajib : ID film
//
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

	vars := mux.Vars(r)
	idFilm := vars["idfilm"]
	query := "SELECT judul, sinopsis, pemainutama, genre, sutradara, tahunrilis FROM film WHERE idfilm = " + idFilm + " "
	resultSet, errQuery := DBConnection.Query(query)
	if errQuery != nil {
		var response models.Response
		ResponseManager(&response, 500, errQuery.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var film models.Film
	var films []models.Film
	for resultSet.Next() {
		if err := resultSet.Scan(&film.Judul, &film.Sinopsis, &film.PemainUtama, &film.Genre, &film.Sutradara, &film.TahunRilis); err != nil {
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

// Christian
//	Update Data Film berfungsi mengubah data film berdasarkan id film
// Method : PUT
// Parameter : path params
// Nilai Parameter Wajib : ID Film
// Nilai Parameter Pilihan : Judul, Sinopsis, Pemain Utama, Genre, Sutradara, Tahun Rilis
//
func UpdateDataFilm(w http.ResponseWriter, r *http.Request) {
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

	idFilm := r.Form.Get("idFilm")

	query := "UPDATE film SET "
	filterBY, valuesList := GenerateSQLWhere(r, []string{"judul", "sinopsis", "pemainutama", "genre", "sutradara", "tahunrilis"}, ",", "GET")
	valuesList = append(valuesList, idFilm)
	query += filterBY + " WHERE idfilm=?"
	resultSet, err := DBConnection.Exec(query, valuesList...)

	var response models.Response
	if err != nil {
		ResponseManager(&response, 500, err.Error())
	} else {
		nums, _ := resultSet.RowsAffected()
		if nums > 0 {
			ResponseManager(&response, 200, "Success UPDATE User")
		} else {
			ResponseManager(&response, 400, "Bad Request - Update Failed")
		}

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

// femi
// Menonton berfungsi berfungsi untuk menonton, dengan memasukkan idfilm, setelah menonton(sudah memasukkan id film)
// nantinya akan tercatat di history
// Method : POST
// Parameter : Path params
// Nilai Parameter Wajib : idfilm
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

	vars := mux.Vars(r)
	idfilm := vars["idfilm"]
	println(idfilm)
	query := "SELECT judul, sinopsis, pemainutama, genre, sutradara, tahunrilis FROM film WHERE idfilm = ?"

	if len(idfilm) == 0 {
		var response models.Response
		ResponseManager(&response, 400, " Error when parsing form")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var response models.FilmResponse
	resultSet, errQuery := DBConnection.Query(query, idfilm)
	var Film models.Film
	var Films []models.Film

	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		for resultSet.Next() {
			if errors := resultSet.Scan(&Film.Judul, &Film.Sinopsis, &Film.PemainUtama, &Film.Genre, &Film.Sutradara, &Film.TahunRilis); errors != nil {

			} else {
				Films = append(Films, Film)
			}
		}
		if len(Films) > 0 {
			response.Data = Films
			ResponseManager(&response.Response, 200, " ")
			TambahHistory(w, r)
		} else {
			ResponseManager(&response.Response, 400, errQuery.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
