package dto

type UsuarioDto struct {
	CI       string `json:"ci"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sucursal string `json:"sucursal"`
	Rol      string `json:"rol"`
}
