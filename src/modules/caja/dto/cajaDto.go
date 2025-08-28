package dto

type CajaDto struct {
	MontoInicial float64 `json:"montoInicial" binding:"required"`
}
