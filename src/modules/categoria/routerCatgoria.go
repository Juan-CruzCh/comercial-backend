package categoria

import "github.com/gin-gonic/gin"

func RouterCategoria(router *gin.RouterGroup) {
	router.POST("/categorias", CrearCategoria)
	router.GET("/categorias", ListarCategoria)
}
