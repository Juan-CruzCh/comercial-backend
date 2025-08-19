package dto

import "time"

type StockDto struct {
	Cantidad         int       `json:"cantidad"`
	FechaVencimiento time.Time `json:"fechaVencimiento"`
	Producto         string    `json:"producto"`
	PrecioUnitario   float64   `json:"precioUnitario"`
	MontoTotal       float64   `json:"montoTotal"`
}
