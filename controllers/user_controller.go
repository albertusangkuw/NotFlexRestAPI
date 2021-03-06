package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
)

// Albertus
// Member registrasi berfungsi
// Method : POST
// Parameter : Body  x-www-form-urlencoded
// Nilai Parameter Wajib : email, password
//
func MemberRegistration(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	// Melakukan pengecheckan ke server database
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

	email := r.Form.Get("email")
	namalengkap := r.Form.Get("namalengkap")
	password := r.Form.Get("password")
	tanggallahir := r.Form.Get("tanggallahir")
	jeniskelamin := r.Form.Get("jeniskelamin")
	asalnegara := r.Form.Get("asalnegara")

	if len(email) > 0 && len(namalengkap) > 0 && len(password) > 7 && len(tanggallahir) > 0 && len(jeniskelamin) > 0 && len(asalnegara) > 0 {
		ctx := context.Background()
		tx, err := DBConnection.BeginTx(ctx, nil)
		if err != nil {
			ResponseManager(&response, 500, "")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			log.Fatal(err)
			return
		}
		newIDUser := GetMD5Hash(email + namalengkap)
		queryUser := "INSERT INTO user VALUES(?,?,?,?,?)  "
		_, err = tx.ExecContext(ctx, queryUser, newIDUser, email, namalengkap, GetSHA256(password), 2)
		if err != nil {
			tx.Rollback()
			ResponseManager(&response, 500, err.Error()+" 1")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		queryMember := "INSERT INTO  member VALUES(?,?,?,?,?)"
		_, err = tx.ExecContext(ctx, queryMember, newIDUser, jeniskelamin, asalnegara, 1, tanggallahir)
		if err != nil {
			tx.Rollback()
			ResponseManager(&response, 500, err.Error()+" 2")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		err = tx.Commit()
		if err != nil {
			ResponseManager(&response, 500, err.Error())
			log.Fatal(err)
		} else {
			ResponseManager(&response, 200, "Success Insert New Member")
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		ResponseManager(&response, 400, "Form does not meet the requirement")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
}

// Albertus
// Member Login
// Method : Get
// Parameter : Body  x-www-form-urlencoded
// Nilai Parameter Wajib : email, password
//
func UserLogin(w http.ResponseWriter, r *http.Request) {
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	query := "SELECT iduser,type FROM user WHERE email=? AND password=? "

	email := r.URL.Query()["email"]
	password := r.URL.Query()["password"]

	var response models.Response
	if len(email) == 0 || len(password) == 0 {
		ResponseManager(&response, 400, "Login Failed")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	rows, err := DBConnection.Query(query, email[0], GetSHA256(password[0]))

	if err != nil {
		ResponseManager(&response, 500, "")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	//Check user login by name
	id := ""
	tipe := 0
	for rows.Next() {
		if err := rows.Scan(&id, &tipe); err != nil {
			log.Print(err.Error())
		}
	}
	status := true
	if tipe == 2 {
		query = "SELECT status FROM member WHERE iduser=?"
		rows, err = DBConnection.Query(query, id)
		if err != nil {
			ResponseManager(&response, 500, "")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
		statusUser := 0
		for rows.Next() {
			if err := rows.Scan(&statusUser); err != nil {
				log.Print(err.Error())
			}
		}
		if statusUser != 1 {
			status = false
		}
	}
	if status && len(id) > 0 {
		generateToken(w, id, email[0], tipe)
		ResponseManager(&response, 200, "Success Login")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		ResponseManager(&response, 401, "Login Failed")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&response)
	}
}

// Albertus
// Member Logout
// Method : Get
// Parameter :-
// Nilai Parameter Wajib : -
//
func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)
	var response models.Response
	ResponseManager(&response, 200, "Success Logout")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
