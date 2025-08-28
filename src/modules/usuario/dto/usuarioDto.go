package dto

type UsuarioDto struct {
	CI        string `json:"ci" validate:"required"`
	Nombre    string `json:"nombre" validate:"required"`
	Apellidos string `json:"apellidos" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Sucursal  string `json:"sucursal" validate:"required"`
	Rol       string `json:"rol" validate:"required"`
}
