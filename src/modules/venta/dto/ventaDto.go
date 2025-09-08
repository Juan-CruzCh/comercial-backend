package dto

type VentaDto struct {
	Descuento    *float64       `json:"descuento" binding:"required,gte=0"`
	DetalleVenta []DetalleVenta `json:"detalleVenta" binding:"required,dive,required"`
}

type DetalleVenta struct {
	Stock               string  `json:"stock" binding:"required"`
	Cantidad            int     `json:"cantidad" binding:"required,gte=0"`
	PrecioUnitario      float64 `json:"precioUnitario" binding:"required,gte=0"`
	DescripcionProducto string  `json:"descripcionProducto" binding:"required"`
}

type BuscadorVentaDto struct {
	Codigo      string `json:"codigo"`
	Sucursal    string `json:"sucursal"`
	Usuario     string `json:"usuario"`
	FechaInicio string `json:"fechaInicio"`
	FechaFin    string `json:"fechaFin"`
}
