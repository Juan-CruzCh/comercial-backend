package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/stock/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
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
	var stock model.StockModel
	err := collection.FindOne(ctx, bson.M{"producto": productoId, "fechaVencimiento": fechaVencimiento}).Decode(&stock)

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
