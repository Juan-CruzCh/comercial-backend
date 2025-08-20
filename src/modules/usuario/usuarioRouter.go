package usuario

import "github.com/gin-gonic/gin"

func UsuarioRouter(router *gin.RouterGroup) {
	router.POST("/usuarios", crearUsuarioController)
	router.GET("/usuarios", listarUsuarioController)
	router.GET("/usuarios/:id", obtenerUsuarioController)
	router.PUT("/usuarios/:id", actualizarUsuarioController)
	router.DELETE("/usuarios/:id", eliminarUsuarioController)
}
