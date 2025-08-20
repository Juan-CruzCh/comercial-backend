package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleVentaModel struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	Producto       bson.ObjectID `bson:"producto,omitempty"`
	Cantidad       int           `bson:"cantidad"`
	PrecioUnitario float64       `bson:"precioUnitario"`
	preciototal    float64       `bson:"preciototal"`
	Flag           bool          `bson:"flag"`
	Fecha          time.Time     `bson:"fecha"`
	Descripcion    string        `bson:"descripcion"`
}
