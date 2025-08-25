package dto

type VentaDto struct {
	MontoTotal   float64        `json:"montoTotal"  binding:"gte=0"`
	Descuento    float64        `json:"descuento"  binding:"gte=0"`
	DetalleVenta []DetalleVenta `json:"detalleVenta"    binding:"required,dive"`
}

type DetalleVenta struct {
	Stock       string  `json:"stock"  binding:"required"`
	Cantidad    int     `json:"cantidad"  binding:"gte=0"`
	Preciototal float64 `json:"precioTotal"  binding:"gte=0"`
}
