package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/usuario/dto"
	"comercial-backend/src/modules/usuario/model"
	"comercial-backend/src/modules/usuario/repository"
	utilsUsuario "comercial-backend/src/modules/usuario/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CrearUsuarioService(u *dto.UsuarioDto, ctx context.Context) error {
	_, err := repository.VeficarUsuarioExisteRepository(&u.Username, ctx)
	if err != nil {
		return err
	}
	hash, err := utilsUsuario.EncriptarPassword(u.Password)
	sucurcalID, err := utils.ValidadIdMongo(u.Sucursal)
	if err != nil {

		return err
	}
	var user model.UsuarioModel = model.UsuarioModel{
		CI:        u.CI,
		Nombre:    u.Nombre,
		Apellidos: u.Apellidos,
		Username:  u.Username,
		Password:  hash,
		Sucursal:  *sucurcalID,
		Rol:       u.Rol,
		Fecha:     time.Now(),
		Flag:      enum.EstadoNuevo,
	}
	err = repository.CrearUsuarioRepository(&user, ctx)
	if err != nil {
		return err
	}
	return nil
}

func ListarUsuarioService(ctx context.Context) (*[]bson.M, error) {
	data, err := repository.ListarUsuarioRepository(ctx)
	if err != nil {
		return &[]bson.M{}, err
	}
	return data, nil
}

func obtenerUsuarioService(id string, ctx context.Context) {

}
func actualizarUsuarioService(id string, u *dto.UsuarioDto, ctx context.Context) {

}

func eliminarUsuarioService(id string, ctx context.Context) {

}
