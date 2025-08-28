package router

import (
	"comercial-backend/src/modules/caja/controller"

	"github.com/gin-gonic/gin"
)

func CajaRouter(router *gin.RouterGroup) {
	router.POST("/abrir/caja", controller.AbriCajaController)
	router.POST("/cerrar/caja", controller.CerrarCajaController)
	router.POST("/verificar/caja", controller.VerificarCajaController)
}
