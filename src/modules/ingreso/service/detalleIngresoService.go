package service

import (
	"comercial-backend/src/modules/ingreso/repository"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ListarDetalleIngresoService(id *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {
	resultado, err := repository.ListarDetalleIngresoRepository(id, ctx)
	fmt.Println("service")
	if err != nil {
		return &[]bson.M{}, err
	}
	return resultado, nil
}
