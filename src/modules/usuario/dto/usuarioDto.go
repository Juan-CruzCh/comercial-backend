package dto

type UsuarioDto struct {
	CI        string `json:"ci" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellidos string `json:"apellidos" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Sucursal  string `json:"sucursal" binding:"required"`
	Rol       string `json:"rol" binding:"required"`
}

type ActualizarUsuarioDto struct {
	CI        string `json:"ci" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Apellidos string `json:"apellidos" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Sucursal  string `json:"sucursal" binding:"required"`
	Rol       string `json:"rol" binding:"required"`
}
