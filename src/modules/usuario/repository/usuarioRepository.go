package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/usuario/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func VeficarUsuarioExisteRepository(username *string, ctx context.Context) (*model.UsuarioModel, error) {
	collection := config.MongoDatabase.Collection(enum.Usuario)
	cantidad, err := collection.CountDocuments(ctx, bson.M{"flag": enum.EstadoNuevo, "username": username})
	if err != nil {
		return &model.UsuarioModel{}, err
	}
	if cantidad > 0 {
		return &model.UsuarioModel{}, errors.New("El usuario ya existe")
	}
	return nil, nil
}

func CrearUsuarioRepository(data *model.UsuarioModel, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Usuario)
	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil

}

func VeficarUsuarioRepository(username *string, ctx context.Context) (*model.UsuarioModel, error) {
	collection := config.MongoDatabase.Collection(enum.Usuario)
	var usuario model.UsuarioModel
	err := collection.FindOne(ctx, bson.M{"flag": enum.EstadoNuevo, "username": username}).Decode(&usuario)
	if err != nil {
		return &model.UsuarioModel{}, err
	}
	return &usuario, nil
}
