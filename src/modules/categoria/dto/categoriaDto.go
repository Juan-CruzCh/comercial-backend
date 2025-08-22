package dto

type CategoriaDto struct {
	Nombre string `json:"nombre"  binding:"required"`
}
