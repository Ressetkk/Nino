package music

import (
	"context"
	"encoding/json"
	"github.com/Ressetkk/nino/internal/db"
	"github.com/clevergo/jsend"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Artist struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Bio        string             `json:"bio,omitempty" bson:"bio,omitempty"`
	PictureUrl string             `json:"picture_url" bson:"picture_url"`
}

//TODO add validation of fields
func addArtist(w http.ResponseWriter, r *http.Request) error {
	var newArtist Artist
	if err := json.NewDecoder(r.Body).Decode(&newArtist); err != nil {
		return err
	}
	col := db.GetCollection(artists)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	id, err := col.InsertOne(ctx, newArtist)
	if err != nil {
		return err
	}
	return jsend.Success(w, id)
}

func getArtists(w http.ResponseWriter, r *http.Request) error {
	col := db.GetCollection(artists)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var a []primitive.M
	for cur.Next(ctx) {
		var elem primitive.M
		if err := cur.Decode(&elem); err != nil {
			return err
		}
		a = append(a, elem)
	}
	return jsend.Success(w, a)
}
