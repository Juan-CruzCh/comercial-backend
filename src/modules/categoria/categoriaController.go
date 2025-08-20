package categoria

import (
	"comercial-backend/src/modules/categoria/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Controllers para categoria:

func CrearCategoriaController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var categoria dto.CategoriaDto
	if err := c.ShouldBind(&categoria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := crearCategoriaService(&categoria, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "registrado"})
}

func ListarCategoriaController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	data, err := ListarCategoriaService(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func ObtenerCategoriaController(c *gin.Context) {

}

func ActualizarCategoriaController(c *gin.Context) {

}

func EliminarCategoriaController(c *gin.Context) {

}
