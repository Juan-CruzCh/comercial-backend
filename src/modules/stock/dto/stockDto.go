package dto

import (
	"time"
)

type IngresoStockData struct {
	Proveedor  string     `json:"proveedor"`
	Factura    string     `json:"factura"`
	MontoTotal float64    `json:"montoTotal"`
	Stock      []StockDto `json:"stock"`
}

type StockDto struct {
	Cantidad         int       `json:"cantidad"`
	FechaVencimiento time.Time `json:"fechaVencimiento"`
	Producto         string    `json:"producto"`
	PrecioUnitario   float64   `json:"precioUnitario"`
	MontoTotal       float64   `json:"montoTotal"`
	Descuento        float64   `json:"descuento"`
	SudTotal         float64   `json:"sudTotal"`
}
