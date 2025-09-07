package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CajaModel struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	FechaApertura time.Time     `bson:"fechaApertura"`
	MontoInicial  float64       `bson:"montoInicial"`
	FechaCierre   time.Time     `bson:"fechaCierre,omitempty"`
	MontoFinal    float64       `bson:"montoFinal"`
	TotalVentas   float64       `bson:"totalVentas"`
	Egresos       float64       `bson:"egresos"`
	Estado        enum.CajaEnum `bson:"estado"`
	Usuario       bson.ObjectID `bson:"usuario"`
	Sucursal      bson.ObjectID `bson:"sucursal"`
	Fecha         time.Time     `bson:"fecha"`
	Flag          enum.Estado   `bson:"flag"`
}
