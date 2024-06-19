package structures

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Music struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

type Playlist struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	UserID int     `json:"user_id"`
	Musics []Music `json:"musics"`
}
