package producto

import (
	"comercial-backend/src/config"
	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ListarProductoService() {

}

func RegistrarProductoService(productoDto *dto.ProductoDto, categoria bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("Producto")
	model := model.ProductoModel{
		Codigo:    "falta",
		Nombre:    productoDto.Nombre,
		Categoria: categoria,
	}
	fmt.Printf("categoria (tipo): %T\n", model.Categoria)
	_, err := collection.InsertOne(ctx, model)

	if err != nil {

		return err
	}
	return nil
}
