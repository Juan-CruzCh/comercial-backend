package proveedor

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/proveedor/dto"
	"comercial-backend/src/modules/proveedor/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func registrarProveedorService(proveedor *dto.ProveedorDto, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("Proveedor")
	var proveedorModel =model.ProveedorModel {
		Nombre: proveedor.Nombre,
		CI: proveedor.CI,
		Apellidos: proveedor.Apellidos,
		Empresa: proveedor.Empresa,
		Flag: enum.EstadoNuevo,
		Fecha: time.Now(),
		Celular: proveedor.Celular,
	}
	_,err :=collection.InsertOne(ctx,proveedorModel)
	if err != nil {
		return err
	}
	return  nil

}

func listarProveedorService(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection("Proveedor")
	
	cursor,err :=collection.Find(ctx, bson.M{"flag":enum.EstadoNuevo})
	if err != nil {
		return nil, err
	}
	var proveedores []bson.M
	err = cursor.All(ctx, &proveedores)
	if err != nil {
		return nil, err
	}
	
	return  proveedores, nil

}