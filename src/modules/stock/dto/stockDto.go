package dto

import (
	"time"
)

type IngresoStockDto struct {
	Proveedor string               `json:"proveedor"  binding:"required"`
	Factura   string               `json:"factura" binding:"required"`
	Stock     []StockDtoDetalleDto `json:"stock"  binding:"required,dive,required"`
}

type StockDtoDetalleDto struct {
	Cantidad             int       `json:"cantidad" binding:"required,gt=0"`
	FechaVencimiento     time.Time `json:"fechaVencimiento" binding:"omitempty"`
	Producto             string    `json:"producto"  binding:"required"`
	PrecioUnitario       float64   `json:"precioUnitario" binding:"gte=0"`
	PrecioUnitarioCompra float64   `json:"precioUnitarioCompra" binding:"gte=0"`
	Descuento            float64   `json:"descuento" binding:"gte=0"`
}
