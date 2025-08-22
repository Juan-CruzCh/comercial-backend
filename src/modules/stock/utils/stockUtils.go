package utils

import (
	"comercial-backend/src/core/utils"

	"comercial-backend/src/modules/ingreso/structIngreso"
	"comercial-backend/src/modules/stock/dto"
)

func ConvertirDtoAIngreso(dto dto.IngresoStockData) (structIngreso.IngresoStockData, error) {
	ProveedorID, err := utils.ValidadIdMongo(dto.Proveedor)
	if err != nil {
		return structIngreso.IngresoStockData{}, err
	}
	stock, err := convertirStock(dto.Stock)
	if err != nil {
		return structIngreso.IngresoStockData{}, err
	}
	return structIngreso.IngresoStockData{
		Proveedor:  *ProveedorID,
		Factura:    dto.Factura,
		MontoTotal: dto.MontoTotal,
		Stock:      stock,
	}, nil
}

func convertirStock(stockDto []dto.StockDto) ([]structIngreso.StockDto, error) {
	var result []structIngreso.StockDto

	for _, s := range stockDto {
		productoID, err := utils.ValidadIdMongo(s.Producto)
		if err != nil {
			return nil, err
		}
		result = append(result, structIngreso.StockDto{
			Cantidad:         s.Cantidad,
			FechaVencimiento: s.FechaVencimiento,
			Producto:         *productoID,
			PrecioUnitario:   s.PrecioUnitario,
			MontoTotal:       s.MontoTotal,
			Descuento:        s.Descuento,
			SudTotal:         s.SudTotal,
		})
	}
	return result, nil
}

func GenerarCodigoStock() {

}
