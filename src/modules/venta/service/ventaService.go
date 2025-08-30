package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	stockRopository "comercial-backend/src/modules/stock/repository"
	"errors"
	"strconv"

	"comercial-backend/src/modules/venta/dto"
	"comercial-backend/src/modules/venta/model"
	"comercial-backend/src/modules/venta/repository"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RealizarVentaService(body *dto.VentaDto, ctx context.Context) (*bson.ObjectID, error) {

	fecha := utils.FechaHoraBolivia()
	err := validaStockProduct(&body.DetalleVenta, ctx)
	if err != nil {
		return &bson.NilObjectID, err
	}
	usuarioID, err := utils.ValidadIdMongo("68a8b4dbdb1be7def32f34a0")
	if err != nil {
		return &bson.NilObjectID, err
	}
	sucursalID, err := utils.ValidadIdMongo("68a8b4dbdb1be7def32f34a0")
	if err != nil {
		return &bson.NilObjectID, err
	}

	cantidad, _ := repository.CountDocumentsVentaRepository(ctx)
	var codigo string = "VEN-" + strconv.Itoa(int(cantidad))

	var venta model.VentaModel = model.VentaModel{
		Codigo:     codigo,
		MontoTotal: body.MontoTotal,
		FechaVenta: fecha,
		Fecha:      fecha,
		Usuario:    *usuarioID,
		Sucursal:   *sucursalID,
		TipoPago:   enum.Efectivo,
		Estado:     enum.Realizada,
		Flag:       enum.EstadoNuevo,
		Descuento:  *body.Descuento,
		SubTotal:   body.SudTotal,
	}
	ventaID, err := repository.RealizarVentaRepository(&venta, ctx)
	if err != nil {
		return &bson.NilObjectID, err

	}

	for _, v := range body.DetalleVenta {
		stockID, err := utils.ValidadIdMongo(v.Stock)
		if err != nil {
			return &bson.NilObjectID, err
		}
		stock, err := stockRopository.BuscarStockRepository(stockID, ctx)
		if err != nil {
			return &bson.NilObjectID, err
		}
		var nuevaCantidad int = stock.Cantidad - v.Cantidad
		err = stockRopository.ActualizarStockRepository(stock.ID, nuevaCantidad, ctx)
		if err != nil {
			return &bson.NilObjectID, err
		}
		var detalleVenta model.DetalleVentaModel = model.DetalleVentaModel{
			Producto:       stock.Producto,
			Stock:          *stockID,
			Cantidad:       v.Cantidad,
			Descripcion:    v.DescripcionProducto,
			Venta:          *ventaID,
			Fecha:          fecha,
			Flag:           enum.EstadoNuevo,
			PrecioUnitario: v.PrecioUnitario,
			PrecioTotal:    v.PrecioTotal,
		}
		err = repository.RealizarVentaDetalleRepository(&detalleVenta, ctx)
	}

	return ventaID, nil

}

func validaStockProduct(detalleVenta *[]dto.DetalleVenta, ctx context.Context) error {
	for _, v := range *detalleVenta {
		stockID, _ := utils.ValidadIdMongo(v.Stock)
		stock, err := stockRopository.BuscarStockRepository(stockID, ctx)
		if err != nil {
			return err
		}
		if v.Cantidad > stock.Cantidad {
			return errors.New("La cantidad del producto es mayor ala de stock: " + v.DescripcionProducto)
		}
	}
	return nil

}

func ListarVentasRealizas(ctx context.Context) (*[]bson.M, error) {
	resultado, err := repository.ListarVentasRepository(ctx)
	if err != nil {
		return &[]bson.M{}, err
	}
	return resultado, nil
}
