package music

import (
	"context"
	"github.com/Ressetkk/nino/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	albums    = "musicAlbums"
	albumInfo = "albumInfo"
)

type Album struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	ReleaseDate string             `json:"release_date" bson:"release_date"`
	Path        string             `json:"path" bson:"path"`
	AlbumArtUrl string             `json:"album_art_url" bson:"album_art_url"`
	NumOfTracks int                `json:"num_of_tracks" bson:"num_of_tracks"`
	Artist      string             `json:"artist_id" bson:"artist_id"`
}

func AddAlbum(album Album) (*primitive.ObjectID, error) {
	col := db.GetCollection(albums)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	res, err := col.InsertOne(ctx, album)
	if err != nil {
		return nil, err
	}
	insertedID := res.InsertedID.(primitive.ObjectID)
	return &insertedID, nil
}

func GetAlbum(id primitive.ObjectID) (*Album, error) {
	col := db.GetCollection(albums)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	res := col.FindOne(ctx, bson.M{"_id": id})
	var a Album
	if err := res.Decode(&a); err != nil {
		return nil, err
	}
	return &a, nil
}
