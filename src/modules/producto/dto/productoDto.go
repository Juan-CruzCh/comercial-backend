package dto

type ProductoDto struct {
	Nombre       string `json:"nombre" binding:"required"`
	Descripcion  string `json:"descripcion"`
	Categoria    string `json:"categoria" binding:"required"`
	UnidadManejo string `json:"unidadManejo" binding:"required"`
}
