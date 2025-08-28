package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/usuario/model"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func VeficarUsuarioRepository( username *string, ctx context.Context) (*model.UsuarioModel, error) {
	collection:=config.MongoDatabase.Collection(enum.Usuario)
	var usuario model.UsuarioModel 
	err := collection.FindOne(ctx, bson.M{"flag":enum.EstadoNuevo, "username":username}).Decode(&usuario)
	if err != nil {
		return &model.UsuarioModel{}, err
	}
	return &usuario, nil

}