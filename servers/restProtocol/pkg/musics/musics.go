package musics

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"restProtocol/pkg/structures"

	"github.com/gorilla/mux"
)

func CreateMusic(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var music structures.Music
	json.NewDecoder(r.Body).Decode(&music)

	err := db.QueryRow("INSERT INTO musics(name, artist) VALUES($1, $2) RETURNING id", music.Name, music.Artist).Scan(&music.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(music)
}

func GetMusics(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, artist FROM musics")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var musics []structures.Music
	for rows.Next() {
		var music structures.Music
		if err := rows.Scan(&music.ID, &music.Name, &music.Artist); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		musics = append(musics, music)
	}

	json.NewEncoder(w).Encode(musics)
}

func GetMusic(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id := params["id"]

	var music structures.Music
	err := db.QueryRow("SELECT id, name, artist FROM musics WHERE id=$1", id).Scan(&music.ID, &music.Name, &music.Artist)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(music)
}

func UpdateMusic(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id := params["id"]

	var music structures.Music
	json.NewDecoder(r.Body).Decode(&music)

	_, err := db.Exec("UPDATE musics SET name=$1, artist=$2 WHERE id=$3", music.Name, music.Artist, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(music)
}

func DeleteMusic(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM musics WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
