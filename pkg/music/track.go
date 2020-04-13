package music

import (
	"context"
	"github.com/Ressetkk/nino/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddTrack(track Track) (*primitive.ObjectID, error) {
	col := db.GetCollection(tracks)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	res, err := col.InsertOne(ctx, track)
	if err != nil {
		return nil, err
	}
	insertedID := res.InsertedID.(primitive.ObjectID)
	return &insertedID, nil
}
