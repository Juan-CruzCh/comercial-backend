package dto

type SucursalDto struct {
	Nombre    string `json:"nombre"   binding:"required"`
	Direccion string `json:"direccion"   binding:"required"`
}
