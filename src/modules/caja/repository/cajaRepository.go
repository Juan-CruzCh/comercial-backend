package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/structCore"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/caja/model"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AbrirCajaRepository(data *model.CajaModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func VerificarCajaAbierto(usuario *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	filter := bson.M{
		"usuario": usuario,
		"estado":  enum.Abierto,
	}
	cantidad, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if cantidad > 0 {
		return errors.New("La caja ya se encuentra abierta")
	}
	return nil
}

func CerrarCajaRepository(idCaja *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	fecha := utils.FechaHoraBolivia()
	_, err := collection.UpdateOne(ctx, bson.M{"_id": idCaja, "flag": enum.EstadoNuevo}, bson.M{"$set": bson.M{"estado": enum.Cerrado, "fechaCierre": fecha}})
	if err != nil {
		return err
	}
	return nil
}

func BuscarCajaUsuarioRepository(usuario *bson.ObjectID, ctx context.Context) (*model.CajaModel, error) {
	collection := config.MongoDatabase.Collection(enum.Caja)
	var caja model.CajaModel
	err := collection.FindOne(ctx, bson.M{"usuario": usuario, "estado": enum.Abierto, "flag": enum.EstadoNuevo}).Decode(&caja)
	if err != nil {
		return &model.CajaModel{}, err
	}
	return &caja, nil
}

func VerificarCajaRepository(usuario *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	filter := bson.M{
		"usuario": usuario,
		"estado":  enum.Abierto,
		"flag":    enum.EstadoNuevo,
	}
	cantidad, err := collection.CountDocuments(ctx, filter)
	fmt.Println(cantidad)
	if err != nil {
		return err
	}

	if cantidad > 0 {
		return nil
	}
	return errors.New("debe abrir la caja")
}

func AsignarTotalVentasCajaRepository(usuario *bson.ObjectID, totalVenta float64, montoFinal float64, totalDescuento float64, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	_, err := collection.UpdateOne(ctx, bson.M{"usuario": usuario, "estado": enum.Abierto, "flag": enum.EstadoNuevo}, bson.M{"$set": bson.M{
		"totalVentas":    totalVenta,
		"montoFinal":     montoFinal,
		"totalDescuento": totalDescuento,
	}})
	if err != nil {
		return err
	}
	return nil
}

func ListarCajaRespository(pagina int, limite int, ctx context.Context) (*structCore.ResultadoPaginado, error) {
	collection := config.MongoDatabase.Collection(enum.Caja)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
				{Key: "estado", Value: enum.Cerrado},
			}},
		},
		utils.Lookup("Usuario", "usuario", "_id", "usuario"),
		utils.Lookup("Sucursal", "sucursal", "_id", "sucursal"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "fechaApertura", Value: 1},
				{Key: "montoInicial", Value: 1},
				{Key: "montoFinal", Value: 1},
				{Key: "totalVentas", Value: 1},
				{Key: "estado", Value: 1},
				{Key: "fechaCierre", Value: 1},
				{Key: "usuario", Value: utils.ArrayElemAt("$usuario.username", 0)},
				{Key: "sucursal", Value: utils.ArrayElemAt("$sucursal.nombre", 0)},
			}},
		},
		bson.D{
			{Key: "$skip", Value: utils.Skip(pagina, limite)},
		},
		bson.D{
			{Key: "$limit", Value: limite},
		},
		bson.D{
			{Key: "$sort", Value: bson.D{
				{Key: "fechaCierre", Value: -1},
			}},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}
	cantidad, err := collection.CountDocuments(ctx, bson.D{
		{Key: "flag", Value: enum.EstadoNuevo},
		{Key: "estado", Value: enum.Cerrado},
	})
	if err != nil {
		return nil, err
	}
	var paginas int = utils.CalcularPaginas(int(cantidad), limite)
	var data []bson.M
	defer cursor.Close(ctx)
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
