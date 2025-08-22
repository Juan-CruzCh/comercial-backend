package stock

import (
	"comercial-backend/src/modules/stock/dto"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegitrarStockController(c *gin.Context) {
	var body dto.IngresoStockData
	err := c.ShouldBindJSON(&body)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	err = RegitrarStockService(body, ctx)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "stock registrado exitosamente"})
}
