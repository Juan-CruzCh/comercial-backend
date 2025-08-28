package service

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso/model"
	"comercial-backend/src/modules/ingreso/repository"
	ingresoRepository "comercial-backend/src/modules/ingreso/repository"
	"comercial-backend/src/modules/ingreso/structIngreso"
	"context"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegistrarIngresoStockService(body *structIngreso.IngresoStockData, ctx context.Context) (*bson.ObjectID, error) {
	fecha := utils.FechaHoraBolivia()

	documento, err := ingresoRepository.CountDocumentsIngresoRepository(ctx)
	if err != nil {
		return  &bson.NilObjectID, err
	}
	var codigo string = "IGR-" + strconv.Itoa(int(documento))
	var ingreso = model.IngresoModel{
		Codigo:     codigo,
		Fecha:      fecha,
		Proveedor:  body.Proveedor,
		Factura:    body.Factura,
		MontoTotal: body.MontoTotal,
		Flag:       enum.EstadoNuevo,
	}
	ingresoID, err := ingresoRepository.CrearIngresoRepository(&ingreso, ctx)
	if err != nil {
		return  &bson.NilObjectID,err
	}
	var detalleIngreso []model.DetalleIngresoModel
	for _, v := range body.Stock {

		detalleIngreso = append(detalleIngreso, model.DetalleIngresoModel{
			Producto:       v.Producto,
			Cantidad:       v.Cantidad,
			Fecha:          fecha,
			PrecioUnitario: v.PrecioUnitario,
			Flag:           enum.EstadoNuevo,
			Ingreso:        *ingresoID,
			MontoTotal:     v.MontoTotal,
			Descuento:      v.Descuento,
			SudTotal:       v.SudTotal,
		})

	}

	err = repository.CrearDetalleIngresoManyRepository(detalleIngreso, ctx)
	if err != nil {
		return  &bson.NilObjectID ,errors.New("ocurrio un error al ingresar el detalle de ingreso")
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
