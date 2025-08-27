package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/ingreso/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CrearDetalleIngresoManyRepository(data []model.DetalleIngresoModel, ctx context.Context) error {
	collectionDetalleIngreso := config.MongoDatabase.Collection(enum.DetalleIngreso)
	_, err := collectionDetalleIngreso.InsertMany(ctx, data)
	return err
}

func ListarDetalleIngresoRepository(id *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {
	fmt.Println("id", id)
	collection := config.MongoDatabase.Collection(enum.DetalleIngreso)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
				//{Key: "ingreso", Value: id},
			}},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println(err)
		return &[]bson.M{}, err
	}
	defer cursor.Close(ctx)
	var resultado []bson.M
	err = cursor.All(ctx, &resultado)
	if err != nil {
		fmt.Println(err)
		return &[]bson.M{}, err
	}
	fmt.Println(resultado)
	return &[]bson.M{}, nil
}
