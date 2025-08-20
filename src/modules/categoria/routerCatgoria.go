package categoria

import "github.com/gin-gonic/gin"

func RouterCategoria(router *gin.RouterGroup) {
	router.POST("/categoria", CrearCategoriaController)
	router.GET("/categoria", ListarCategoriaController)
	router.GET("/categoria/:id", ObtenerCategoriaController)
	router.PUT("/categoria/:id", ActualizarCategoriaController)
	router.DELETE("/categoria/:id", EliminarCategoriaController)
}
