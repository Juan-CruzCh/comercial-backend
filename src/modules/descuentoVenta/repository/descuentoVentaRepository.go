package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/descuentoVenta/model"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
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
