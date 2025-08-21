package dto

type ProveedorDto struct {
	CI        string `json:"ci" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellidos string `json:"apellidos" binding:"required"`
	Celular   string `json:"celular" binding:"required"`
	Empresa   string `json:"empresa"`
}