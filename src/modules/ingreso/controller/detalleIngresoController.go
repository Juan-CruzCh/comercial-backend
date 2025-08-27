package controller

import (
	"comercial-backend/src/core/utils"
	"comercial-backend/src/modules/ingreso/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ListarDetalleIngresoController(c *gin.Context) {
	var idIngreso string = c.Param("id")
	id, err := utils.ValidadIdMongo(idIngreso)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	reultado, err := service.ListarDetalleIngresoService(id, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reultado)
}
