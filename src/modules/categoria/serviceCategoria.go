package categoria

import (
	"comercial-backend/src/config"
	"comercial-backend/src/modules/categoria/dto"
	"comercial-backend/src/modules/categoria/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func crearCategoriaService(categoria *dto.CategoriaDto) (*dto.CategoriaDto , error) {
	collection := config.MongoDatabase.Collection("Categoria")	

	count, err := collection.CountDocuments(context.TODO(), bson.M{"nombre": categoria.Nombre})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("la categoría '%s' ya existe", categoria.Nombre)
	}

	data  := model.Categoria {
		Nombre:categoria.Nombre,
		Fecha: time.Now(),
		Flag: "nuevo",
	}
	_, err =collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return categoria, nil
}
func ListarCategoriaService() ([]bson.M, error){
	collection := config.MongoDatabase.Collection("Categoria")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	
		if err != nil {
		return nil, err
	}
	var data []bson.M 
	err =cursor.All(context.TODO(), & data)
		if err != nil {
		return nil, err
	}
	return  data , nil

}