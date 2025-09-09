package dto

type DescuentoVentaDto struct {
	Alquiler float64 `json:"alquiler" binding:"required,gte=0"`
	Vendedor float64 `json:"vendedor" binding:"required,gte=0"`
	Sucursal string  `json:"sucursal" binding:"required"`
}
