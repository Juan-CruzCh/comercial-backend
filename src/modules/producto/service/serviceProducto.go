package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/producto/dto"
	"comercial-backend/src/modules/producto/model"
	"comercial-backend/src/modules/producto/repository"
	"comercial-backend/src/modules/producto/structs"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ListarProductoService(filtros *structs.FiltrosProductoStruct, pagina int, limite int, ctx context.Context) ([]bson.M, error) {
	data, err := repository.ListarProductoRepository(filtros, pagina, limite, ctx)
	if err != nil {

		return []bson.M{}, err
	}
	if data == nil {

		return []bson.M{}, nil
	}

	return data, nil
}

func RegistrarProductoService(productoDto *dto.ProductoDto, categoria *bson.ObjectID, unidadManejo *bson.ObjectID, ctx context.Context) (bson.M, error) {
	cantidad, err := repository.CountDocumentsProductoRepository(ctx)
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
	producto, err := repository.CrearProductoRepository(&model, ctx)
	if err != nil {

		return bson.M{}, err
	}
	return producto, nil
}
