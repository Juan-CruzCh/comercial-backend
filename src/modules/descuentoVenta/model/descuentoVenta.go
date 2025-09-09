package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DescuentoVenta struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Alquiler float64       `bson:"alquiler"`
	Vendedor float64       `bson:"vendedor"`
	Sucursal bson.ObjectID `bson:"sucursal"`
	Fecha    time.Time     `bson:"fecha"`
	Flag     enum.Estado   `bson:"flag"`
}
