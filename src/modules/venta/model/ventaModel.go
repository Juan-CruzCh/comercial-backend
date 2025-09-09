package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type VentaModel struct {
	ID                 bson.ObjectID    `bson:"_id,omitempty"`
	Codigo             string           `bson:"codigo"`
	MontoTotal         float64          `bson:"montoTotal"`
	Descuento          float64          `bson:"descuento"`
	SubTotal           float64          `bson:"subTotal"`
	FechaVenta         time.Time        `bson:"fechaVenta"`
	Fecha              time.Time        `bson:"fecha"`
	Flag               enum.Estado      `bson:"flag"`
	Estado             enum.EstadoVenta `bson:"estado"`
	TipoPago           enum.TipoPago    `bson:"tipoPago"`
	Usuario            bson.ObjectID    `bson:"usuario"`
	Sucursal           bson.ObjectID    `bson:"sucursal"`
	DescuentoAlquiller float64          `bson:"descuentoAlquiller"`
	DescuentoVendedor  float64          `bson:"descuentoVendedor"`
	TotalGanancia      float64          `bson:"totalGanancia"`
}
