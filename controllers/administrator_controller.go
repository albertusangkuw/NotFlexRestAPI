package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/models"
	"github.com/gorilla/mux"
)

// Christian
// cari member berfungsi mencari data member berdasarkan email
// Method : GET
// Parameter : Params
// Nilai Parameter Wajib : Email
//
func SearchMemberBasedEmail(w http.ResponseWriter, r *http.Request) {
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
	var response models.MemberResponse
	email := r.URL.Query()["email"]
	query := "SELECT u.iduser, u.email, u.fullname, u.password, u.type, m.gender, m.countryoforigin, m.status, m.dateofbirth FROM user u JOIN member m where u.email = '" + email[0] + "' "
	resultSet, errQuery := DBConnection.Query(query)
	if errQuery != nil {
		ResponseManager(&response.Response, 500, errQuery.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var member models.Member
	var members []models.Member
	for resultSet.Next() {
		if err := resultSet.Scan(&member.IDUser, &member.Email, &member.FullName, &member.Password, &member.Type, &member.Gender, &member.OriginCountry, &member.Status, &member.DateOfBirth); err != nil {
			log.Print(err.Error())
		} else {
			members = append(members, member)
		}
	}

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

// Femi
// Blokir member berfungsi untuk melakukan suspend terhadap user
// Method : GET
// Parameter :Path Param
// Nilai Parameter Wajib : iduser
func BlockedMember(w http.ResponseWriter, r *http.Request) {
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
	iduser := vars["iduser"]
	query := "UPDATE member SET status = 0 WHERE iduser = ?"

	if len(iduser) == 0 {
		var response models.Response
		ResponseManager(&response, 400, " Iduser required")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var response models.Response
	resultSet, errQuery := DBConnection.Exec(query, iduser)

	if errQuery != nil {
		ResponseManager(&response, 500, errQuery.Error())
	} else {
		affectedRows, _ := resultSet.RowsAffected()
		if affectedRows > 0 {
			ResponseManager(&response, 200, "")
			fmt.Println("Id User ada")
		} else {
			ResponseManager(&response, 400, "No member was updated")
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
