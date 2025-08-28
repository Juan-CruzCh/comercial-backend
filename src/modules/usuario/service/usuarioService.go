package service

import (
	"comercial-backend/src/modules/usuario/dto"
	"comercial-backend/src/modules/usuario/repository"
	"context"
	"fmt"
)

func CrearUsuarioService(u *dto.UsuarioDto, ctx context.Context) error{
	usuario, err:=repository.VeficarUsuarioRepository(&u.Username, ctx)
	if err != nil {
		fmt.Println(err)
		return  err
	}
	fmt.Println("usuario",usuario)
	return  nil
}

func listarUsuarioService(ctx context.Context) {

}

func obtenerUsuarioService(id string, ctx context.Context) {

}
func actualizarUsuarioService(id string, u *dto.UsuarioDto, ctx context.Context) {

}

func eliminarUsuarioService(id string, ctx context.Context) {

}
