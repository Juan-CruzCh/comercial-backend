package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleIngresoModel struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	Producto         bson.ObjectID `bson:"producto"`
	Ingreso          bson.ObjectID `bson:"ingreso"`
	Cantidad         int           `bson:"cantidad"`
	Fecha            time.Time     `bson:"fecha"`
	PrecioUnitario   float64       `bson:"precioUnitario"`
	FechaVencimiento *time.Time    `bson:"fechaVencimiento,omitempty"`
	MontoTotal       float64       `bson:"montoTotal"`
	Descuento        float64       `bson:"descuento"`
	SudTotal         float64       `bson:"sudTotal"`
	Flag             enum.Estado   `bson:"flag"`
}
