package dto

import (
	"time"
)

type IngresoStockDto struct {
	Proveedor  string               `json:"proveedor"  binding:"required"`
	Factura    string               `json:"factura" binding:"required"`
	MontoTotal float64              `json:"montoTotal"  binding:"gte=0"`
	Stock      []StockDtoDetalleDto `json:"stock"  binding:"required,dive,required"`
}

type StockDtoDetalleDto struct {
	Cantidad         int       `json:"cantidad" binding:"required,gt=0"`
	FechaVencimiento time.Time `json:"fechaVencimiento" binding:"omitempty"`
	Producto         string    `json:"producto"  binding:"required"`
	PrecioUnitario   float64   `json:"precioUnitario" binding:"gte=0"`
	MontoTotal       float64   `json:"montoTotal" binding:"gte=0"`
	Descuento        float64   `json:"descuento" binding:"gte=0"`
	SudTotal         float64   `json:"sudTotal" binding:"gte=0"`
}
