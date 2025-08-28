package service

import (
	"comercial-backend/src/core/enum"
	coreUtil "comercial-backend/src/core/utils"
	ingreso "comercial-backend/src/modules/ingreso/service"
	productoRepository "comercial-backend/src/modules/producto/repository"
	"comercial-backend/src/modules/stock/dto"
	"comercial-backend/src/modules/stock/model"
	repositoryStock "comercial-backend/src/modules/stock/repository"
	stockUtil "comercial-backend/src/modules/stock/utils"
	"context"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegitrarStockService(body dto.IngresoStockData, ctx context.Context) (*bson.ObjectID, error) {
	fecha := coreUtil.FechaHoraBolivia()

	ingresoStock, err := stockUtil.ConvertirDtoAIngreso(body)
	if err != nil {
		return &bson.NilObjectID,err
	}
	idIngreso, err := ingreso.RegistrarIngresoStockService(&ingresoStock, ctx)
	if err != nil {
		return  &bson.NilObjectID,err
	}
	documentos, err := repositoryStock.CountDocumentsStockRepository(ctx)
	if err != nil {
		return &bson.NilObjectID,err
	}
	contador := int(documentos)

	for _, v := range body.Stock {
		contador++
		productoId, err := coreUtil.ValidadIdMongo(v.Producto)

		if err != nil {
			return &bson.NilObjectID,err
		}

		stock, err := repositoryStock.VerificarStockRepository(*productoId, &v.FechaVencimiento, ctx)

		if err != nil {

			if err == mongo.ErrNoDocuments {

				producto, err := productoRepository.VerificarProductoRepository(*productoId, ctx)
				if err != nil {
					return &bson.NilObjectID,errors.New("Ocurrio un error al verificar el prodcuto " + err.Error())
				}

				consonante := coreUtil.GenerarCodigo(producto.Nombre)
				var codigo string = consonante + "-" + strconv.Itoa(int(contador))
				var nuevoStock = model.StockModel{
					Cantidad:         v.Cantidad,
					FechaVencimiento: v.FechaVencimiento,
					Codigo:           codigo,
					Producto:         *productoId,
					PrecioUnitario:   v.PrecioUnitario,
					Flag:             enum.EstadoNuevo,
					Fecha:            fecha,
				}
				err = repositoryStock.RegistrarStockRepository(&nuevoStock, ctx)
				if err != nil {
					return &bson.NilObjectID,errors.New("ocurrio un error al registrar el stock " + err.Error())
				}
			} else {
				return &bson.NilObjectID,err
			}

		} else {
			var cantidad int = v.Cantidad + stock.Cantidad
			err = repositoryStock.ActualizarStockRepository(stock.ID, cantidad, ctx)
			if err != nil {
				return &bson.NilObjectID,err
			}
		}

	}

	return idIngreso , nil
}

func ListarStockService(ctx context.Context) (*[]bson.M, error) {
	resultado, err := repositoryStock.ListarStockRepository(ctx)
	if err != nil {
		return &[]bson.M{}, err
	}
	return resultado, nil
}
