package almacen

import "github.com/gin-gonic/gin"

func AlmacenRouter(router *gin.RouterGroup) {
	router.POST("/almacen", registrarAlmacenController)
	router.GET("/almacen", listarAlmacenesController)
	router.GET("/almacen/:id", obtenerAlmacenController)
	router.PUT("/almacen/:id", actualizarAlmacenController)
	router.DELETE("/almacen/:id", eliminarAlmacenController)
}
