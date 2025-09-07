package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/structCore"
	"comercial-backend/src/core/utils"
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
		return &bson.NilObjectID, errors.New("se prodcuto un erro al ingresa la venta")
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

func ListarVentasRepository(pagina int, limite int, ctx context.Context) (*structCore.ResultadoPaginado, error) {
	collection := config.MongoDatabase.Collection(enum.Venta)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},
		utils.Lookup("Sucursal", "sucursal", "_id", "sucursal"),
		utils.Lookup("Usuario", "usuario", "_id", "usuario"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "codigo", Value: 1},
				{Key: "montoTotal", Value: 1},
				{Key: "subTotal", Value: 1},
				{Key: "fechaVenta", Value: 1},
				{Key: "descuento", Value: 1},
				{Key: "sucursal", Value: utils.ArrayElemAt("$sucursal.nombre", 0)},
				{Key: "vendedor", Value: utils.ArrayElemAt("$usuario.username", 0)},
			}},
		},
		bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "fechaVenta", Value: -1},
			}},
		},

		bson.D{
			{Key: "$skip", Value: utils.Skip(pagina, limite)},
		},
		bson.D{
			{Key: "$limit", Value: limite},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {

		return nil, err
	}
	defer cursor.Close(ctx)
	cantidad, err := collection.CountDocuments(ctx, bson.M{"flag": enum.EstadoNuevo})
	if err != nil {

		return nil, err
	}
	paginas := utils.CalcularPaginas(int(cantidad), limite)
	var data []bson.M
	err = cursor.All(ctx, &data)
	if err != nil {

		return nil, err
	}
	var resultado structCore.ResultadoPaginado = structCore.ResultadoPaginado{
		Data:    data,
		Paginas: paginas,
	}
	return &resultado, nil

}
func BuscarVentaPorIdRespository(idVenta *bson.ObjectID, ctx context.Context) (*bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Venta)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "_id", Value: idVenta},
			}},
		},
		utils.Lookup("DetalleVenta", "_id", "venta", "detalleVenta"),
		utils.Lookup("Usuario", "usuario", "_id", "usuario"),
		utils.Lookup("Sucursal", "sucursal", "_id", "sucursal"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "codigo", Value: 1},
				{Key: "usuario", Value: utils.ArrayElemAt("$usuario.username", 0)},
				{Key: "sucursal", Value: utils.ArrayElemAt("$sucursal.nombre", 0)},
				{Key: "detalleVenta", Value: 1},
				{Key: "fechaVenta", Value: 1},
				{Key: "montoTotal", Value: 1},
				{Key: "descuento", Value: 1},
				{Key: "subTotal", Value: 1},
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
	return &resultado[0], nil

}
