package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/producto/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ListarProductoRepository(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection("Producto")
	var pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},
		utils.Lookup("Categoria", "categoria", "_id", "categoria"),
		utils.Lookup("UnidadManejo", "unidadManejo", "_id", "unidadManejo"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "nombre", Value: 1},
				{Key: "descripcion", Value: 1},
				{Key: "codigo", Value: 1},
				{Key: "categoria", Value: utils.ArrayElemAt("$categoria.nombre", 0)},
				{Key: "unidadManejo", Value: utils.ArrayElemAt("$unidadManejo.nombre", 0)},
			},
			},
		},

		bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "fecha", Value: -1},
			}},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var producto []bson.M
	err = cursor.All(ctx, &producto)
	if err != nil {
		return nil, err
	}
	return producto, nil
}

func CrearProductoRepository(data *model.ProductoModel, ctx context.Context) (bson.M, error) {
	collection := config.MongoDatabase.Collection("Producto")
	resultado, err := collection.InsertOne(ctx, data)

	if err != nil {

		return bson.M{}, err
	}
	id, ok := resultado.InsertedID.(bson.ObjectID)
	if !ok {
		return bson.M{}, errors.New("ocurrio un error al insertar el ingreso")
	}
	var producto bson.M
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&producto)

	if err != nil {

		return bson.M{}, err
	}
	return producto, nil

}

func CountDocumentsProductoRepository(ctx context.Context) (int64, error) {
	collection := config.MongoDatabase.Collection("Producto")
	countDocuments, err := collection.CountDocuments(ctx, bson.M{"flag": enum.EstadoNuevo})
	if err != nil {
		return 0, err
	}
	return countDocuments, nil
}

func VerificarProductoRepository(producuto bson.ObjectID, ctx context.Context) (*model.ProductoModel, error) {
	collection := config.MongoDatabase.Collection("Producto")
	var producto model.ProductoModel
	err := collection.FindOne(ctx, bson.M{"_id": producuto, "flag": enum.EstadoNuevo}).Decode(&producto)
	if err != nil {
		return &model.ProductoModel{}, err
	}
	return &producto, nil
}
