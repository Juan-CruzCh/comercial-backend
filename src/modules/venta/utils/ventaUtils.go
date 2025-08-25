package utils

import (
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/venta/dto"
)

func ValidarDetalleVentaBsonObjectId(detalleVenta *[]dto.DetalleVenta) error {
	for _, v := range *detalleVenta {
		_, err := utils.ValidadIdMongo(v.Stock)
		if err != nil {
			return err
		}
	}
	return nil
}
