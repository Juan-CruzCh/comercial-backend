package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/venta/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func CountDocumentsVentaRepository(ctx context.Context) (int64, error) {
	collection := config.MongoDatabase.Collection(enum.Venta)
	countDocuments, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return countDocuments, nil
}

func ListarVentasRepository(ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Venta)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "codigo", Value: 1},
				{Key: "montoTotal", Value: 1},
				{Key: "subTotal", Value: 1},
				{Key: "fechaVenta", Value: 1},
				{Key: "descuento", Value: 1},
			}},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
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
