package categoria

import "github.com/gin-gonic/gin"

func RouterCategoria(router *gin.RouterGroup) {
	router.POST("/categorias", CrearCategoriaController)
	router.GET("/categorias", ListarCategoriaController)
	router.GET("/categorias/:id", ObtenerCategoriaController)
	router.PUT("/categorias/:id", ActualizarCategoriaController)
	router.DELETE("/categorias/:id", EliminarCategoriaController)
}
