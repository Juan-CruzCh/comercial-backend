package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/stock/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CountDocumentsStockRepository(ctx context.Context) (int64, error) {
	collection := config.MongoDatabase.Collection(enum.Stock)
	documentos, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return documentos, nil
}

func VerificarStockRepository(productoId bson.ObjectID, fechaVencimiento *time.Time, ctx context.Context) (*model.StockModel, error) {
	collection := config.MongoDatabase.Collection(enum.Stock)
	var filter bson.M = bson.M{
		"producto": productoId,
	}
	if fechaVencimiento != nil && !fechaVencimiento.IsZero() {
		filter["fechaVencimiento"] = fechaVencimiento
	}
	var stock model.StockModel
	err := collection.FindOne(ctx, filter).Decode(&stock)

	if err != nil {
		return &model.StockModel{}, err
	}
	return &stock, nil
}

func RegistrarStockRepository(data *model.StockModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Stock)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return errors.New("Ocurrio un error al ingresar el estock" + err.Error())
	}
	return nil

}

func ActualizarStockRepository(stock bson.ObjectID, cantidad int, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Stock)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": stock}, bson.M{"$set": bson.M{"cantidad": cantidad}})
	if err != nil {
		return err
	}
	return nil
}

func BuscarStockRepository(stock *bson.ObjectID, ctx context.Context) (*model.StockModel, error) {
	collection := config.MongoDatabase.Collection(enum.Stock)
	var stockModel model.StockModel
	err := collection.FindOne(ctx, bson.M{"_id": stock, "flag": enum.EstadoNuevo}).Decode(&stockModel)
	if err != nil {
		return &model.StockModel{}, errors.New("no se enontro el estock: %v" + err.Error())
	}
	return &stockModel, nil
}

func ListarStockRepository(ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Stock)
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},
		utils.Lookup("Producto", "producto", "_id", "producto"),
		utils.Unwind("$producto", false),
		utils.Lookup("Categoria", "producto.categoria", "_id", "categoria"),
		utils.Lookup("UnidadManejo", "producto.unidadManejo", "_id", "unidadManejo"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "codigo", Value: 1},
				{Key: "cantidad", Value: 1},
				{Key: "precioUnitario", Value: 1},
				{Key: "fechaVencimiento", Value: 1},
				{Key: "descripcion", Value: "$producto.descripcion"},
				{Key: "producto", Value: "$producto.nombre"},
				{Key: "categoria", Value: utils.ArrayElemAt("$categoria.nombre", 0)},
				{Key: "unidadManejo", Value: utils.ArrayElemAt("$unidadManejo.nombre", 0)},
			},
			},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return &[]bson.M{}, err
	}
	defer cursor.Close(ctx)
	var stock []bson.M
	err = cursor.All(ctx, &stock)
	if err != nil {
		return &[]bson.M{}, err
	}

	return &stock, nil
}
