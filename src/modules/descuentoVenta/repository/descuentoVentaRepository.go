package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/descuentoVenta/model"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CrearDescuentoVentaRepository(data *model.DescuentoVenta, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.DescuentoVenta)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func ObtenerDescuentoVentaRepository(sucursalID *bson.ObjectID, ctx context.Context) (*model.DescuentoVenta, error) {
	collection := config.MongoDatabase.Collection(enum.DescuentoVenta)
	var resultado model.DescuentoVenta
	err := collection.FindOne(ctx, bson.M{"flag": enum.EstadoNuevo, "sucursal": sucursalID}, options.FindOne().SetSort(bson.D{{Key: "fecha", Value: -1}})).Decode(&resultado)
	if err != nil {
		return nil, err
	}
	return &resultado, nil
}

func ListarDescuentoVentaRepository(ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.DescuentoVenta)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},
		utils.Lookup("Sucursal", "sucursal", "_id", "sucursal"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "sucursal", Value: utils.ArrayElemAt("$sucursal.nombre", 0)},
				{Key: "alquiler", Value: 1},
				{Key: "vendedor", Value: 1},
			}},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var resultado []bson.M
	err = cursor.All(ctx, &resultado)
	if err != nil {
		return nil, err
	}
	return &resultado, nil
}

func EliminarDescuentoVentaRepository(id *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.DescuentoVenta)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"flag": enum.EstadoEliminado}})
	if err != nil {
		return err
	}
	return nil
}
