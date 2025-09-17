package router

import (
	"comercial-backend/src/modules/usuario/controller"

	"github.com/gin-gonic/gin"
)

func UsuarioRouter(router *gin.RouterGroup) {
	router.POST("/usuario", controller.CrearUsuarioController)
	router.GET("/usuario", controller.ListarUsuarioController)
	router.GET("/usuario/:id", controller.ObtenerUsuarioController)
	router.PATCH("/usuario/:id", controller.ActualizarUsuarioController)
	router.DELETE("/usuario/:id", controller.EliminarUsuarioController)
	router.GET("/usuario/verificar", controller.VerificarAutenticacionUsuarioController)
	router.GET("/usuario/logout", controller.CerrarAuntenticacionUsuarioController)
}
