package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type SucursalModel struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Nombre    string        `bson:"nombre"`
	Direccion string        `bson:"direccion"`
	Fecha     time.Time     `bson:"fecha"`
	Flag      enum.Estado        `bson:"flag"`
}
