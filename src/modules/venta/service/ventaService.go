package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/structCore"
	"comercial-backend/src/core/utils"
	cajaRopository "comercial-backend/src/modules/caja/repository"
	descuentoVentaRepository "comercial-backend/src/modules/descuentoVenta/repository"
	stockRopository "comercial-backend/src/modules/stock/repository"
	"errors"
	"strconv"

	"comercial-backend/src/modules/venta/dto"
	"comercial-backend/src/modules/venta/model"
	"comercial-backend/src/modules/venta/repository"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RealizarVentaService(body *dto.VentaDto, ctx context.Context, usuarioID *bson.ObjectID, sucursalID *bson.ObjectID) (*bson.ObjectID, error) {
	caja, err := cajaRopository.BuscarCajaUsuarioRepository(usuarioID, ctx)
	if err != nil {
		return nil, errors.New("Ocurrio un erro en la caja de venta " + err.Error())
	}

	fecha := utils.FechaHoraBolivia()
	err = validaStockProducto(&body.DetalleVenta, ctx)
	if err != nil {
		return &bson.NilObjectID, err
	}

	var montoTotal float64 = 0
	var subTotal float64 = 0
	for _, v := range body.DetalleVenta {
		subTotal += v.PrecioUnitario * float64(v.Cantidad)
		montoTotal += v.PrecioUnitario * float64(v.Cantidad)
	}
	if *body.Descuento > subTotal {
		return nil, errors.New("el descuento no puede ser mayor al subTotal")
	}
	cantidad, _ := repository.CountDocumentsVentaRepository(ctx)
	var codigo string = "VEN-" + strconv.Itoa(int(cantidad))
	montoTotal = montoTotal - *body.Descuento
	alquiler, vendedor, ganancia, descuentoAcumulado := realizarDescuentoVenta(sucursalID, montoTotal, ctx)
	var venta model.VentaModel = model.VentaModel{
		Codigo:             codigo,
		MontoTotal:         utils.RoundFloat(montoTotal, 2),
		FechaVenta:         fecha,
		Fecha:              fecha,
		Usuario:            *usuarioID,
		Sucursal:           *sucursalID,
		TipoPago:           enum.Efectivo,
		Estado:             enum.Realizada,
		Flag:               enum.EstadoNuevo,
		Descuento:          *body.Descuento,
		DescuentoAlquiller: alquiler,
		DescuentoVendedor:  vendedor,
		TotalGanancia:      ganancia,
		DescuentoAcumulado: descuentoAcumulado,
		SubTotal:           utils.RoundFloat(subTotal, 2),
	}
	ventaID, err := repository.RealizarVentaRepository(&venta, ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range body.DetalleVenta {
		stockID, err := utils.ValidadIdMongo(v.Stock)
		if err != nil {
			return nil, err
		}
		stock, err := stockRopository.BuscarStockRepository(stockID, ctx)
		if err != nil {
			return nil, err
		}
		var nuevaCantidad int = stock.Cantidad - v.Cantidad
		err = stockRopository.ActualizarStockRepository(stock.ID, nuevaCantidad, ctx)
		if err != nil {
			return nil, err
		}
		var precioTotalDetalle float64 = utils.RoundFloat(v.PrecioUnitario*float64(v.Cantidad), 2)
		var detalleVenta model.DetalleVentaModel = model.DetalleVentaModel{
			Producto:       stock.Producto,
			Stock:          *stockID,
			Cantidad:       v.Cantidad,
			Descripcion:    v.DescripcionProducto,
			Venta:          *ventaID,
			Fecha:          fecha,
			Flag:           enum.EstadoNuevo,
			PrecioUnitario: v.PrecioUnitario,
			PrecioTotal:    precioTotalDetalle,
		}
		_ = repository.RealizarVentaDetalleRepository(&detalleVenta, ctx)
	}

	var totalVenta float64 = utils.RoundFloat(caja.TotalVentas+montoTotal, 2)
	var montoFinal float64 = utils.RoundFloat(totalVenta+caja.MontoInicial, 2)
	var totalDescuento float64 = utils.RoundFloat(*body.Descuento+caja.TotalDescuento, 2)
	err = cajaRopository.AsignarTotalVentasCajaRepository(usuarioID, totalVenta, montoFinal, totalDescuento, ctx)
	if err != nil {
		return &bson.NilObjectID, errors.New("Ocurrio un error en la caja de venta al asignar el total vendido " + err.Error())
	}
	return ventaID, nil

}
func realizarDescuentoVenta(sucursal *bson.ObjectID, montoTotal float64, ctx context.Context) (al float64, ven float64, ganacia float64, totalDescuento float64) { //realiza descuento de cada venta un porcentaje para el alquiler y el vendedor
	var alquiler float64 = 0
	var vendedor float64 = 0

	data, err := descuentoVentaRepository.ObtenerDescuentoVentaRepository(sucursal, ctx)
	if err != nil {
		return 0, 0, montoTotal, 0
	}

	var porcentajeAlquiler float64 = utils.Porcentaje(data.Alquiler)
	var porcentajeVendedor float64 = utils.Porcentaje(data.Vendedor)
	alquiler = (porcentajeAlquiler / 100) * montoTotal //se extrae la en plara el porcentaje
	vendedor = (porcentajeVendedor / 100) * montoTotal //se extrae la en plara el porcentaje
	alquiler = utils.RoundFloat(alquiler, 2)
	vendedor = utils.RoundFloat(vendedor, 2)
	var descuentoAcumulado float64 = alquiler + vendedor
	var totalGanancia float64 = montoTotal - descuentoAcumulado
	totalGanancia = utils.RoundFloat(totalGanancia, 2)

	return alquiler, vendedor, totalGanancia, descuentoAcumulado
}

func validaStockProducto(detalleVenta *[]dto.DetalleVenta, ctx context.Context) error {
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

func ListarVentasRealizas(filtros *dto.BuscadorVentaDto, pagina int, limite int, ctx context.Context) (*structCore.ResultadoPaginado, error) {
	resultado, err := repository.ListarVentasRepository(filtros, pagina, limite, true, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func BuscarVentaPorIdService(idVenta *bson.ObjectID, ctx context.Context) (*bson.M, error) {
	data, err := repository.BuscarVentaPorIdRespository(idVenta, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReporteVentasService(filtros *dto.BuscadorVentaDto, ctx context.Context) (*structCore.ResultadoPaginado, error) {
	data, err := repository.ListarVentasRepository(filtros, 0, 0, false, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
