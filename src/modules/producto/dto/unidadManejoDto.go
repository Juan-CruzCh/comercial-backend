package dto

type UnidadManejoDto struct {
	Nombre string `json:"nombre" binding:"required"`
}
