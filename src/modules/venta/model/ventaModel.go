package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type VentaModel struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Codigo     string        `bson:"codigo"`
	Total      float64       `bson:"total"`
	FechaVenta time.Time     `bson:"fechaVenta"`
	Fecha      time.Time     `bson:"fecha"`
	Flag       enum.Estado   `bson:"flag"`
	Estado     string        `bson:"estado"`
	TipoPago   string        `bson:"tipoPago"`
	Usuario    bson.ObjectID `bson:"usuario"`
	Sucursal   bson.ObjectID `bson:"sucursal"`
}
