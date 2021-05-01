package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/NotFlexRestAPI/models"
	"github.com/gorilla/mux"
)

// Albertus
//
// Method : POST
// Parameter : Path Params
// Nilai Parameter Wajib : -
var passlockTokenSubsribe = "rJe8ZzzWHVz-%3QnZQFH?pH5&-p^sTLfuXhW8!Z9-@=arWgt8H324!WKM%Nkp6AU"

func AddSubscribe(w http.ResponseWriter, r *http.Request) {
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
	vars := mux.Vars(r)
	iduser := vars["iduser"]
	tanggaltransaksi := r.Form.Get("tanggaltransaksi")
	jenislangganan := r.Form.Get("jenislangganan")
	awaltanggalberlangganan := r.Form.Get("awaltanggalberlangganan")

	bulanberlaku, errBulan := strconv.Atoi(r.Form.Get("bulanberlaku"))
	tahunberlaku, errTahun := strconv.Atoi(r.Form.Get("tahunberlaku"))
	nomorkartukredit := r.Form.Get("nomorkartukredit")
	kodecvc := r.Form.Get("kodecvc")

	if errBulan != nil || errTahun != nil {
		ResponseManager(&response, 400, " Year or Month is not valid")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	year, month, _ := time.Now().Date()
	if bulanberlaku < int(month) || tahunberlaku < year {
		ResponseManager(&response, 400, " Card is expired")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if len(kodecvc) < 3 || len(kodecvc) > 5 || len(nomorkartukredit) < 12 {
		ResponseManager(&response, 400, " Not valid cvc code or card number")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	if !checkLuhn(nomorkartukredit) {
		ResponseManager(&response, 400, " Not valid card number")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	if len(iduser) > 0 && len(tanggaltransaksi) > 0 && len(jenislangganan) > 0 && len(awaltanggalberlangganan) > 0 {
		awalberlangganan, err := time.Parse("2006-01-02 15:04:05", awaltanggalberlangganan)
		akhirberlangganan := awalberlangganan.Add(30 * 24 * time.Hour)

		token := GetSHA256(GetSHA256(iduser + tanggaltransaksi + passlockTokenSubsribe + jenislangganan + awaltanggalberlangganan))

		query := "INSERT INTO subscribe(iduser,subscriptionstatus,transactiondate,subscriptiontype, startofsubscriptiondate, endofthesubscriptiondate,token) VALUES(?,?,?,?,?,?,?)  "
		_, err = DBConnection.Exec(query, iduser, "1", tanggaltransaksi, jenislangganan, awaltanggalberlangganan, akhirberlangganan.Format("2006-01-02 15:04:05"), token)
		if err != nil {
			ResponseManager(&response, 500, err.Error())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
		ResponseManager(&response, 200, "Success Insert New Langganan")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		ResponseManager(&response, 400, "Form does not meet the requirement")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
}

//Algorithm from https://www.geeksforgeeks.org/luhn-algorithm/
func checkLuhn(cardNo string) bool {
	nDigits := len(cardNo)

	nSum := 0
	isSecond := false
	for i := nDigits - 1; i >= 0; i-- {
		var d int
		d = int(cardNo[i] - '0')

		if isSecond == true {
			d = d * 2
		}

		nSum += d / 10
		nSum += d % 10

		isSecond = !isSecond
	}
	return (nSum%10 == 0)
}

// Elangel
// StopSubscribe berfungsi untuk melakukan update status data langganan pada database agar user tidak berlangganan lagi
// Method : UPDATE
// Parameter : Path Params
// Nilai Parameter Wajib : iduser
func StopSubscribe(w http.ResponseWriter, r *http.Request) {
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
	vars := mux.Vars(r)
	idUser := vars["iduser"]

	if len(idLangganan) > 0 && len(idUser) > 0 {
		query := "UPDATE subscribe SET subscriptionstatus=? WHERE 	idsubscription=? AND iduser=?"
		_, errQuery := DBConnection.Exec(query, 0, idLangganan, idUser)
		if errQuery != nil {
			ResponseManager(&response, 400, errQuery.Error())
		} else {
			ResponseManager(&response, 200, "Success Stop Langganan ")
		}
	} else {
		ResponseManager(&response, 400, "Error Stop Langganan "+idLangganan)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
