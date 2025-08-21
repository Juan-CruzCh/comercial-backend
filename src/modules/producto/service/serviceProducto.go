package service

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/model"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ListarProductoService(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection("Producto")
	cursor, err := collection.Find(ctx, bson.M{"flag": "nuevo"})
	if err != nil {
		return nil, err
	}
	var producto []bson.M
	err = cursor.All(ctx, &producto)
	if err != nil {
		return nil, err
	}
	return producto, nil

}

func RegistrarProductoService(productoDto *dto.ProductoDto, categoria *bson.ObjectID, unidadManejo *bson.ObjectID, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("Producto")
	cantidad, err := collection.CountDocuments(ctx, bson.M{"flag": "nuevo"})
	if err != nil {
		return err
	}
	var codigo string = utils.GenerarCodigo(productoDto.Nombre)
	var cantidadSrt string = strconv.Itoa(int(cantidad))
	codigo = codigo + "-" + cantidadSrt
	model := model.ProductoModel{
		Codigo:       codigo,
		Nombre:       productoDto.Nombre,
		Categoria:    *categoria,
		Fecha:        time.Now(),
		Flag:         enum.EstadoNuevo,
		UnidadManejo: *unidadManejo,
		Descripcion:  productoDto.Descripcion,
	}

	_, err = collection.InsertOne(ctx, model)

	if err != nil {

		return err
	}
	return nil
}
