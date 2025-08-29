package router

import (
	"comercial-backend/src/modules/autenticacion/controller"

	"github.com/gin-gonic/gin"
)

func AutenticacionRouter(router *gin.RouterGroup) {
	router.POST("/autenticacion", controller.AutenticacionController)
	router.GET("/logout", controller.CerrarAuntenticacionController)
}
