package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/structCore"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/venta/dto"
	"comercial-backend/src/modules/venta/model"
	"context"
	"errors"
	"fmt"
	"time"

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

func ListarVentasRepository(filtros *dto.BuscadorVentaDto, pagina int, limite int, halibitarPaginador bool, ctx context.Context) (*structCore.ResultadoPaginado, error) {
	collection := config.MongoDatabase.Collection(enum.Venta)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},
		utils.Lookup("Sucursal", "sucursal", "_id", "sucursal"),
		utils.Lookup("Usuario", "usuario", "_id", "usuario"),
	}
	if filtros.Codigo != "" {
		pipeline = append(pipeline, utils.RegexMatch("codigo", filtros.Codigo))
	}
	if filtros.FechaFin != "" && filtros.FechaInicio != "" {
		f1, f2, err := utils.NormalizarRangoDeFechas(filtros.FechaInicio, filtros.FechaFin)
		if err != nil {
			return nil, err
		}
		pipeline = append(pipeline, bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "fechaVenta", Value: bson.D{
					{Key: "$gte", Value: f1},
					{Key: "$lte", Value: f2},
				}},
			}},
		})
	}
	if filtros.Sucursal != "" {
		ID, err := utils.ValidadIdMongo(filtros.Sucursal)
		if err != nil {
			return nil, err
		}
		pipeline = append(pipeline, bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "sucursal._id", Value: ID},
			}},
		})
	}
	if filtros.Usuario != "" {
		ID, err := utils.ValidadIdMongo(filtros.Usuario)
		if err != nil {
			return nil, err
		}
		pipeline = append(pipeline, bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "usuario._id", Value: ID},
			}},
		})
	}

	pipeline = append(pipeline, bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "codigo", Value: 1},
			{Key: "montoTotal", Value: 1},
			{Key: "subTotal", Value: 1},
			{Key: "fechaVenta", Value: 1},
			{Key: "descuento", Value: 1},
			{Key: "descuentoAlquiller", Value: 1},
			{Key: "descuentoVendedor", Value: 1},
			{Key: "descuentoAcumulado", Value: 1},
			{Key: "totalGanancia", Value: 1},
			{Key: "sucursal", Value: utils.ArrayElemAt("$sucursal.nombre", 0)},
			{Key: "vendedor", Value: utils.ArrayElemAt("$usuario.username", 0)},
		}},
	})
	pipeline = append(pipeline, bson.D{
		{Key: "$sort", Value: bson.D{
			{Key: "fechaVenta", Value: -1},
		}},
	})
	if halibitarPaginador {
		pipeline = append(pipeline, bson.D{
			{Key: "$skip", Value: utils.Skip(pagina, limite)},
		})
		pipeline = append(pipeline,
			bson.D{
				{Key: "$limit", Value: limite},
			},
		)

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
func ListarVentaMesualRepository(sucursal *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Venta)
	now := time.Now()
	inicioMes := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	hoy := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 99, now.Location())

	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "sucursal", Value: sucursal},
				{Key: "fechaVenta", Value: bson.D{
					{Key: "$gte", Value: inicioMes},
					{Key: "$lte", Value: hoy},
				}},
			}},
		},
		utils.Sort("fechaVenta"),
		utils.Lookup("DetalleVenta", "_id", "venta", "detalleVenta"),
		utils.Unwind("$detalleVenta", false),
		bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "aqo", Value: bson.D{{Key: "$year", Value: "$fechaVenta"}}},
					{Key: "mes", Value: bson.D{{Key: "$month", Value: "$fechaVenta"}}},
					{Key: "dia", Value: bson.D{{Key: "$dayOfMonth", Value: "$fechaVenta"}}},
				}},
				{Key: "fecha", Value: bson.D{{Key: "$first", Value: "$fechaVenta"}}},
				{Key: "montoTotal", Value: bson.D{{Key: "$sum", Value: "$montoTotal"}}},
				{Key: "cantidad", Value: bson.D{{Key: "$sum", Value: "$detalleVenta.cantidad"}}},
			}},
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "montoTotal", Value: 1},
				{Key: "cantidad", Value: 1},
				{Key: "fecha", Value: 1},
				{Key: "_id", Value: 0},
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
		fmt.Println(err)
		return nil, err
	}

	return &resultado, nil
}
