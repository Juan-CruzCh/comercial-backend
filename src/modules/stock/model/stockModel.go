package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type StockModel struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	Cantidad         int           `bson:"cantidad"`
	FechaVencimiento time.Time     `bson:"fechaVencimiento"`
	Codigo           string        `bson:"codigo"`
	Producto         bson.ObjectID `bson:"producto"`
	PrecioUnitario   float64       `bson:"precioUnitario"`
	MontoTotal       float64       `bson:"MontoTotal"`
	Descuento        float64       `bson:"descuento"`
	SubTotal         float64       `bson:"subTotal"`
	Fecha            time.Time     `bson:"fecha"`
	Flag             enum.Estado   `bson:"flag"`
}
