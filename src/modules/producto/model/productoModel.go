package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductoModel struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Codigo    string        `bson:"codigo"`
	Nombre    string        `bson:"nombre"`
	Descripcion string             `bson:"descripcion,omitempty"`
	Categoria bson.ObjectID `bson:"categoria"`
	Fecha     time.Time     `bson:"fecha"`
	Flag      string        `bson:"flag"`
}
