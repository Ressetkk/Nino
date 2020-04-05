package music

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ressetkk/nino/internal/db"
	"github.com/Ressetkk/nino/internal/helpers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"net/http"
)

type Album struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	ReleaseDate  primitive.DateTime `json:"release_date" bson:"release_date"`
	Path         string             `json:"path" bson:"path"`
	AlbumArtPath string             `json:"album_art_path" bson:"album_art_path"`
	NumOfTracks  int                `json:"num_of_tracks" bson:"num_of_tracks"`
	Tracks       []*Track           `json:"tracks" bson:"tracks"`
	Artist       *Artist            `json:"artist" bson:"artist"`
}

type Track struct {
	Title string `json:"title" bson:"title"`
	Path  string `json:"path" bson:"path"`
}

type Artist struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

func AddRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/music").Subrouter()
	s.HandleFunc("/albums", getAlbums).Methods("GET")
	s.HandleFunc("/albums", addAlbum).Methods("POST")
	s.HandleFunc("/albums/{id:[0-9]+}", getAlbumInfo).Methods("GET")
	s.HandleFunc("/albums/albums/{id:[0-9]+}", editAlbum).Methods("PUT", "DELETE")
	s.HandleFunc("/albums/{id:[0-9]+}/tracks", getAlbumTracks).Methods("GET")
	return s
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := db.GetCollection("music")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Error(helpers.WriteErrorResponse(w, http.StatusBadRequest, err))
		return
	}
	defer func() {
		if err := cur.Close(context.TODO()); err != nil {
			log.Error(err)
		}
	}()

	var albums []*Album
	for cur.Next(context.TODO()) {
		var elem Album
		if err := cur.Decode(&elem); err != nil {
			log.Error(err)
		}
		albums = append(albums, &elem)
	}
	resp, err := json.Marshal(albums)
	if err != nil {
		log.Error(helpers.WriteErrorResponse(w, http.StatusInternalServerError, err))
		return
	}
	w.Write(resp)
}

func addAlbum(w http.ResponseWriter, r *http.Request) {

}

func getAlbumInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	io.WriteString(w, id)
}

func editAlbum(w http.ResponseWriter, r *http.Request) {

}

func getAlbumTracks(w http.ResponseWriter, r *http.Request) {

}
