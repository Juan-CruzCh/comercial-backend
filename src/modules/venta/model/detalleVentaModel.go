package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleVentaModel struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Producto    bson.ObjectID `bson:"producto"`
	Stock       bson.ObjectID `bson:"stock"`
	Venta       bson.ObjectID `bson:"venta"`
	Cantidad    int           `bson:"cantidad"`
	PrecioTotal float64       `bson:"precioTotal"`
	PrecioUnitario float64 `bson:"precioUnitario"`
	Flag        enum.Estado   `bson:"flag"`
	Fecha       time.Time     `bson:"fecha"`
	Descripcion string        `bson:"descripcion"`
}
