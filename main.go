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
	router.HandleFunc("/film/cari", controllers.Authenticate(controllers.SearchCollectionFilm, 0)).Methods("GET")
	router.HandleFunc("/film/{idfilm}", controllers.Authenticate(controllers.Watching, 0)).Methods("GET")
	router.HandleFunc("/administrator/member/suspend/{iduser}", controllers.Authenticate(controllers.BlokirMember, 1)).Methods("GET")

	//Elangel Endpoint
	router.HandleFunc("/administrator/film/tambah", controllers.Authenticate(controllers.AddFilm, 1)).Methods("POST")
	router.HandleFunc("/administrator/koleksi/tambah", controllers.Authenticate(controllers.AddCollection, 1)).Methods("POST")
	router.HandleFunc("/administrator/koleksi/tambah/film", controllers.Authenticate(controllers.AddFilmToCollection, 1)).Methods("POST")
	router.HandleFunc("/administrator/detailkoleksi/lihat", controllers.Authenticate(controllers.GetDetailCollection, 1)).Methods("GET")
	router.HandleFunc("/member/lihatHistoryFilm/{iduser}", controllers.Authenticate(controllers.GetHistoryFilm, 2)).Methods("GET")
	router.HandleFunc("/member/update", controllers.Authenticate(controllers.UpdateProfile, 2)).Methods("PUT")
	router.HandleFunc("/member/tambahBerlangganan/{iduser}", controllers.Authenticate(controllers.AddSubscribe, 2)).Methods("POST")
	router.HandleFunc("/member/berhentiBerlangganan/{iduser}", controllers.Authenticate(controllers.StopSubscribe, 2)).Methods("DELETE")

	//Christian endpoint
	router.HandleFunc("/administrator/member", controllers.Authenticate(controllers.SearchMemberBasedEmail, 1)).Methods("GET")
	router.HandleFunc("/administrator/film/update", controllers.Authenticate(controllers.UpdateDataFilm, 1)).Methods("PUT")
	router.HandleFunc("/administrator/film", controllers.Authenticate(controllers.SearchFilmDataBasedTitle, 1)).Methods("GET")
	router.HandleFunc("/administrator/film/{idfilm}", controllers.Authenticate(controllers.SearchFilmDataBasedId, 1)).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler := corsHandler.Handler(router)
	http.Handle("/", handler)
	fmt.Println("Connected to port 8085")
	log.Fatal(http.ListenAndServe(":8085", router))
}
