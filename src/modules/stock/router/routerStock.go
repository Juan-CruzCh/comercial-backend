package router

import (
	"comercial-backend/src/modules/stock/controller"

	"github.com/gin-gonic/gin"
)

func RouterStock(router *gin.RouterGroup) {
	router.POST("/stock", controller.RegitrarStockController)
	router.GET("/stock",controller.ListarStockController )
}
