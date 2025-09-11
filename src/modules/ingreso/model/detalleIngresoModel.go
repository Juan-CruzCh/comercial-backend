package model

import (
	"comercial-backend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleIngresoModel struct {
	ID                        bson.ObjectID `bson:"_id,omitempty"`
	CodigoStock               string        `bson:"codigoStock"`
	Producto                  bson.ObjectID `bson:"producto"`
	Ingreso                   bson.ObjectID `bson:"ingreso"`
	Cantidad                  int           `bson:"cantidad"`
	Fecha                     time.Time     `bson:"fecha"`
	PrecioUnitario            float64       `bson:"precioUnitario"`
	PrecioUnitarioCompra      float64       `bson:"precioUnitarioCompra"`
	FechaVencimiento          *time.Time    `bson:"fechaVencimiento,omitempty"`
	PrecioUnitarioTotal       float64       `bson:"precioUnitarioTotal"`
	PrecioUnitarioTotalCompra float64       `bson:"precioUnitarioTotalCompra"`
	Descuento                 float64       `bson:"descuento"`
	SubTotal                  float64       `bson:"subTotal"`
	Flag                      enum.Estado   `bson:"flag"`
}
