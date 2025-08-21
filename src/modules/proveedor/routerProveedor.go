package proveedor

import "github.com/gin-gonic/gin"

func RouterProveedor(router *gin.RouterGroup) {
	router.POST("/proveedor", registrarProveedorController)
	router.GET("/proveedor", listarProveedorController)
}
