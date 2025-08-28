package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/caja/dto"
	"comercial-backend/src/modules/caja/model"
	"comercial-backend/src/modules/caja/repository"
	"context"
)

func AbriCajaService(body *dto.CajaDto, ctx context.Context) error {
	fecha := utils.FechaHoraBolivia()
	usuarioID, err := utils.ValidadIdMongo("68b06561b72e50f06889d3ee")
	if err != nil {
		return err
	}
	err = repository.VerificarCajaAbierto(usuarioID, ctx)

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
	repository.AbrirCajaRepository(&model, ctx)
	return nil

}
