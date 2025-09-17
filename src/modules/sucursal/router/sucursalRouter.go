package sucursal

import (
	"comercial-backend/src/modules/sucursal/controller"

	"github.com/gin-gonic/gin"
)

func SucursalRouter(router *gin.RouterGroup) {
	router.POST("/sucursal", controller.RegistrarSucursalController)
	router.GET("/sucursal", controller.ListarSucursalController)
	router.DELETE("/sucursal/:id", controller.EliminarSucursalController)
}
