package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/producto/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CrearUnidadManejoRepository(data model.UnidadManejoModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("UnidadManejo")

	countDocuments, err := collection.CountDocuments(ctx, bson.M{"nombre": data.Nombre, "flag": enum.EstadoNuevo})
	if err != nil {

		return err
	}
	if countDocuments > 0 {
		return errors.New("La unidad de manejo ya existe")
	}
	_, err = collection.InsertOne(ctx, data)
	if err != nil {

		return err
	}
	return nil
}

func ListarUnidadManejoRepository(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection("UnidadManejo")
	var resultado []bson.M
	cursor, err := collection.Find(ctx, bson.M{"flag": enum.EstadoNuevo})
	if err != nil {

		return []bson.M{}, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &resultado)
	if err != nil {
		return []bson.M{}, err
	}
	return resultado, nil
}
