package ingreso

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
)

func RegistrarIngresoStockService(body *structIngreso.IngresoStockData, ctx context.Context) error {
	fecha, err := utils.FechaHoraBolivia()
	if err != nil {
		return err
	}
	documento, err := ingresoRepository.CountDocumentsIngresoRepository(ctx)
	if err != nil {
		return err
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
		return err
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
		return errors.New("ocurrio un error al ingresar el detalle de ingreso")
	}
	return nil

}
