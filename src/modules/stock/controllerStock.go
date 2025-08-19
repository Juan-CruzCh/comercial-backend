package stock

import (
	"comercial-backend/src/modules/stock/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegitrarStockController(c *gin.Context) {
	var stock []dto.StockDto
	err := c.ShouldBind(&stock)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()
	err = RegitrarStockService(stock, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "stock registrado exitosamente"})
}
