package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"

	"comercial-backend/src/modules/descuentoVenta/dto"
	"comercial-backend/src/modules/descuentoVenta/model"
	"comercial-backend/src/modules/descuentoVenta/repository"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CrearDescuentoVentaService(body *dto.DescuentoVentaDto, ctx context.Context) error {
	fecha := utils.FechaHoraBolivia()
	sucursalID, err := utils.ValidadIdMongo(body.Sucursal)
	var data model.DescuentoVenta = model.DescuentoVenta{
		Alquiler: body.Alquiler,
		Vendedor: body.Vendedor,
		Fecha:    fecha,
		Sucursal: *sucursalID,
		Flag:     enum.EstadoNuevo,
	}
	err = repository.CrearDescuentoVentaRepository(&data, ctx)
	if err != nil {
		return err
	}
	return nil
}

func ListarDescuentoVentaService(ctx context.Context) (*[]bson.M, error) {
	data, err := repository.ListarDescuentoVentaRepository(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
