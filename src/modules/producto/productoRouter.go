package producto

import (
	"comercial-backend/src/modules/producto/controller"

	"github.com/gin-gonic/gin"
)

func RouterProducto(router *gin.RouterGroup) {
	//producto
	router.GET("/producto", controller.ListarProductoController)
	router.POST("/producto", controller.RegitrarProductoController)

	//unidadManejo

	router.POST("/unidad/manejo", controller.CrearUnidadController)
}
