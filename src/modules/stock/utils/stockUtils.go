package utils

import (
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso"
	"comercial-backend/src/modules/stock/dto"
)

func ConvertirDtoAIngreso(dto dto.IngresoStockData) (ingreso.IngresoStockData, error) {
	ProveedorID, err := utils.ValidadIdMongo(dto.Proveedor)
	if err != nil {
		return ingreso.IngresoStockData{}, err
	}
	stock, err := convertirStock(dto.Stock)
	if err != nil {
		return ingreso.IngresoStockData{}, err
	}
	return ingreso.IngresoStockData{
		Proveedor:  *ProveedorID,
		Factura:    dto.Factura,
		MontoTotal: dto.MontoTotal,
		Stock:      stock,
	}, nil
}

func convertirStock(stockDto []dto.StockDto) ([]ingreso.StockDto, error) {
	var result []ingreso.StockDto

	for _, s := range stockDto {
		productoID, err := utils.ValidadIdMongo(s.Producto)
		if err != nil {
			return nil, err
		}
		result = append(result, ingreso.StockDto{
			Cantidad:         s.Cantidad,
			FechaVencimiento: s.FechaVencimiento,
			Producto:         *productoID,
			PrecioUnitario:   s.PrecioUnitario,
		})
	}
	return result, nil
}
