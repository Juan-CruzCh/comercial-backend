package repository

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/categoria/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CrearCategoriaRepository(data *model.Categoria, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Categoria)

	count, err := collection.CountDocuments(ctx, bson.M{"nombre": data.Nombre})
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("la categoría '%s' ya existe", data.Nombre)
	}

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func ListarCategoriaRepository(ctx context.Context) (*[]bson.M, error) {
	collection := config.MongoDatabase.Collection(enum.Categoria)
	cursor, err := collection.Find(ctx, bson.M{"flag": enum.EstadoNuevo})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var resultado []bson.M

	cursor.All(ctx, &resultado)

	return &resultado, nil
}

func EliminarCategoriaRepository(categoriaId *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection(enum.Categoria)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": categoriaId}, bson.M{"$set": bson.M{"flag": enum.EstadoEliminado}})
	if err != nil {
		return err
	}

	return nil
}
