package proveedor

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/modules/proveedor/dto"
	"comercial-backend/src/modules/proveedor/model"
	"context"
	"time"
)

func registrarProveedorService(proveedor *dto.ProveedorDto, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("Proveedor")
	var proveedorModel =model.ProveedorModel {
		Nombre: proveedor.Nombre,
		CI: proveedor.CI,
		Apellidos: proveedor.Apellidos,
		Empresa: proveedor.Empresa,
		Flag: "nuevo",
		Fecha: time.Now(),
	}
	_,err :=collection.InsertOne(ctx,proveedorModel)
	if err != nil {
		return err
	}
	return  nil

}