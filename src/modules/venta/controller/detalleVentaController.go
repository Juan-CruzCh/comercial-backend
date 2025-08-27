package controller

import (
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/venta/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DetalleVentaController(c *gin.Context) {
	var idVenta string = c.Param("id")
	venta, err := utils.ValidadIdMongo(idVenta)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := service.DetalleVentaService(venta, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultado)

}
