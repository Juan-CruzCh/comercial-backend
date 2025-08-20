package dto

type AlmacenDto struct {
	Nombre string `json:"nombre"  binding:"required"`
}
