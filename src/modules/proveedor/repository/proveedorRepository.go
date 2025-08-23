package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/proveedor/model"

	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CrearProveedorRepository(data *model.ProveedorModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Proveedor)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func ListarProveedorRepository(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Proveedor)

	cursor, err := collection.Find(ctx, bson.M{"flag": enum.EstadoNuevo})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var proveedores []bson.M
	err = cursor.All(ctx, &proveedores)
	if err != nil {
		return nil, err
	}

	return proveedores, nil
}
