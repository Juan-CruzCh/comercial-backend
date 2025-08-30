package structIngreso

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type IngresoStockData struct {
	Proveedor  bson.ObjectID
	Usuario    bson.ObjectID
	Factura    string
	MontoTotal float64
	Stock      []StockDto
}

type StockDto struct {
	Cantidad         int
	FechaVencimiento time.Time
	Producto         bson.ObjectID
	PrecioUnitario   float64
	MontoTotal       float64
	Descuento        float64
	SudTotal         float64
}
