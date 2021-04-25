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

	router.HandleFunc("/login", controllers.UserLogin).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")
	router.HandleFunc("/registrasi", controllers.MemberRegistration).Methods("POST")
	router.HandleFunc("/users/update", controllers.UpdateProfile).Methods("POST")

	router.HandleFunc("/registrasi", controllers.CariFilm).Methods("POST")
	router.HandleFunc("/registrasi", controllers.Menonton).Methods("POST")
	router.HandleFunc("/registrasi", controllers.TambahHistory).Methods("POST")

	router.HandleFunc("/administrator", controllers.CariMemberBerdasarkanEmail).Methods("GET")
	//router.HandleFunc("/administrator", controllers.UbahDataFilm).Methods("PUT")
	//router.HandleFunc("/administrator", controllers.CariDataFilmBerdasarkanJudul).Methods("GET")
	router.HandleFunc("/administrator", controllers.CariDataFilmBerdasarkanId).Methods("GET")

	router.HandleFunc("/cari", controllers.CariDataFilmBerdasarkanId).Methods("POST")
	/*
		router.HandleFunc("/users", controllers..GetAllUsers).Methods("GET")
		router.HandleFunc("/users/{userID}",controllers..GetAllUsers).Methods("GET")
		router.HandleFunc("/users/{userID}/photo", controllers.Authenticate(controllers.GetPhotoProfile, 0)).Methods("GT")
		router.HandleFunc("/users/{IDuser}/update", controllers.Authenticate(controllers.UpdateUser, 0)).Methods("PUT")
		router.HandleFunc("/users/{userID}/delete", controllers.Authenticate(controllers.DeleteUser, 0)).Methods("DELETE")

		/*
		router.HandleFunc("/users/{userID}/following/{anotherID}", controllers.Authenticate(c.FollowedUser, 0)).Methods("POST")
		router.HandleFunc("/users/{userID}/following/{anotherID}/remove", controllers.Authenticate(c.UnFollowedUser, 0)).Methods("DELETE")
		router.HandleFunc("/music", controllers.Authenticate(c.GetSong, 0)).Methods("GET")
		router.HandleFunc("/music/{IDmusic}", controllers.Authenticate(c.GetSong 0)).Methods("GET")
		router.HandleFunc("/music/{IDmusic}/data", c.GetSongFile).Methods("GET")

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
