package router

import (
	"comercial-backend/src/modules/caja/controller"

	"github.com/gin-gonic/gin"
)

func CajaRouter(router *gin.RouterGroup) {
	router.POST("/abrir/caja", controller.AbriCajaController)
}
