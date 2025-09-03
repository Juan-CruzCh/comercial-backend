package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso/model"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CrearDetalleIngresoManyRepository(data []model.DetalleIngresoModel, ctx context.Context) error {
	collectionDetalleIngreso := config.MongoDatabase.Collection(enum.DetalleIngreso)
	_, err := collectionDetalleIngreso.InsertMany(ctx, data)
	return err
}

func ListarDetalleIngresoRepository(id *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.DetalleIngreso)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
				{Key: "ingreso", Value: id},
			}},
		},
		utils.Lookup("Producto", "producto", "_id", "producto"),

		utils.Lookup("Categoria", "producto.0.categoria", "_id", "categoria"),

		utils.Lookup("UnidadManejo", "producto.0.unidadManejo", "_id", "unidaManejo"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "cantidad", Value: 1},
				{Key: "precioUnitario", Value: 1},
				{Key: "montoTotal", Value: 1},
				{Key: "descuento", Value: 1},
				{Key: "sudTotal", Value: 1},
				{Key: "fechaVencimiento", Value: 1},
				{Key: "producto", Value: utils.ArrayElemAt("$producto.nombre", 0)},
				{Key: "descripcion", Value: utils.ArrayElemAt("$producto.descripcion", 0)},
				{Key: "codigo", Value: utils.ArrayElemAt("$producto.codigo", 0)},
				{Key: "categoria", Value: utils.ArrayElemAt("$categoria.nombre", 0)},
				{Key: "unidadManejo", Value: utils.ArrayElemAt("$unidaManejo.nombre", 0)},
			}},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return &[]bson.M{}, err
	}
	defer cursor.Close(ctx)
	var resultado []bson.M
	err = cursor.All(ctx, &resultado)
	if err != nil {
		return &[]bson.M{}, err
	}

	return &resultado, nil
}
