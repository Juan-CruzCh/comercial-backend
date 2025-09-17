package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/sucursal/model"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CrearSucursalRepository(data *model.SucursalModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Sucursal)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil

}

func ListarSucursalRepository(ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Sucursal)
	cursor, err := collection.Find(ctx, bson.M{"flag": enum.EstadoNuevo})
	if err != nil {
		return &[]bson.M{}, err
	}
	var resultado []bson.M
	err = cursor.All(ctx, &resultado)
	if err != nil {
		return &[]bson.M{}, err
	}
	defer cursor.Close(ctx)
	return &resultado, nil

}

func EliminarSucursalRepository(id *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Sucursal)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"flag": enum.EstadoEliminado}})
	if err != nil {
		return err
	}
	return nil
}
