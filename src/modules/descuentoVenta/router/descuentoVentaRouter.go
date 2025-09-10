package router

import (
	"comercial-backend/src/modules/descuentoVenta/controller"

	"github.com/gin-gonic/gin"
)

func DescuentoVentaRouter(router *gin.RouterGroup) {
	router.POST("/descuento/venta", controller.CrearDescuentoVentaController)
	router.GET("/descuento/venta", controller.ListarDescuentoVentaController)
}
