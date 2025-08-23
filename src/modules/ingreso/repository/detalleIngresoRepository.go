package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/ingreso/model"
	"context"
)

func CrearDetalleIngresoManyRepository(data []model.DetalleIngresoModel, ctx context.Context) error {
	collectionDetalleIngreso := config.MongoDatabase.Collection(enum.DetalleIngreso)
	_, err := collectionDetalleIngreso.InsertMany(ctx, data)
	return err
}
