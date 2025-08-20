package almacen

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/almacen/dto"
	"comercial-backend/src/modules/almacen/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func registrarAlmacenService(almacenDto *dto.AlmacenDto, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("Almacen")
	countDocuments, err := collection.CountDocuments(ctx, bson.M{"flag": enum.EstadoNuevo, "nombre": almacenDto.Nombre})
	if err != nil {
		return err
	}
	if countDocuments > 0 {
		return errors.New("el almacen ya se encuetra registrado")
	}
	var model = model.AlmacenModel{
		Nombre: almacenDto.Nombre,
		Fecha:  time.Now(),
		Flag:   enum.EstadoNuevo,
	}
	_, err = collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}
	return nil
}

func listarAlmacenService(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection("Almacen")
	cursor, err := collection.Find(ctx, bson.M{"flag": enum.EstadoNuevo})
	if err != nil {
		return nil, err
	}
	var almacen []bson.M
	cursor.All(ctx, &almacen)
	return almacen, nil
}
