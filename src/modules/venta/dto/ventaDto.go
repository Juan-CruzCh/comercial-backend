package dto

type VentaDto struct {
	MontoTotal   float64        `json:"montoTotal" binding:"required,gte=0"`
	Descuento    *float64       `json:"descuento" binding:"required,gte=0"`
	SudTotal     float64        `json:"sudTotal" binding:"required,gte=0"`
	DetalleVenta []DetalleVenta `json:"detalleVenta" binding:"required,dive,required"`
}

type DetalleVenta struct {
	Stock               string  `json:"stock" binding:"required"`
	Cantidad            int     `json:"cantidad" binding:"required,gte=0"`
	PrecioTotal         float64 `json:"precioTotal" binding:"required,gte=0"`
	PrecioUnitario      float64 `json:"precioUnitario" binding:"required,gte=0"`
	DescripcionProducto string  `json:"descripcionProducto" binding:"required"`
}
