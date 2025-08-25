package router

import (
	"comercial-backend/src/modules/venta/controller"

	"github.com/gin-gonic/gin"
)

func VentaRouter(router *gin.RouterGroup) {
	router.POST("/venta", controller.RealizarVenta)
}
