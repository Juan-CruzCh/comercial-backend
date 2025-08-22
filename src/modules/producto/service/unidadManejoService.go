package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"strings"

	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/model"
	"comercial-backend/src/modules/producto/repository"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CrearUnidadManejoService(unidadManejoDto *dto.UnidadManejoDto, ctx context.Context) error {

	fecha := utils.FechaHoraBolivia()

	var data model.UnidadManejoModel = model.UnidadManejoModel{
		Nombre: strings.ToUpper(unidadManejoDto.Nombre),
		Fecha:  fecha,
		Flag:   enum.EstadoNuevo,
	}

	err := repository.CrearUnidadManejoRepository(data, ctx)
	if err != nil {

		return err
	}

	return err
}

func ObtenerUnidadManejoPorIDService(ctx context.Context) {

}

func ActualizarUnidadManejoService(ctx context.Context, id string) {

}

func EliminarUnidadManejoService(ctx context.Context, id string) {

}

func ListarUnidadManejoService(ctx context.Context) ([]bson.M, error) {

	data, err := repository.ListarUnidadManejoRepository(ctx)
	if err != nil {

		return []bson.M{}, err
	}
	return data, nil

}
