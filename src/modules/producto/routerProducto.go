package producto

import "github.com/gin-gonic/gin"

func RouterProduct(router *gin.RouterGroup) {
	router.POST("/producto", RegitrarProductoController)
	router.GET("/producto", ListarProductoController)
}
