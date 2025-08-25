package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	stockRopository "comercial-backend/src/modules/stock/repository"
	"comercial-backend/src/modules/venta/dto"
	"comercial-backend/src/modules/venta/model"
	"comercial-backend/src/modules/venta/repository"
	"context"
	"fmt"
)

func RealizarVentaService(body *dto.VentaDto, ctx context.Context) error {

	fecha := utils.FechaHoraBolivia()

	usuarioID, err := utils.ValidadIdMongo("68a8b4dbdb1be7def32f34a0")
	if err != nil {
		return err
	}
	sucursalID, err := utils.ValidadIdMongo("68a8b4dbdb1be7def32f34a0")
	if err != nil {
		return err
	}
	var venta model.VentaModel = model.VentaModel{
		Codigo:     "FALTA",
		MontoTotal: body.MontoTotal,
		FechaVenta: fecha,
		Fecha:      fecha,
		Usuario:    *usuarioID,
		Sucursal:   *sucursalID,
		TipoPago:   enum.Efectivo,
		Estado:     enum.Completado,
		Flag:       enum.EstadoNuevo,
	}
	_, err = repository.RealizarVentaRepository(&venta, ctx)
	if err != nil {
		return err

	}

	for _, v := range body.DetalleVenta {
		stockID, err := utils.ValidadIdMongo(v.Stock)
		if err != nil {
			return err
		}
		stock, err := stockRopository.BuscarStockRepository(stockID, ctx)
		if err != nil {
			return err
		}
		var nuevaCantidad int = stock.Cantidad - v.Cantidad
		fmt.Println(nuevaCantidad)
		err = stockRopository.ActualizarStockRepository(stock.ID, nuevaCantidad, ctx)
		if err != nil {
			return err
		}
	}

	return nil

}
