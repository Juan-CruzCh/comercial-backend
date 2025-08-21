package stock

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/enum"
	coreUtil "comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso"
	"comercial-backend/src/modules/stock/dto"
	"comercial-backend/src/modules/stock/model"
	structstock "comercial-backend/src/modules/stock/structStock"
	stockUtil "comercial-backend/src/modules/stock/utils"
	"strconv"

	//"comercial-backend/src/modules/stock/utils"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegitrarStockService(body dto.IngresoStockData, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("Stock")
	collectionProducto := config.MongoDatabase.Collection("Producto")
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
	documentos, err := collection.CountDocuments(ctx, bson.M{})
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
		var stock model.StockModel
		err = collection.FindOne(ctx, bson.M{"producto": productoId, "fechaVencimiento": v.FechaVencimiento}).Decode(&stock)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				var producto structstock.Producto
				err = collectionProducto.FindOne(ctx, bson.M{"_id": productoId}).Decode(&producto)
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
				}
				_, err = collection.InsertOne(ctx, nuevoStock)
			} else {
				return err
			}

		} else {
			var cantidad int = v.Cantidad + stock.Cantidad
			_, err := collection.UpdateOne(ctx, bson.M{"_id": stock.ID}, bson.M{"$set": bson.M{"cantidad": cantidad}})
			if err != nil {
				return err
			}
		}

	}

	return nil
}
