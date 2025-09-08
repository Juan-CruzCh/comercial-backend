package router

import (
	"comercial-backend/src/modules/venta/controller"

	"github.com/gin-gonic/gin"
)

func VentaRouter(router *gin.RouterGroup) {
	router.POST("/venta", controller.RealizarVenta)
	router.POST("/venta/listar", controller.ListarVentasRealizas)
	router.GET("/buscar/ventaId/:id", controller.BuscarVentaPorIdController)
}
