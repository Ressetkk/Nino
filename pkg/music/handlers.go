package music

import (
	"context"
	"encoding/json"
	"github.com/Ressetkk/nino/internal/db"
	"github.com/Ressetkk/nino/pkg/logging"
	"github.com/clevergo/jsend"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

const (
	artists = "musicArtists"
	tracks  = "musicTracks"
)

type Track struct {
	Title       string             `json:"title" bson:"title"`
	Path        string             `json:"path" bson:"path"`
	TrackNumber int                `json:"track_number"`
	Album       primitive.ObjectID `json:"album_id" bson:"album_id"`
}

func AddRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/music").Subrouter()
	s.Handle("/albums", logging.ErrorHandler(getAlbums)).Methods("GET")
	s.Handle("/albums", logging.ErrorHandler(addAlbum)).Methods("POST")
	s.Handle("/albums/{id:[a-f0-9]+}", logging.ErrorHandler(getAlbumInfo)).Methods("GET")
	s.Handle("/albums/{id:[a-f0-9]+}", logging.ErrorHandler(editAlbum)).Methods("PATCH", "DELETE")
	s.Handle("/albums/{id:[a-f0-9]+}/tracks", logging.ErrorHandler(getAlbumTracks)).Methods("GET")

	s.Handle("/artists", logging.ErrorHandler(getArtists)).Methods("GET")
	s.Handle("/artists", logging.ErrorHandler(addArtist)).Methods("POST")
	//s.HandleFunc("/artists/{id:[a-f0-9]+}", getArtistInfo).Methods("GET")
	//s.HandleFunc("/artists/{id:[a-f0-9]+}/albums", getArtistAlbums).Methods("POST")
	return s
}

// TODO write context with timeout for timing out on requests
func getAlbums(w http.ResponseWriter, r *http.Request) error {
	col := db.GetCollection(albumInfo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var albums []primitive.M
	for cur.Next(ctx) {
		var album primitive.M
		if err := cur.Decode(&album); err != nil {
			return err
		}
		albums = append(albums, album)
	}
	return jsend.Success(w, albums)
}

func addAlbum(w http.ResponseWriter, r *http.Request) error {
	var newAlbum Album
	err := json.NewDecoder(r.Body).Decode(&newAlbum)
	if err != nil {
		return err
	}
	insertedID, err := AddAlbum(newAlbum)
	if err != nil {
		return err
	}
	album, err := GetAlbum(*insertedID)
	if err != nil {
		return err
	}
	return jsend.Success(w, album)
}

func getAlbumInfo(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	objId, _ := primitive.ObjectIDFromHex(params["id"])
	collection := db.GetCollection(albumInfo)
	result := collection.FindOne(context.TODO(), bson.M{"_id": objId})
	var album primitive.M
	if err := result.Decode(&album); err != nil {
		return err
	}
	return jsend.Success(w, album)
}

func editAlbum(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "PATCH" {

	}
	return nil
}

func getAlbumTracks(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	objId, _ := primitive.ObjectIDFromHex(params["id"])
	col := db.GetCollection(tracks)
	cur, err := col.Find(context.TODO(), bson.M{"album_id": objId})
	if err != nil {
		return err
	}
	var elms []Track
	for cur.Next(context.TODO()) {
		var t Track
		if err := cur.Decode(&t); err != nil {
			return err
		}
		elms = append(elms, t)
	}
	return jsend.Success(w, elms)
}
