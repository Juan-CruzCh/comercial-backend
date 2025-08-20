package dto

type ProveedorDto struct {
	CI        string `json:"ci" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellidos string `json:"apellidos" binding:"required"`
	Empresa   string `json:"empresa"`
}