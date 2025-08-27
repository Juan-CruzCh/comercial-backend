package service

import (
	"comercial-backend/src/modules/venta/repository"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func DetalleVentaService(idVenta *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {
	resultado, err := repository.DetalleVentaRepository(idVenta, ctx)
	if err != nil {
		return &[]bson.M{}, err
	}
	return resultado, nil
}
