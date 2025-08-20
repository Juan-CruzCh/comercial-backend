package service

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"time"

	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/model"
	"context"
)

func CrearUnidadManejoService(unidadManejoDto *dto.UnidadManejoDto, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("UnidadManejo")
	var model = model.UnidadManejoModel{
		Nombre: unidadManejoDto.Nombre,
		Fecha:  time.Now(),
		Flag:   enum.EstadoNuevo,
	}
	_, err := collection.InsertOne(ctx, model)
	if err != nil {

		return err
	}
	return nil
}

func ObtenerUnidadManejoPorIDService(ctx context.Context) {

}

func ActualizarUnidadManejoService(ctx context.Context, id string) {

}

func EliminarUnidadManejoService(ctx context.Context, id string) {

}

func ListarUnidadManejoService(ctx context.Context) {

}
