package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/usuario/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func ListarUsuarioRepository(ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Usuario)
	var usuario []bson.M
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},
		utils.Lookup("Sucursal", "sucursal", "_id", "sucursal"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "ci", Value: 1},
				{Key: "nombre", Value: 1},
				{Key: "apellidos", Value: 1},
				{Key: "rol", Value: 1},
				{Key: "username", Value: 1},
				{Key: "sucursal", Value: utils.ArrayElemAt("$sucursal.nombre", 0)},
			}},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return &[]bson.M{}, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &usuario)
	return &usuario, nil
}

func EliminarUsuarioRepository(id *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Usuario)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"flag": enum.EstadoEliminado}})
	if err != nil {
		return err
	}
	return nil
}

func BuscarUsuarioIdRepository(id *bson.ObjectID, ctx context.Context) (*bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Usuario)
	var resultado []bson.M

	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "_id", Value: id},
				{Key: "flag", Value: enum.EstadoNuevo},
			}},
		},
		utils.Lookup("Sucursal", "sucursal", "_id", "sucursal"),
		bson.D{
			{Key: "$project", Value: bson.D{

				{Key: "nombre", Value: 1},
				{Key: "apellidos", Value: 1},
				{Key: "rol", Value: 1},
				{Key: "username", Value: 1},
				{Key: "sucursal", Value: utils.ArrayElemAt("$sucursal.nombre", 0)},
			}},
		},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	cursor.All(ctx, &resultado)
	return &resultado[0], nil
}
