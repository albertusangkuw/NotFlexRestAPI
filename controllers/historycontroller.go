package controllers

import (
	"encoding/json"
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
