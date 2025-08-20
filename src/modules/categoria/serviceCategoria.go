package categoria

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/modules/categoria/dto"
	"comercial-backend/src/modules/categoria/model"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func crearCategoriaService(categoria *dto.CategoriaDto, ctx context.Context) (*dto.CategoriaDto, error) {
	collection := config.MongoDatabase.Collection("Categoria")

	count, err := collection.CountDocuments(ctx, bson.M{"nombre": categoria.Nombre})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("la categoría '%s' ya existe", categoria.Nombre)
	}

	data := model.Categoria{
		Nombre: categoria.Nombre,
		Fecha:  time.Now(),
		Flag:   enum.EstadoNuevo,
	}
	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return categoria, nil
}
func ListarCategoriaService(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection("Categoria")
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	var data []bson.M
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func ObtenerCategoriaService(c *gin.Context) {

}

func ActualizarCategoriaService(c *gin.Context) {

}

func EliminarCategoriaService(c *gin.Context) {

}
