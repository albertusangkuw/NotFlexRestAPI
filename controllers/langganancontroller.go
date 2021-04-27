package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/NotFlexRestAPI/models"
)

// Elangel
func BerhentiBerlangganan(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	if !connect() {
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
	idUser := vars["idUser"]
	println(idUser)

	query := "DELETE FROM member WHERE iduser=?"
	_, errQuery := DBConnection.Exec(query, idUser)

	if errQuery != nil {
		ResponseManager(&response, 500, "Error Stop Langganan")
	} else {
		ResponseManager(&response, 200, "Success Stop Langganan")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
