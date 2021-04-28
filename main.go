package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NotFlexRestAPI/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	//Albertus Endpoint
	router.HandleFunc("/login", controllers.UserLogin).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")
	router.HandleFunc("/registrasi", controllers.MemberRegistration).Methods("POST")

	// Femi EndPoint
	router.HandleFunc("/film/cari", controllers.CariFilm).Methods("GET")
	router.HandleFunc("/film/{idfilm}", controllers.Menonton).Methods("POST")
	router.HandleFunc("/administrator/member/suspend/{iduser}", controllers.BlokirMember).Methods("GET")

	//Elangel Endpoint
	router.HandleFunc("/administrator/film/tambah", controllers.TambahFilm).Methods("POST")
	router.HandleFunc("/administrator/koleksi/tambah", controllers.TambahKoleksi).Methods("POST")
	router.HandleFunc("/administrator/koleksi/tambah/film", controllers.TambahFilmKeKoleksi).Methods("POST")
	router.HandleFunc("/administrator/detailkoleksi/lihat", controllers.LihatDetailKoleksi).Methods("GET")
	router.HandleFunc("/member/lihatHistoryFilm", controllers.LihatHistoryFilm).Methods("GET")
	router.HandleFunc("/member/update", controllers.UpdateProfile).Methods("PUT")
	router.HandleFunc("/member/berhentiBerlangganan/{idUser}", controllers.BerhentiBerlangganan).Methods("DELETE")

	//Christian endpoint
	router.HandleFunc("/administrator/member", controllers.CariMemberBerdasarkanEmail).Methods("GET")
	router.HandleFunc("/administrator/film/update", controllers.UpdateDataFilm).Methods("PUT")
	router.HandleFunc("/administrator/film", controllers.CariDataFilmBerdasarkanJudul).Methods("GET")
	router.HandleFunc("/administrator/film/{idfilm}", controllers.CariDataFilmBerdasarkanId).Methods("GET")

	//router.HandleFunc("/cari", controllers.CariDataFilmBerdasarkanId).Methods("POST")
	/*
		router.HandleFunc("/users", controllers..GetAllUsers).Methods("GET")
		router.HandleFunc("/users/{userID}",controllers..GetAllUsers).Methods("GET")
		router.HandleFunc("/users/{userID}/photo", controllers.Authenticate(controllers.GetPhotoProfile, 0)).Methods("GT")
		router.HandleFunc("/users/{IDuser}/update", controllers.Authenticate(controllers.UpdateUser, 0)).Methods("PUT")
		router.HandleFunc("/users/{userID}/delete", controllers.Authenticate(controllers.DeleteUser, 0)).Methods("DELETE")
	*/

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler := corsHandler.Handler(router)
	http.Handle("/", handler)
	fmt.Println("Connected to port 8085")
	log.Fatal(http.ListenAndServe(":8085", router))
}
