package stock

import "github.com/gin-gonic/gin"

func RouterStock(router *gin.RouterGroup) {
	router.POST("/stock", RegitrarStockController)
}
