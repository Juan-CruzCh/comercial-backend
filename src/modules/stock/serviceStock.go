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
		StockModel = append(StockModel, model.StockModel{
			Cantidad:         v.Cantidad,
			FechaVencimiento: v.FechaVencimiento,
			Producto:         bson.ObjectID{},
			PrecioUnitario:   v.PrecioUnitario,
			MontoTotal:       v.MontoTotal,
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
