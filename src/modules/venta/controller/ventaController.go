package controller

import (
	"comercial-backend/src/modules/venta/dto"
	"comercial-backend/src/modules/venta/service"
	"comercial-backend/src/modules/venta/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RealizarVenta(c *gin.Context) {
	validate := validator.New()
	ctx, cancel := context.WithTimeout(c.Request.Context(), 20*time.Second)
	defer cancel()
	var body dto.VentaDto
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = utils.ValidarDetalleVentaBsonObjectId(&body.DetalleVenta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.RealizarVentaService(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
