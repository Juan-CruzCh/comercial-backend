package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso/model"
	"comercial-backend/src/modules/ingreso/repository"
	ingresoRepository "comercial-backend/src/modules/ingreso/repository"
	"comercial-backend/src/modules/stock/dto"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegistrarIngresoStockService(body *dto.IngresoStockDto, ctx context.Context, usuario *bson.ObjectID) (*bson.ObjectID, error) {
	fecha := utils.FechaHoraBolivia()
	ProveedorID, err := utils.ValidadIdMongo(body.Proveedor)
	if err != nil {
		return nil, err
	}
	documento, err := ingresoRepository.CountDocumentsIngresoRepository(ctx)
	if err != nil {
		return nil, err
	}
	var PrecioUnitarioTotalCompra float64 = 0
	var PrecioUnitarioTotal float64 = 0
	var subTotal float64 = 0
	var descuento float64 = 0
	var cantidadTotal = 0
	for _, v := range body.Stock {
		PrecioUnitarioTotal += v.PrecioUnitario * float64(v.Cantidad)
		subTotal += (v.PrecioUnitarioCompra * float64(v.Cantidad))
		descuento += v.Descuento
		PrecioUnitarioTotalCompra += (float64(v.Cantidad) * v.PrecioUnitarioCompra) - v.Descuento
		cantidadTotal += v.Cantidad
	}

	var codigo string = "IGR-" + strconv.Itoa(int(documento))
	var ingreso = model.IngresoModel{
		Codigo:                    codigo,
		Fecha:                     fecha,
		Proveedor:                 *ProveedorID,
		Factura:                   body.Factura,
		PrecioUnitarioTotal:       PrecioUnitarioTotal,
		TotalDescuento:            descuento,
		Flag:                      enum.EstadoNuevo,
		Usuario:                   *usuario,
		SudTotal:                  subTotal,
		PrecioUnitarioTotalCompra: PrecioUnitarioTotalCompra,
		CantidadTotal:             cantidadTotal,
	}
	ingresoID, err := ingresoRepository.CrearIngresoRepository(&ingreso, ctx)
	if err != nil {
		return &bson.NilObjectID, err
	}

	return ingresoID, err

}

func RegitrarDetalleIngresoService(detalle dto.StockDtoDetalleDto, ingresoID *bson.ObjectID, codigoStock string, ctx context.Context) error {
	fecha := utils.FechaHoraBolivia()
	productoID, err := utils.ValidadIdMongo(detalle.Producto)
	if err != nil {
		return err
	}
	var fechaVencimiento *time.Time
	if !detalle.FechaVencimiento.IsZero() {
		fechaVencimiento = &detalle.FechaVencimiento
	}
	var detalleIngreso model.DetalleIngresoModel = model.DetalleIngresoModel{
		Producto:                  *productoID,
		Cantidad:                  detalle.Cantidad,
		Fecha:                     fecha,
		PrecioUnitario:            detalle.PrecioUnitario,
		Flag:                      enum.EstadoNuevo,
		Ingreso:                   *ingresoID,
		PrecioUnitarioCompra:      detalle.PrecioUnitarioCompra,
		Descuento:                 detalle.Descuento,
		SubTotal:                  (detalle.PrecioUnitarioCompra * float64(detalle.Cantidad)),
		FechaVencimiento:          fechaVencimiento,
		CodigoStock:               codigoStock,
		PrecioUnitarioTotal:       detalle.PrecioUnitario * float64(detalle.Cantidad),
		PrecioUnitarioTotalCompra: detalle.PrecioUnitarioCompra*float64(detalle.Cantidad) - detalle.Descuento,
	}
	err = repository.CrearDetalleIngresoRepository(detalleIngreso, ctx)
	if err != nil {
		return err
	}
	return nil
}

func ListarIngresoService(ctx context.Context) (*[]bson.M, error) {

	resultado, err := repository.ListarIngresoRepository(ctx)
	if err != nil {
		return &[]bson.M{}, err
	}
	return resultado, nil
}
