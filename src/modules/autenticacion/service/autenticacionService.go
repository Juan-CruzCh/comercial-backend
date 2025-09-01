package service

import (
	"comercial-backend/src/modules/autenticacion/dto"
	"comercial-backend/src/modules/autenticacion/utils"
	"comercial-backend/src/modules/usuario/repository"
	usuarioUtil "comercial-backend/src/modules/usuario/utils"
	"context"
	"errors"
	"fmt"
)

func AutenticacionService(dto *dto.AutenticacionDto, cxt context.Context) (string, error) {
	usuario, err := repository.VeficarUsuarioRepository(&dto.Username, cxt)
	if err != nil {

		return "", errors.New("Contraseña invalida")
	}
	ok, err := usuarioUtil.ComparePasswordAndHash(dto.Password, usuario.Password)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if !ok {
		return "", errors.New("Contraseña invalida")
	}
	token, err := utils.GenraraToken(usuario.ID, usuario.Sucursal)
	if err != nil {
		return "", err
	}

	return token, nil

}
