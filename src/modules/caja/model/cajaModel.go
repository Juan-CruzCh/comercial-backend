package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CajaModel struct {
    ID                 bson.ObjectID `bson:"_id,omitempty"`
    FechaApertura      time.Time          `bson:"fecha_apertura"`
    MontoInicial       float64            `bson:"monto_inicial"`
    FechaCierre        *time.Time         `bson:"fecha_cierre,omitempty"`
    MontoFinal         float64            `bson:"monto_final"`
    TotalVentas        float64            `bson:"total_ventas"`
    Egresos            float64            `bson:"egresos"`
    Estado             string             `bson:"estado"` 
    Usuario         bson.ObjectID      `bson:"usuario"`
	Fecha        *time.Time         `bson:"fecha"`
	Flag        enum.Estado         `bson:"fecha"`
}
