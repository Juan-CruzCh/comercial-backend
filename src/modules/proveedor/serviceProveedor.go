package proveedor

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/structCore"
	"comercial-backend/src/modules/proveedor/dto"
	"comercial-backend/src/modules/proveedor/model"
	"comercial-backend/src/modules/proveedor/repository"
	"context"
	"time"
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

func listarProveedorService(ci string, nombre string, celular string, empresa string, pagina int, limite int, ctx context.Context) (*structCore.ResultadoPaginado, error) {
	data, err := repository.ListarProveedorRepository(ci, nombre, celular, empresa, pagina, limite, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil

}
