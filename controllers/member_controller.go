package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NotFlexRestAPI/models"
)

// Elangel
// UpdateProfile berfungsi untuk melakukan update data user member pada dataase, untuk data yang bisa diupdate hanya namalengkap, tanggallahir, dan jeniskelamin
// Method : GET
// Parameter : Body  x-www-form-urlencoded
// Nilai Parameter Wajib : iduser
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
		queryUser := "UPDATE user SET fullname=? WHERE iduser=?"
		queryMember := "UPDATE member SET dateofbirth=?, gender=? WHERE iduser=?"
		_, errQueryMember := DBConnection.Exec(queryMember, tanggallahir, jeniskelamin, idUser)
		_, errQueryUser := DBConnection.Exec(queryUser, namalengkap, idUser)
		if errQueryMember != nil && errQueryUser != nil {
			ResponseManager(&response, 400, errQueryMember.Error())
		} else {
			ResponseManager(&response, 200, "Success Update Profile")
		}
	} else {
		ResponseManager(&response, 400, "Error Update Profile")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
