package proveedor

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/proveedor/dto"
	"comercial-backend/src/modules/proveedor/model"
	"comercial-backend/src/modules/proveedor/repository"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func registrarProveedorService(proveedor *dto.ProveedorDto, ctx context.Context) error {

	var proveedorModel = model.ProveedorModel{
		Nombre:    proveedor.Nombre,
		CI:        proveedor.CI,
		Apellidos: proveedor.Apellidos,
		Empresa:   proveedor.Empresa,
		Flag:      enum.EstadoNuevo,
		Fecha:     time.Now(),
		Celular:   proveedor.Celular,
	}
	err := repository.CrearProveedorRepository(&proveedorModel, ctx)
	if err != nil {
		return err
	}
	return nil

}

func listarProveedorService(ctx context.Context) ([]bson.M, error) {
	data, err := repository.ListarProveedorRepository(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil

}
