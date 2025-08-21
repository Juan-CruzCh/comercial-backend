package service

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"time"

	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/model"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
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

func ListarUnidadManejoService(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection("UnidadManejo")
	var resultado []bson.M
	cursor, err := collection.Find(ctx, bson.M{"flag": enum.EstadoNuevo})
	if err != nil {

		return []bson.M{}, err
	}
	err = cursor.All(ctx, &resultado)
	if err != nil {
		return []bson.M{}, err
	}
	return resultado, nil
}
