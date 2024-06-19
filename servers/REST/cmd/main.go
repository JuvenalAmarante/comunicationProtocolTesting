package main

import (
	"REST/pkg/database"
	"REST/pkg/musics"
	"REST/pkg/playlists"
	"REST/pkg/users"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var con = database.Init()

	r := mux.NewRouter()

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users.CreateUser(w, r, con)
	}).Methods("POST")
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users.GetUsers(w, r, con)
	}).Methods("GET")
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		users.GetUser(w, r, con)
	}).Methods("GET")
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		users.UpdateUser(w, r, con)
	}).Methods("PUT")
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		users.DeleteUser(w, r, con)
	}).Methods("DELETE")

	r.HandleFunc("/musics", func(w http.ResponseWriter, r *http.Request) {
		musics.CreateMusic(w, r, con)
	}).Methods("POST")
	r.HandleFunc("/musics", func(w http.ResponseWriter, r *http.Request) {
		musics.GetMusics(w, r, con)
	}).Methods("GET")
	r.HandleFunc("/musics/{id}", func(w http.ResponseWriter, r *http.Request) {
		musics.GetMusic(w, r, con)
	}).Methods("GET")
	r.HandleFunc("/musics/{id}", func(w http.ResponseWriter, r *http.Request) {
		musics.UpdateMusic(w, r, con)
	}).Methods("PUT")
	r.HandleFunc("/musics/{id}", func(w http.ResponseWriter, r *http.Request) {
		musics.DeleteMusic(w, r, con)
	}).Methods("DELETE")

	r.HandleFunc("/playlists", func(w http.ResponseWriter, r *http.Request) {
		playlists.CreatePlaylist(w, r, con)
	}).Methods("POST")
	r.HandleFunc("/playlists", func(w http.ResponseWriter, r *http.Request) {
		playlists.GetPlaylists(w, r, con)
	}).Methods("GET")
	r.HandleFunc("/playlists/{id}", func(w http.ResponseWriter, r *http.Request) {
		playlists.GetPlaylist(w, r, con)
	}).Methods("GET")
	r.HandleFunc("/playlists/{id}", func(w http.ResponseWriter, r *http.Request) {
		playlists.UpdatePlaylist(w, r, con)
	}).Methods("PUT")
	r.HandleFunc("/playlists/{id}", func(w http.ResponseWriter, r *http.Request) {
		playlists.DeletePlaylist(w, r, con)
	}).Methods("DELETE")

	fmt.Print("Rodando na porta 8000!")

	log.Fatal(http.ListenAndServe(":8000", r))
}
