package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"

	"gRPC/pkg/database"
	pb "gRPC/service" // Altere para o caminho correto

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedMusicServiceServer
	db *sql.DB
}

func (s *server) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	err := s.db.QueryRow("INSERT INTO users(name, age) VALUES($1, $2) RETURNING id", in.Name, in.Age).Scan(&in.Id)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *server) GetUsers(ctx context.Context, in *pb.Empty) (*pb.Users, error) {
	rows, err := s.db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var usuario pb.User
		if err := rows.Scan(&usuario.Id, &usuario.Name, &usuario.Age); err != nil {
			return nil, err
		}
		users = append(users, &usuario)
	}

	return &pb.Users{Users: users}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	var usuario pb.User
	err := s.db.QueryRow("SELECT id, name, age FROM users WHERE id=$1", in.Id).Scan(&usuario.Id, &usuario.Name, &usuario.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario not found")
		}
		return nil, err
	}
	return &usuario, nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	_, err := s.db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", in.Name, in.Age, in.Id)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.UserId) (*pb.Empty, error) {
	_, err := s.db.Exec("DELETE FROM users WHERE id=$1", in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *server) CreateMusic(ctx context.Context, in *pb.Music) (*pb.Music, error) {
	err := s.db.QueryRow("INSERT INTO musics(name, artist) VALUES($1, $2) RETURNING id", in.Name, in.Artist).Scan(&in.Id)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *server) GetMusics(ctx context.Context, in *pb.Empty) (*pb.Musics, error) {
	rows, err := s.db.Query("SELECT id, name, artist FROM musics")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var musics []*pb.Music
	for rows.Next() {
		var music pb.Music
		if err := rows.Scan(&music.Id, &music.Name, &music.Artist); err != nil {
			return nil, err
		}
		musics = append(musics, &music)
	}

	return &pb.Musics{Musics: musics}, nil
}

func (s *server) GetMusic(ctx context.Context, in *pb.MusicId) (*pb.Music, error) {
	var music pb.Music
	err := s.db.QueryRow("SELECT id, name, artist FROM musics WHERE id=$1", in.Id).Scan(&music.Id, &music.Name, &music.Artist)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("music not found")
		}
		return nil, err
	}
	return &music, nil
}

func (s *server) UpdateMusic(ctx context.Context, in *pb.Music) (*pb.Music, error) {
	_, err := s.db.Exec("UPDATE musics SET name=$1, artist=$2 WHERE id=$3", in.Name, in.Artist, in.Id)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *server) DeleteMusic(ctx context.Context, in *pb.MusicId) (*pb.Empty, error) {
	_, err := s.db.Exec("DELETE FROM musics WHERE id=$1", in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *server) CreatePlaylist(ctx context.Context, in *pb.Playlist) (*pb.Playlist, error) {
	err := s.db.QueryRow("INSERT INTO playlists(name, user_id) VALUES($1, $2) RETURNING id", in.Name, in.UserId).Scan(&in.Id)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *server) GetPlaylists(ctx context.Context, in *pb.Empty) (*pb.Playlists, error) {
	rows, err := s.db.Query("SELECT id, name, user_id FROM playlists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlists []*pb.Playlist
	for rows.Next() {
		var playlist pb.Playlist
		if err := rows.Scan(&playlist.Id, &playlist.Name, &playlist.UserId); err != nil {
			return nil, err
		}
		playlists = append(playlists, &playlist)
	}

	return &pb.Playlists{Playlists: playlists}, nil
}

func (s *server) GetPlaylist(ctx context.Context, in *pb.PlaylistId) (*pb.Playlist, error) {
	var playlist pb.Playlist
	err := s.db.QueryRow("SELECT id, name, user_id FROM playlists WHERE id=$1", in.Id).Scan(&playlist.Id, &playlist.Name, &playlist.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("playlist not found")
		}
		return nil, err
	}
	return &playlist, nil
}

func (s *server) UpdatePlaylist(ctx context.Context, in *pb.Playlist) (*pb.Playlist, error) {
	_, err := s.db.Exec("UPDATE playlists SET name=$1, user_id=$2 WHERE id=$3", in.Name, in.UserId, in.Id)
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (s *server) DeletePlaylist(ctx context.Context, in *pb.PlaylistId) (*pb.Empty, error) {
	_, err := s.db.Exec("DELETE FROM playlists WHERE id=$1", in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func main() {
	db := database.Init()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMusicServiceServer(s, &server{db: db})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
