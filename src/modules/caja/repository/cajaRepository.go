package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/caja/model"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func AbrirCajaRepository(data *model.CajaModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func VerificarCajaAbierto(usuario *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	filter := bson.M{
		"usuario": usuario,
		"estado":  enum.Abierto,
	}
	cantidad, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if cantidad > 0 {
		return errors.New("La caja ya se encuentra abierta")
	}
	return nil
}

func CerrarCajaRepository(idCaja *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	fecha := utils.FechaHoraBolivia()
	_, err := collection.UpdateOne(ctx, bson.M{"_id": idCaja, "flag": enum.EstadoNuevo}, bson.M{"$set": bson.M{"estado": enum.Cerrado, "fechaCierre": fecha}})
	if err != nil {
		return err
	}
	return nil
}

func BuscarCajaUsuarioRepository(usuario *bson.ObjectID, ctx context.Context) (*model.CajaModel, error) {
	collection := config.MongoDatabase.Collection(enum.Caja)
	var caja model.CajaModel
	err := collection.FindOne(ctx, bson.M{"usuario": usuario, "estado": enum.Abierto, "flag": enum.EstadoNuevo}).Decode(&caja)
	if err != nil {
		return &model.CajaModel{}, err
	}
	return &caja, nil
}

func VerificarCajaRepository(usuario *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Caja)
	filter := bson.M{
		"usuario": usuario,
		"estado":  enum.Abierto,
		"flag":    enum.EstadoNuevo,
	}
	cantidad, err := collection.CountDocuments(ctx, filter)
	fmt.Println(cantidad)
	if err != nil {
		return err
	}

	if cantidad > 0 {
		return nil
	}
	return errors.New("Debe abrir la caja")
}
