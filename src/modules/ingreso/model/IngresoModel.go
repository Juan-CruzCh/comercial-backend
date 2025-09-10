package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IngresoModel struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	Codigo         string        `bson:"codigo"`
	Fecha          time.Time     `bson:"fecha"`
	Proveedor      bson.ObjectID `bson:"proveedor"`
	Usuario        bson.ObjectID `bson:"usuario"`
	Factura        string        `bson:"factura"`
	MontoTotal     float64       `bson:"montoTotal"`
	TotalDescuento float64       `bson:"totalDescuento"`
	SudTotal       float64       `bson:"sudTotal"`
	Flag           enum.Estado   `bson:"flag"`
}
