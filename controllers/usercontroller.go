package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/NotFlexRestAPI/models"
)

func contohControllerHandler(w http.ResponseWriter, r *http.Request) {
	// Melakukan pengecheckan ke server database
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	//Untuk Dapat value dari URL contoh http://host/users/uh972y7r2rh  => http://host/users/{iduser}
	err := r.ParseForm()
	if err != nil {
		var response models.Response
		ResponseManager(&response, 400, " Error when parsing form")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	vars := mux.Vars(r)
	iduser := vars["iduser"]
	println(iduser)

	//Untuk Dapat value dari GET  contoh http://host/users?name=Budi  => http://host/users
	testData := r.URL.Query()["name"]
	println(testData[0])

	//Untuk Dapat value dari POST  contoh http://host/users/new   {BODY x-www-form-url:  name="Ani"}    => http://host/users
	newName := r.Form.Get("name")
	println(newName)

	// Untuk Select atau mendapatkan data dari DB gunakan  Query
	query := "SELECT iduser,username,email FROM user"
	resultSet, _ := DBConnection.Query(query)
	for resultSet.Next() {
		//lakukan sesuatu dengan data
	}

	// Untuk Delete, Update, Insert gunakan Exec
	query = "INSERT INTO song_like(iduser,idsong) VALUES(?,?)"
	data1 := "1"
	data2 := "Indonesia Raya"
	_, errQuery := DBConnection.Exec(query, data1, data2)

	//Untuk menambahkan data diresponse yang sudah dibuat di model
	var response models.UserResponse
	ResponseManager(&response.Response, 200, "Success GET User")
	response.Data = nil

	//Mengatasi error
	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Custom Generate SQL
	filterBY, valuesList := GenerateSQLWhere(r, []string{"id", "title"}, "OR", "POST")
	// hasil filterBY (string) = id=? OR title=?
	query += "WHERE " + filterBY
	resultSet, _ = DBConnection.Query(query, valuesList)
}

// Albertus
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
		// Starting transaction manager https://www.sohamkamani.com/golang/sql-transactions/
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
			ResponseManager(&response, 500, "")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		queryMember := "INSERT INTO  member VALUES(?,?,?,?,?)"
		_, err = tx.ExecContext(ctx, queryMember, newIDUser, jeniskelamin, asalnegara, tanggallahir)
		if err != nil {
			tx.Rollback()
			ResponseManager(&response, 500, "")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		err = tx.Commit()
		if err != nil {
			ResponseManager(&response, 500, "")
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
func UserLogin(w http.ResponseWriter, r *http.Request) {
	if !connect() {
		var response models.Response
		ResponseManager(&response, 500, " Database Server Not Responding")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	query := "SELECT iduser FROM user WHERE email=? AND password=? "

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

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			log.Print(err.Error())
		}
	}

	if len(id) > 0 {
		//generateToken(w, id, name, 0)
		ResponseManager(&response, 200, "Success Login")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		ResponseManager(&response, 404, "Login Failed")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&response)
	}
}

// Albertus
func Logout(w http.ResponseWriter, r *http.Request) {
	//resetUserToken(w)
	var response models.Response
	ResponseManager(&response, 200, "Success Logout")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Elangel
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
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

	idUser := r.Form.Get("iduser")
	namalengkap := r.Form.Get("namalengkap")
	tanggallahir := r.Form.Get("tanggallahir")
	jeniskelamin := r.Form.Get("jeniskelamin")

	if len(idUser) > 0 && len(namalengkap) > 0 && len(tanggallahir) > 0 && len(jeniskelamin) > 0 {
		queryMember := "UPDATE member SET tanggalLahir=?, jenisKelamin=? WHERE iduser=?"
		queryUser := "UPDATE user SET namalengkap=? WHERE iduser=?"
		_, errQueryMember := DBConnection.Exec(queryMember, tanggallahir, jeniskelamin, idUser)
		_, errQueryUser := DBConnection.Exec(queryUser, namalengkap, idUser)
		if errQueryMember != nil && errQueryUser != nil {
			ResponseManager(&response, 400, errQueryMember.Error())
		} else {
			ResponseManager(&response, 200, "Success Update Profile")
		}
	} else {
		ResponseManager(&response, 400, "Error Insert New Film")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
