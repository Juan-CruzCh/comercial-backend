package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleVentaModel struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	Producto       bson.ObjectID `bson:"producto"`
	Venta          bson.ObjectID `bson:"venta"`
	Cantidad       int           `bson:"cantidad"`
	PrecioUnitario float64       `bson:"precioUnitario"`
	preciototal    float64       `bson:"preciototal"`
	Flag           enum.Estado   `bson:"flag"`
	Fecha          time.Time     `bson:"fecha"`
	Descripcion    string        `bson:"descripcion"`
}
