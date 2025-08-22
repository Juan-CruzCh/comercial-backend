package stock

import (
	"comercial-backend/src/core/enum"
	coreUtil "comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso"
	productoRepository "comercial-backend/src/modules/producto/repository"
	"comercial-backend/src/modules/stock/dto"
	"comercial-backend/src/modules/stock/model"
	repositoryStock "comercial-backend/src/modules/stock/repository"
	stockUtil "comercial-backend/src/modules/stock/utils"
	"strconv"

	//"comercial-backend/src/modules/stock/utils"
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegitrarStockService(body dto.IngresoStockData, ctx context.Context) error {
	fecha, err := coreUtil.FechaHoraBolivia()
	if err != nil {
		return err
	}
	ingresoStock, err := stockUtil.ConvertirDtoAIngreso(body)
	if err != nil {
		return err
	}
	err = ingreso.RegistrarIngresoStockService(&ingresoStock, ctx)
	if err != nil {
		return err
	}
	documentos, err := repositoryStock.CountDocumentsStockRepository(ctx)
	if err != nil {
		return err
	}
	contador := int(documentos)

	for _, v := range body.Stock {
		contador++
		productoId, err := coreUtil.ValidadIdMongo(v.Producto)

		if err != nil {
			return err
		}

		stock, err := repositoryStock.VerificarStockRepository(*productoId, v.FechaVencimiento, ctx)

		if err != nil {
			if err == mongo.ErrNoDocuments {

				producto, err := productoRepository.VerificarProductoRepository(*productoId, ctx)
				if err != nil {
					return err
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
					MontoTotal:       v.MontoTotal,
					Descuento:        v.Descuento,
					SubTotal:         v.SudTotal,
				}
				err = repositoryStock.RegistrarStockRepository(&nuevoStock, ctx)
				if err != nil {
					return err
				}
			} else {
				return err
			}

		} else {
			var cantidad int = v.Cantidad + stock.Cantidad
			err = repositoryStock.ActualizarStockRepository(stock.ID, cantidad, ctx)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
