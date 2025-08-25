package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/venta/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RealizarVentaRepository(venta *model.VentaModel, ctx context.Context) (*bson.ObjectID, error) {
	collection := config.MongoDatabase.Collection(enum.Venta)
	resultado, err := collection.InsertOne(ctx, venta)
	if err != nil {
		return &bson.NilObjectID, err
	}
	ventaID, ok := resultado.InsertedID.(bson.ObjectID)
	if !ok {
		return &bson.NilObjectID, errors.New("Se prodcuto un erro al ingresa la venta")
	}

	return &ventaID, nil

}
