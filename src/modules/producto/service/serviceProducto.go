package service

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/model"
	"context"
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func ListarProductoService(ctx context.Context) ([]bson.M, error) {
	collection := config.MongoDatabase.Collection("Producto")
	var pipeline = mongo.Pipeline{
		bson.D{
            {Key: "$match", Value: bson.D{
                {Key: "flag", Value: enum.EstadoNuevo},
            }},  
    },
	utils.Lookup("Categoria", "categoria","_id","categoria"),
	utils.Lookup("UnidadManejo", "unidadManejo","_id","unidadManejo"),	
	bson.D{
		{Key: "$project", Value: bson.D {
				{Key: "nombre", Value: 1},
				{Key: "descripcion", Value: 1},
				{Key: "codigo", Value: 1},
				{Key: "categoria", Value: bson.D{
					 {Key: "$arrayElemAt", Value: bson.A{"$categoria.nombre", 0}},
				}},
				{Key: "unidadManejo", Value: bson.D{
					 {Key: "$arrayElemAt", Value: bson.A{"$unidadManejo.nombre", 0}},
				}},
		} ,
		},
	}	,

	bson.D{
		{Key: "$sort", Value: bson.D{
			{Key: "fecha",Value: -1},
		},},
	},
}
	cursor, err := collection.Aggregate(ctx, pipeline)
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

func RegistrarProductoService(productoDto *dto.ProductoDto, categoria *bson.ObjectID, unidadManejo *bson.ObjectID, ctx context.Context) (bson.M, error) {
	collection := config.MongoDatabase.Collection("Producto")
	cantidad, err := collection.CountDocuments(ctx, bson.M{"flag": enum.EstadoNuevo})
	if err != nil {
		return bson.M{}, err
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

	resultado, err := collection.InsertOne(ctx, model)

	if err != nil {

		return bson.M{}, err
	}
	id, ok := resultado.InsertedID.(bson.ObjectID)
	if !ok {
		return bson.M{}, errors.New("ocurrio un error al insertar el ingreso")
	}
	var producto bson.M
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&producto)

	if err != nil {

		return bson.M{}, err
	}
	return producto, nil
}
