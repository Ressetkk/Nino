package music

import (
	"context"
	"encoding/json"
	"github.com/Ressetkk/nino/internal/db"
	"github.com/clevergo/jsend"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	s.HandleFunc("/albums/{id:[a-f0-9]+}", getAlbumInfo).Methods("GET")
	s.HandleFunc("/albums/albums/{id:[a-f0-9]+}", editAlbum).Methods("PATCH", "DELETE")
	s.HandleFunc("/albums/{id:[a-f0-9]+}/tracks", getAlbumTracks).Methods("GET")
	return s
}

// TODO write helper methods for error handling
// TODO write context with timeout for timing out on requests
func getAlbums(w http.ResponseWriter, r *http.Request) {
	collection := db.GetCollection("music")
	cur, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		if err := jsend.Error(w, err.Error(), http.StatusBadRequest); err != nil {
			log.Error(err)
		}
		return
	}
	defer db.Close(cur, context.TODO())
	var albums []*Album
	for cur.Next(context.TODO()) {
		var elem Album
		if err := cur.Decode(&elem); err != nil {
			log.Error(err)
		}
		albums = append(albums, &elem)
	}

	if err := jsend.Success(w, albums); err != nil {
		log.Error(err)
	}
}

func addAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		log.Error(err)
	}

}

func getAlbumInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	collection := db.GetCollection("music")
	obj := collection.FindOne(context.TODO(), filter, options.FindOne())
	var album Album
	if err := obj.Decode(&album); err != nil {
		log.Error(err)
		if err := jsend.Error(w, err.Error(), http.StatusBadRequest); err != nil {
			log.Error(err)
		}
		return
	}
	if err := jsend.Success(w, album); err != nil {
		log.Error(err)
	}
}

func editAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {

	}
}

func getAlbumTracks(w http.ResponseWriter, r *http.Request) {

}
