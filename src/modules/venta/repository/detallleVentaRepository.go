package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/venta/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RealizarVentaDetalleRepository(detalleVenta *model.DetalleVentaModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.DetalleVenta)
	_, err := collection.InsertOne(ctx, detalleVenta)
	if err != nil {
		return errors.New("Ocurrio un error en el detalle venta " + err.Error())
	}
	return nil
}

func DetalleVentaRepository(idVenta *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {

	collection := config.MongoDatabase.Collection(enum.DetalleVenta)
	cursor, err := collection.Find(ctx, bson.M{"venta": idVenta})
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
