package dto

type VentaDto struct {
	MontoTotal float64 `json:"montoTotal"  binding:"gte=0"`
}

type DetalleVenta struct {
	Producto    string  `json:"producto"   binding:"required"`
	Stock       string  `json:"stock"  binding:"gte=0"`
	Cantidad    int     `json:"cantidad"  binding:"gte=0"`
	Preciototal float64 `json:"precioTotal"  binding:"gte=0"`
}
