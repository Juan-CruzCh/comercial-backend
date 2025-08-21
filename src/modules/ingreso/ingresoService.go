package ingreso

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso/model"
	"context"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegistrarIngresoStockService(body *IngresoStockData, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("Ingreso")
	collectionDetalleIngreso := config.MongoDatabase.Collection("DetalleIngreso")
	fecha, err := utils.FechaHoraBolivia()
	if err != nil {
		return err
	}
	documento, err := collection.CountDocuments(ctx, bson.M{"flag": enum.EstadoNuevo})
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
	result, err := collection.InsertOne(ctx, ingreso)
	ingresoID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return errors.New("ocurrio un error al insertar el ingreso")
	}
	var detalleIngreso []model.DetalleIngresoModel
	for _, v := range body.Stock {

		detalleIngreso = append(detalleIngreso, model.DetalleIngresoModel{
			Producto:       v.Producto,
			Cantidad:       v.Cantidad,
			Fecha:          fecha,
			PrecioUnitario: v.PrecioUnitario,
			Flag:           enum.EstadoNuevo,
			Ingreso:        ingresoID,
			MontoTotal:     v.MontoTotal,
			Descuento:      v.Descuento,
			SudTotal:       v.SudTotal,
		})

	}
	_, err = collectionDetalleIngreso.InsertMany(ctx, detalleIngreso)
	if err != nil {
		return err
	}
	return nil

}
