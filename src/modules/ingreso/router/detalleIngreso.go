package router

import (
	"comercial-backend/src/modules/ingreso/controller"

	"github.com/gin-gonic/gin"
)

func DetalleIngresoRouter(router *gin.RouterGroup) {
	router.GET("detalle/ingreso/:id", controller.ListarDetalleIngresoController)
}
