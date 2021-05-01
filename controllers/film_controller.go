package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
	"github.com/gorilla/mux"
)

// Elangel
// AddFilm berfungsi untuk melakukan insert data film baru ke dalam database
// Method : POST
// Parameter : Body  x-www-form-urlencoded
// Nilai Parameter Wajib : judul
func AddFilm(w http.ResponseWriter, r *http.Request) {
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
		query := "INSERT INTO film(title, synopsis, mainactor, genre, director, releaseyear) VALUES(?,?,?,?,?,?)"
		_, errQuery := DBConnection.Exec(query, judul, sinopsis, pemainUtama, genre, sutradara, tahunRilis)

		if errQuery != nil {
			ResponseManager(&response, 500, "Error Insert New Film"+errQuery.Error())
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
// Nilai Parameter Optional : judul, sinopsis, pemainutama, genre, sutradara
func SearchCollectionFilm(w http.ResponseWriter, r *http.Request) {
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

	query := "SELECT title, synopsis, mainactor, genre, director, releaseyear FROM film"
	filterBY, valuesList := GenerateSQLWhere(r, []string{"title", "synopsis", "mainactor", "releaseyear", "genre", "director"}, "OR", "GET")
	query += " WHERE " + filterBY
	println(query)

	resultSet, errQuery := DBConnection.Query(query, valuesList...)
	var Film models.Film
	var Films []models.Film

	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
	} else {
		for resultSet.Next() {
			if errors := resultSet.Scan(&Film.Title, &Film.Sinopsis, &Film.MainCast, &Film.Genre, &Film.Director, &Film.ReleaseYear); errors != nil {

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
// Parameter : params
// Nilai Parameter Wajib : judul
//
func SearchFilmDataBasedTitle(w http.ResponseWriter, r *http.Request) {
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

	title := r.URL.Query()["title"]
	query := "SELECT title, synopsis, mainactor, genre, director, releaseyear FROM film where title = '" + title[0] + "' "
	resultSet, errQuery := DBConnection.Query(query)
	var film models.Film
	var films []models.Film
	for resultSet.Next() {
		if err := resultSet.Scan(&film.Title, &film.Sinopsis, &film.MainCast, &film.Genre, &film.Director, &film.ReleaseYear); err != nil {
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
func SearchFilmDataBasedId(w http.ResponseWriter, r *http.Request) {
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
	query := "SELECT title, synopsis, mainactor, genre, director, releaseyear FROM film WHERE idfilm = " + idFilm + " "
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
		if err := resultSet.Scan(&film.Title, &film.Sinopsis, &film.MainCast, &film.Genre, &film.Director, &film.ReleaseYear); err != nil {
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
	filterBY, valuesList := GenerateSQLWhere(r, []string{"title", "synopsis", "mainactor", "releaseyear", "genre", "director"}, ",", "GET")
	valuesList = append(valuesList, idFilm)
	query = query + " " + filterBY + " WHERE idfilm=?"
	println(query)
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
// Watching berfungsi berfungsi untuk Watching, dengan memasukkan idfilm, setelah Watching(sudah memasukkan id film)
// nantinya akan tercatat di history
// Method : GET
// Parameter : Path params
// Nilai Parameter Wajib : idfilm
func Watching(w http.ResponseWriter, r *http.Request) {
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
	query := "SELECT title, synopsis, mainactor, genre, director, releaseyear FROM film WHERE idfilm = ?"

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
			if errors := resultSet.Scan(&Film.Title, &Film.Sinopsis, &Film.MainCast, &Film.Genre, &Film.Director, &Film.ReleaseYear); errors != nil {
				println(errors.Error())
			} else {
				//	Film.IDFilm, _= strconv.ParseUint(strconv.Atoi(idfilm))
				Films = append(Films, Film)
			}
		}
		if len(Films) > 0 {
			response.Data = Films
			ResponseManager(&response.Response, 200, "")
			AddHistory(w, r)
		} else {
			ResponseManager(&response.Response, 400, errQuery.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
