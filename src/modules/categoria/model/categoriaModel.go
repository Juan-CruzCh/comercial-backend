package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categoria struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Nombre string             `bson:"nombre"`
	Fecha  time.Time          `bson:"fecha"`
	Flag   string              `bson:"flag"`
}