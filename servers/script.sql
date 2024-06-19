CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    age INT
);

CREATE TABLE musics (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    artist VARCHAR(100)
);

CREATE TABLE playlists (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    user_id INT REFERENCES users(id)
);

CREATE TABLE playlists_musics (
    playlist_id INT REFERENCES playlists(id),
    music_id INT REFERENCES musics(id),
    PRIMARY KEY (playlist_id, music_id)
);