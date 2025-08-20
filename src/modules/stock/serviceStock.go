package stock

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/modules/stock/dto"
	"comercial-backend/src/modules/stock/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func RegitrarStockService(stockDto []dto.StockDto, ctx context.Context) error {
	collection := config.MongoDatabase.Collection("Stock")
	var StockModel []model.StockModel
	for _, v := range stockDto {
		 productoID, err := bson.ObjectIDFromHex(v.Producto)
		  if err != nil {
				return  err
    		}
		StockModel = append(StockModel, model.StockModel{
			Cantidad:         v.Cantidad,
			FechaVencimiento: v.FechaVencimiento,
			Producto:        productoID,
			PrecioUnitario:   v.PrecioUnitario,
			Codigo:           "",
			Fecha:            time.Now(),
			Flag:             "nuevo",
		})
	}
	_, err := collection.InsertMany(ctx, StockModel)
	if err != nil {

		return err
	}
	return nil
}
