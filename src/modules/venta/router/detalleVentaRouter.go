package router

import (
	"comercial-backend/src/modules/venta/controller"

	"github.com/gin-gonic/gin"
)

func DetalleVentaRouter(router *gin.RouterGroup) {
	router.GET("detalle/venta/:id", controller.DetalleVentaController)
}
