package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type VentaModel struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Codigo     string        `bson:"codigo"`
	Total      float64       `bson:"total"`
	FechaVenta time.Time     `bson:"fechaVenta"`
	Fecha      time.Time     `bson:"fecha"`
	Flag       bool          `bson:"flag"`
	Estado     string        `bson:"estado"`
	TipoPago   string        `bson:"tipoPago"`
	Usuario    bson.ObjectID `bson:"usuario"`
	Sucursal   bson.ObjectID `bson:"sucursal"`
}
