package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso/model"
	"comercial-backend/src/modules/ingreso/repository"
	ingresoRepository "comercial-backend/src/modules/ingreso/repository"
	"comercial-backend/src/modules/stock/dto"
	"context"
	"errors"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegistrarIngresoStockService(body *dto.IngresoStockData, ctx context.Context, usuario *bson.ObjectID) (*bson.ObjectID, error) {
	fecha := utils.FechaHoraBolivia()
	ProveedorID, err := utils.ValidadIdMongo(body.Proveedor)
	documento, err := ingresoRepository.CountDocumentsIngresoRepository(ctx)
	if err != nil {
		return &bson.NilObjectID, err
	}
	var codigo string = "IGR-" + strconv.Itoa(int(documento))
	var ingreso = model.IngresoModel{
		Codigo:     codigo,
		Fecha:      fecha,
		Proveedor:  *ProveedorID,
		Factura:    body.Factura,
		MontoTotal: body.MontoTotal,
		Flag:       enum.EstadoNuevo,
		Usuario:    *usuario,
	}
	ingresoID, err := ingresoRepository.CrearIngresoRepository(&ingreso, ctx)
	if err != nil {
		return &bson.NilObjectID, err
	}
	var detalleIngreso []model.DetalleIngresoModel
	for _, v := range body.Stock {
		productoID, err := utils.ValidadIdMongo(v.Producto)
		if err != nil {
			return nil, err
		}
		var fechaVencimiento *time.Time
		if !v.FechaVencimiento.IsZero() {
			fechaVencimiento = &v.FechaVencimiento
		}
		detalleIngreso = append(detalleIngreso, model.DetalleIngresoModel{
			Producto:         *productoID,
			Cantidad:         v.Cantidad,
			Fecha:            fecha,
			PrecioUnitario:   v.PrecioUnitario,
			Flag:             enum.EstadoNuevo,
			Ingreso:          *ingresoID,
			MontoTotal:       v.MontoTotal,
			Descuento:        v.Descuento,
			SudTotal:         v.SudTotal,
			FechaVencimiento: fechaVencimiento,
		})

	}

	err = repository.CrearDetalleIngresoManyRepository(detalleIngreso, ctx)
	if err != nil {
		return &bson.NilObjectID, errors.New("ocurrio un error al ingresar el detalle de ingreso")
	}
	return ingresoID, err

}

func ListarIngresoService(ctx context.Context) (*[]bson.M, error) {

	resultado, err := repository.ListarIngresoRepository(ctx)
	if err != nil {
		return &[]bson.M{}, err
	}
	return resultado, nil
}
