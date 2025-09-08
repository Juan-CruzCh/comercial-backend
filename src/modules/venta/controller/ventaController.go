package controller

import (
	"comercial-backend/src/core/utils"
	coreUtils "comercial-backend/src/core/utils"
	"comercial-backend/src/modules/venta/dto"
	"comercial-backend/src/modules/venta/service"
	ventaUtils "comercial-backend/src/modules/venta/utils"

	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RealizarVenta(c *gin.Context) {
	usuarioID, sucursalID, err := coreUtils.Request(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	ctx, cancel := context.WithTimeout(c.Request.Context(), 20*time.Second)
	defer cancel()
	var body dto.VentaDto
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ventaUtils.ValidarDetalleVentaBsonObjectId(&body.DetalleVenta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idVenta, err := service.RealizarVentaService(&body, ctx, usuarioID, sucursalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"idVenta": idVenta})
}

func ListarVentasRealizas(c *gin.Context) {
	validate := validator.New()
	pagina, limite, err := utils.Paginador(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	var body dto.BuscadorVentaDto
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado, err := service.ListarVentasRealizas(&body, pagina, limite, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultado)
}

func BuscarVentaPorIdController(c *gin.Context) {
	idVenta := c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	ID, err := utils.ValidadIdMongo(idVenta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := service.BuscarVentaPorIdService(ID, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)

}
