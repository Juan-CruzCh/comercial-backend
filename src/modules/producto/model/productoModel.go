package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductoModel struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Codigo       string        `bson:"codigo"`
	Nombre       string        `bson:"nombre"`
	Descripcion  string        `bson:"descripcion,omitempty"`
	Categoria    bson.ObjectID `bson:"categoria"`
	UnidadManejo bson.ObjectID `bson:"unidadManejo"`
	Fecha        time.Time     `bson:"fecha"`
	Flag         enum.Estado   `bson:"flag"`
}
