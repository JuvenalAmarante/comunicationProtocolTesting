package playlists

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"restProtocol/pkg/structures"

	"github.com/gorilla/mux"
)

func CreatePlaylist(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var playlist structures.Playlist
	json.NewDecoder(r.Body).Decode(&playlist)

	err := db.QueryRow("INSERT INTO playlists(name, user_id) VALUES($1, $2) RETURNING id", playlist.Name, playlist.UserID).Scan(&playlist.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(playlist)
}

func GetPlaylists(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, user_id FROM playlists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var playlists []structures.Playlist
	for rows.Next() {
		var playlist structures.Playlist
		if err := rows.Scan(&playlist.ID, &playlist.Name, &playlist.UserID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		playlists = append(playlists, playlist)
	}

	json.NewEncoder(w).Encode(playlists)
}

func GetPlaylist(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id := params["id"]

	var playlist structures.Playlist
	err := db.QueryRow("SELECT id, name, user_id FROM playlists WHERE id=$1", id).Scan(&playlist.ID, &playlist.Name, &playlist.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(playlist)
}

func UpdatePlaylist(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id := params["id"]

	var playlist structures.Playlist
	json.NewDecoder(r.Body).Decode(&playlist)

	_, err := db.Exec("UPDATE playlists SET name=$1, user_id=$2 WHERE id=$3", playlist.Name, playlist.UserID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(playlist)
}

func DeletePlaylist(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM playlists WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
