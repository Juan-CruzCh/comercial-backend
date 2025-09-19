package router

import (
	"comercial-backend/src/core/enum"
	"comercial-backend/src/core/middleware"
	"comercial-backend/src/modules/venta/controller"

	"github.com/gin-gonic/gin"
)

func VentaRouter(router *gin.RouterGroup) {
	router.POST("/venta", middleware.RoleMiddleware([]string{enum.ADMINISTRADOR, enum.VENDEDOR}), controller.RealizarVenta)
	router.POST("/venta/listar", controller.ListarVentasRealizas)
	router.GET("/buscar/ventaId/:id", controller.BuscarVentaPorIdController)
	router.POST("/reporte/ventas", controller.ReporteVentasController)
	router.GET("/reporte/venta/mensual", controller.ListarReporteVentaMensualController)
}
