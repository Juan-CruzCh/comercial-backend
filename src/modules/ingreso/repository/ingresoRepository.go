package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/modules/ingreso/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CountDocumentsIngresoRepository(ctx context.Context) (int64, error) {
	collection := config.MongoDatabase.Collection("Ingreso")
	countDocuments, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return countDocuments, nil
}

func CrearIngresoRepository(data *model.IngresoModel, ctx context.Context) (*bson.ObjectID, error) {
	collection := config.MongoDatabase.Collection("Ingreso")
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return &bson.NilObjectID, err
	}
	ingresoID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return &bson.NilObjectID, errors.New("ocurrio un error al insertar el ingreso")
	}
	return &ingresoID, nil
}
