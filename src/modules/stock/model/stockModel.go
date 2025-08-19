package model

import (
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
	MontoTotal       float64       `bson:"montoTotal"`
	Fecha            time.Time     `bson:"fecha"`
	Flag             string        `bson:"flag"`
}
