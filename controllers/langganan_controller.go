package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NotFlexRestAPI/models"
)

// Albertus
//
// Method : POST
// Parameter : Path Params
// Nilai Parameter Wajib : iduser
func TambahLangganan(w http.ResponseWriter, r *http.Request) {

}

// Elangel
// BerhentiBerlangganan berfungsi untuk melakukan delete data langganan pada database agar user tidak berlangganan lagi
// Method : DELETE
// Parameter : Path Params
// Nilai Parameter Wajib : iduser
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

	idLangganan := r.Form.Get("idlangganan")
	idUser := r.Form.Get("idlangganan")
	statusLangganan := r.Form.Get("statuslangganan")

	if len(idLangganan) > 0 && len(idUser) > 0 && len(statusLangganan) > 0 {
		query := "UPDATE berlangganan SET statuslangganan=? WHERE idlangganan=? AND iduser=?"
		_, errQuery := DBConnection.Exec(query, statusLangganan, idLangganan, idUser)
		if errQuery != nil {
			ResponseManager(&response, 400, errQuery.Error())
		} else {
			ResponseManager(&response, 200, "Success Stop Langganan")
		}
	} else {
		ResponseManager(&response, 400, "Error Stop Langganan")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
