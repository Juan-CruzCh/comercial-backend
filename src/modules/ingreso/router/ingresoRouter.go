package router

import (
	"comercial-backend/src/modules/ingreso/controller"

	"github.com/gin-gonic/gin"
)

func IngresoRouter(router *gin.RouterGroup) {
	router.GET("/ingreso", controller.ListarIngresoController)
}
