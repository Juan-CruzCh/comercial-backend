package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AlmacenModel struct {
	ID     bson.ObjectID `bson:"_id,omitempty"`
	Nombre string        `bson:"nombre"`
	Fecha  time.Time     `bson:"fecha"`
	Flag   bool          `bson:"flag"`
}
