package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleIngresoModel struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	Producto       bson.ObjectID `bson:"producto"`
	Cantidad       int           `bson:"cantidad"`
	Fecha          time.Time     `bson:"fecha"`
	PrecioUnitario float64       `bson:"precio_unitario"`
	Flag           enum.Estado   `bson:"flag"`
}
