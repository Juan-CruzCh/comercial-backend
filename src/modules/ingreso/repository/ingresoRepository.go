package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CountDocumentsIngresoRepository(ctx context.Context) (int64, error) {
	collection := config.MongoDatabase.Collection(enum.Ingreso)
	countDocuments, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return countDocuments, nil
}

func CrearIngresoRepository(data *model.IngresoModel, ctx context.Context) (*bson.ObjectID, error) {
	collection := config.MongoDatabase.Collection(enum.Ingreso)
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

func ListarIngresoRepository(ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Ingreso)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},

		utils.Lookup("Proveedor", "proveedor", "_id", "proveedor"),
		utils.Lookup("Usuario", "usuario", "_id", "usuario"),

		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "codigo", Value: 1},
				{Key: "fecha", Value: 1},
				{Key: "factura", Value: 1},
				{Key: "precioUnitarioTotal", Value: 1},
				{Key: "precioUnitarioTotalCompra", Value: 1},
				{Key: "cantidadTotal", Value: 1},
				{Key: "totalDescuento", Value: 1},
				{Key: "sudTotal", Value: 1},
				{Key: "proveedorNombre", Value: utils.ArrayElemAt("$proveedor.nombre", 0)},
				{Key: "proveedorApoellido", Value: utils.ArrayElemAt("$proveedor.apellidos", 0)},
				{Key: "usuario", Value: utils.ArrayElemAt("$usuario.username", 0)},
			}},
		},
		bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "fecha", Value: -1},
			}},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return &[]bson.M{}, err
	}
	defer cursor.Close(ctx)
	var resultado []bson.M
	cursor.All(ctx, &resultado)

	return &resultado, nil
}
