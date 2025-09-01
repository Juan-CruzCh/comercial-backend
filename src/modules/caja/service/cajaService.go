package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/caja/dto"
	"comercial-backend/src/modules/caja/model"
	"comercial-backend/src/modules/caja/repository"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func AbriCajaService(body *dto.CajaDto, ctx context.Context, usuarioID *bson.ObjectID) error {
	fecha := utils.FechaHoraBolivia()

	err := repository.VerificarCajaAbierto(usuarioID, ctx)

	if err != nil {
		return err
	}
	var model model.CajaModel = model.CajaModel{
		MontoInicial:  body.MontoInicial,
		FechaApertura: fecha,
		Usuario:       *usuarioID,
		Flag:          enum.EstadoNuevo,
		Estado:        enum.Abierto,
		Fecha:         fecha,
	}
	err = repository.AbrirCajaRepository(&model, ctx)
	if err != nil {
		return err
	}
	return nil

}

func CerrarCajaService(usuario *bson.ObjectID, ctx context.Context) error {
	caja, err := repository.BuscarCajaUsuarioRepository(usuario, ctx)
	if err != nil {
		return errors.New("No existe niguna caja abierda" + err.Error())
	}
	fmt.Println(caja)
	err = repository.CerrarCajaRepository(&caja.ID, ctx)
	if err != nil {
		return err
	}
	return nil
}

func VerificarCajaService(usuario *bson.ObjectID, ctx context.Context) error {
	err := repository.VerificarCajaRepository(usuario, ctx)
	if err != nil {
		return err
	}
	return nil
}
