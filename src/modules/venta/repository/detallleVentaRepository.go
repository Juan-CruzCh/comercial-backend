package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/venta/model"
	"context"
	"errors"
)

func RealizarVentaDetalleRepository(detalleVenta *model.DetalleVentaModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.DetalleVenta)
	_, err := collection.InsertOne(ctx, detalleVenta)
	if err != nil {
		return errors.New("Ocurrio un error en el detalle venta " + err.Error())
	}
	return nil
}
