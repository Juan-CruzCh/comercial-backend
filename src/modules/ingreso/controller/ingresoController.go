package controller

import (
	"comercial-backend/src/modules/ingreso/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ListarIngresoController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	resultado, err := service.ListarIngresoService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, resultado)
}
